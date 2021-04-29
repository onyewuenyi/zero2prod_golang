// main_test.go
package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/onyewuenyi/zero2prod_golang/main"
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

	// clearTable()

	os.Exit(code)
}

func TestHealthCheck(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health_check", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

// // fn ensureTableExists, which does some db checking
// func clearTable() {
// 	app.DB.Exec("DELETE FROM  subscriptions")
// 	app.DB.Exec("ALTER SEQUENCE ")
// }

// // specified in Rust migrations/*.sql
// const tableCreationQuery = `CREATE TABLE IF NOT EXISTS subscriptions(
//     id uuid NOT NULL,
//     PRIMARY KEY (id),
//     email TEXT NOT NULL UNIQUE,
//     name TEXT NOT NULL,
//     subscribed_as timestamptz NOT NULL
// )`

// func ensureTableExists() {
// 	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
// 		log.Fatal(err)
// 	}
// }
