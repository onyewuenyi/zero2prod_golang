// app.go

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(user, password, dbname string) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	app.Router = mux.NewRouter()
}

func (app *App) Run(addr string) {}

type form struct {
	Name  string
	Email string
}

func (f *form) subscribe(form form, db *sql.DB) error {
	return errors.New("Not implemented")
}

// run test against a db
// setup db
// db cleaned after running

// docker run -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres


go mod init github.com/onyewuenyi/zero2prod_golang