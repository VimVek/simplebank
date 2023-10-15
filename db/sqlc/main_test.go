package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

// defining it as a global variable because we will use it extensivley in all our unit tests.
var testQueries *Queries
var testDB *sql.DB

// build a db connection and use it to create the queries object
// test main function is the main entry pointof all unit test inside 1 specific golang package(package db in this case)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
