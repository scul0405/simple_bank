package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/scul0405/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	password := util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreateAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreateAt, user2.CreateAt, time.Second)
}

func TestUpdateUser(t *testing.T){
	user := createRandomUser(t)

	testcases := []struct{
		name string
		runTest func(t *testing.T)
	}{
		{
			name: "Update only full name",
			runTest: func(t *testing.T) {
				newFullName := util.RandomOwner()
				arg := UpdateUserParams{
					Username: user.Username,
					FullName: sql.NullString{
						String: newFullName,
						Valid: true,
					},
				}

				newUser, err := testQueries.UpdateUser(context.Background(), arg)
				require.NoError(t, err)
				require.Equal(t, newUser.Username, user.Username)
				require.Equal(t, newUser.FullName, newFullName)
				require.Equal(t, newUser.Email, user.Email)
			},
		},
		{
			name: "Update only email",
			runTest: func(t *testing.T) {
				newEmail := util.RandomEmail()
				arg := UpdateUserParams{
					Username: user.Username,
					Email: sql.NullString{
						String: newEmail,
						Valid: true,
					},
				}

				newUser, err := testQueries.UpdateUser(context.Background(), arg)
				require.NoError(t, err)
				require.Equal(t, newUser.Username, user.Username)
				require.Equal(t, newUser.FullName, newUser.FullName)
				require.Equal(t, newUser.Email, newEmail)
			},
		},
		{
			name: "Update all fields",
			runTest: func(t *testing.T) {
				newFullName := util.RandomOwner()
				newEmail := util.RandomEmail()
				arg := UpdateUserParams{
					Username: user.Username,
					FullName: sql.NullString{
						String: newFullName,
						Valid: true,
					},
					Email: sql.NullString{
						String: newEmail,
						Valid: true,
					},
				}

				newUser, err := testQueries.UpdateUser(context.Background(), arg)
				require.NoError(t, err)
				require.Equal(t, newUser.Username, user.Username)
				require.Equal(t, newUser.FullName, newFullName)
				require.Equal(t, newUser.Email, newEmail)
			},
		},
	}

	for i := range testcases {
		tc := testcases[i]

		tc.runTest(t)
	}
}