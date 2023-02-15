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

func TestUpdateOnlyFullName(t *testing.T){
	user := createRandomUser(t)
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
	require.Equal(t, newUser.HashedPassword, user.HashedPassword)
	require.Equal(t, newUser.Email, user.Email)

}


func TestUpdateOnlyEmail(t *testing.T){
	user := createRandomUser(t)
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
	require.Equal(t, newUser.HashedPassword, user.HashedPassword)
	require.Equal(t, newUser.Email, newEmail)
}

func TestUpdateOnlyPassword(t *testing.T){
	user := createRandomUser(t)
	newPassword := util.RandomString(6)
	newHashPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	arg := UpdateUserParams{
		Username: user.Username,
		HashedPassword: sql.NullString{
			String: newHashPassword,
			Valid: true,
		},
	}

	newUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.FullName, user.FullName)
	require.Equal(t, newUser.HashedPassword, newHashPassword)
	require.Equal(t, newUser.Email, user.Email)
}

func TestUpdateAllFields(t *testing.T){
	user := createRandomUser(t)
	newFullName := util.RandomOwner()
	newEmail := util.RandomEmail()
	newPassword := util.RandomString(6)
	newHashPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

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
		HashedPassword: sql.NullString{
			String: newHashPassword,
			Valid: true,
		},
	}

	newUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.FullName, newFullName)
	require.Equal(t, newUser.HashedPassword, newHashPassword)
	require.Equal(t, newUser.Email, newEmail)
}