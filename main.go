package main

import (
	"be-4-rs-crud/api"
	db "be-4-rs-crud/src/db/sqlc"
	"be-4-rs-crud/utils"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:mysecretpassword@localhost:5431/simple_trx?sslmode=disable"
// 	serverAddress = "0.0.0.0:8888"
// )

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSoruce)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
