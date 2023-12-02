package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const(
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_cars5?sslmode=disable"
)

var testQueries *Queries

//Test the main function
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}