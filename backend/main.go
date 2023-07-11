package main

import (
	"database/sql"
	"github.com/erinbleiweiss/awesomeApp/backend/api"
	"github.com/erinbleiweiss/awesomeApp/backend/config"
	"github.com/erinbleiweiss/awesomeApp/backend/db"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Init(dbConfig *config.DBConfig) {
	appDB, err := db.NewDB(dbConfig.DBPath)
	if err != nil {
		log.Fatal("Error creating DB: ", err)
	}
	app.DB = appDB

	app.Router = api.CreateRouter(appDB)
}

func (app *App) Run(addr string) {
	srv := &http.Server{
		Handler:      app.Router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server running at http://%s...", addr)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	app := App{}
	dbConfig, err := config.GetConfig()
	if err != nil {
		err = errors.Wrap(err, "unable to read DBConfig")
		log.Fatal(err)
	}

	app.Init(dbConfig)
	app.Run("0.0.0.0:8888")
}
