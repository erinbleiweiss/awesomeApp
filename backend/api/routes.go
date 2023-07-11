package api

import (
	"database/sql"
	"github.com/gorilla/mux"
)

func AddRoutesV1(r *mux.Router, db *sql.DB) {
	app := NewHandler(db)
	r.HandleFunc("/register", app.register).Methods("POST")
	r.HandleFunc("/login", app.login).Methods("POST")
	r.HandleFunc("/user/{id}", app.getUser).Methods("GET")
	r.HandleFunc("/users", app.getUsers).Methods("GET")
}
