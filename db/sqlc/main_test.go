package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	config, err := util.LoadConfig("../..") // location of app.env file
	if err != nil {
		log.Fatal("Cannot load configuration: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}