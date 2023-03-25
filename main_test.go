package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const createTable = `
CREATE TABLE IF NOT EXISTS accounts (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	balance REAL NOT NULL
);
`

const dropTable = `
DROP TABLE IF EXISTS accounts;
`

const insert = `
INSERT INTO accounts (name, balance) VALUES (?, ?) RETURNING *;
`

func TestMain(m *testing.M) {
	db, _ = sql.Open("sqlite3", "file:test.db")
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mig, err := migrate.NewWithDatabaseInstance("file://db/migrations", "test.db", instance)
	if err != nil {
		log.Fatalf("err mig: %v", err)
	}

	if err := mig.Up(); err != nil {
		log.Fatal(err)
	}
	i := m.Run()
	fmt.Println("after")
	fmt.Println("i---", i)
	if err := mig.Down(); err != nil {
		log.Fatal(err)
		os.Exit(i)
	}
	fmt.Println("dropped")
	os.Exit(i)
}
