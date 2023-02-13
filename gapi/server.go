package gapi

import (
	"fmt"

	db "github.com/scul0405/simple_bank/db/sqlc"
	"github.com/scul0405/simple_bank/pb"
	"github.com/scul0405/simple_bank/token"
	"github.com/scul0405/simple_bank/util"
)

// Server serves GRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimplebankServer
	store      db.Store
	config     util.Config
	tokenMaker token.Maker
}

// NewServer creates a new GRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPASETOMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
