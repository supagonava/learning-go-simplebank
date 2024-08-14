package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var conn *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	conn, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}
	defer conn.Close()

	testQueries = New(conn)
	os.Exit(m.Run())
}
