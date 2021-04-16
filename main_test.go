// main_test.go
package main_test

import (
	"os"
	"testing"
)

var app main.App

func TestMain(m *testing.M) {
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PW"),
		os.Getenv("APP_DB_NAME"))

	// TODO what does this do. Rust does not do this check

	// ensureTableExists()

	code := m.Run()
	clearTable()
	os.Exit(code)
}

// fn ensureTableExists, which does some db checking
func clearTable() {
	app.DB.Exec("DELETE FROM  subscriptions")
	app.DB.Exec("ALTER SEQUENCE ")
}

// specified in Rust migrations/*.sql
const tableCreationQuery = `CREATE TABLE IF NOT EXISTS subscriptions(
    id uuid NOT NULL,
    PRIMARY KEY (id),
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    subscribed_as timestamptz NOT NULL
)`
