// app.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// struct that exposes refs to App's router and db
// Init and Run defines the struct behavior to make it useful and testable behavior
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// wire up routes to to res and db details
func (app *App) Initialize(user, password, dbname string) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// start app server
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", app.Router))
}

type Form struct {
	Name  string
	Email string
}

// func (app *App) subscribe(form Form, db *sql.DB) error {
// 	return errors.New("Not implemented")
// }

func (app *App) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(http.StatusOK)
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/health_check", app.healthCheck).Methods("GET")
	// app.Router.HandleFunc("/subscriptions", app.subscribe).Methods("POST")
}

// run test against a db
// setup db
// db cleaned after running

// docker run -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres

// go mod init github.com/onyewuenyi/zero2prod_golang
