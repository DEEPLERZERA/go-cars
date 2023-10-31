package main

import (
	"database/sql"
	"log"

	"github.com/DEEPLERZERA/go-cars/api"
	db "github.com/DEEPLERZERA/go-cars/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/go_cars5?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.ExecuteNewStore(conn)
	server := api.InstanceServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
