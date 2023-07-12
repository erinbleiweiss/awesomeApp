package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type AppHandler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *AppHandler {
	return &AppHandler{
		db: db,
	}
}

func CreateRouter(db *sql.DB) *mux.Router {
	//create a new router
	router := mux.NewRouter()
	var api = router.PathPrefix("/api").Subrouter()

	// TODO: Add middleware to check sessions and prevent certain routes from being accessible to non-admins

	AddRoutesV1(api.PathPrefix("/v1").Subrouter(), db)

	// Commented-out line below contains example for adding new routes for v2 breaking changes
	// AddRoutesV2(router.PathPrefix("/v2").Subrouter(), db)

	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs", http.FileServer(http.Dir("docs"))))

	return router
}

func GenericResponse(w http.ResponseWriter, data interface{}, status int) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Error marshalling json response ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		log.Fatalln(err)
	}
}

func JSONResponse(w http.ResponseWriter, data interface{}) {
	GenericResponse(w, data, http.StatusOK)
}

func ErrorResponse(w http.ResponseWriter, msg string, errorCode int) {
	err := map[string]string{
		"error":  msg,
		"status": strconv.Itoa(errorCode),
	}
	GenericResponse(w, err, errorCode)
}
