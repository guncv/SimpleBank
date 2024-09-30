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
	dbSource = "postgresql://root:secretpassword@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	log.Println("Setting up tests...")
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
	log.Println("Tests finished")
	os.Exit(m.Run())
}

func TestSampleFunction(t *testing.T) {
	t.Run("example test", func(t *testing.T) {
		// ใส่ logic ทดสอบของคุณ
		log.Println("TestSampleFunction finished, cleaning up...")
		if 1+1 != 2 {
			t.Error("Test failed")
		}
	})
}
