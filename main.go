package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/techshool/simplebank/api"
	db "github.com/techshool/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:postgrespw@localhost:49153/simple_bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}