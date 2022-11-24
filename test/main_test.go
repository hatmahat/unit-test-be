package test

import (
	db "be-4-rs-crud/src/db/sqlc"
	"be-4-rs-crud/utils"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSoruce)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
