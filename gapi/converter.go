package gapi

import (
	db "github.com/scul0405/simple_bank/db/sqlc"
	"github.com/scul0405/simple_bank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordChangeAt: timestamppb.New(user.PasswordChangedAt),
		CreateAt:         timestamppb.New(user.CreateAt),
	}
}
