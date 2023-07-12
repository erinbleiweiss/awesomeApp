package api

import (
	"database/sql"
	"fmt"
	"github.com/erinbleiweiss/awesomeApp/backend/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *AppHandler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ErrorResponse(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user := model.User{ID: id}
	if err := user.GetUserById(app.db); err != nil {
		switch err {
		case sql.ErrNoRows:
			msg := fmt.Sprintf("user ID %v not found", id)
			ErrorResponse(w, msg, http.StatusNotFound)
			return
		default:
			msg := fmt.Sprintf("could not get user ID %v: %v", id, err)
			ErrorResponse(w, msg, http.StatusInternalServerError)
			return
		}
	}

	JSONResponse(w, user)
}

func (app *AppHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement pagination to handle a database at scale
	users, err := model.GetUsers(app.db)
	if err != nil {
		msg := fmt.Sprintf("error getting users: %v", err)
		ErrorResponse(w, msg, http.StatusInternalServerError)
		return
	}

	JSONResponse(w, users)
}

// TODO: Implement Update/Delete methods. They are not implemented here for the sake of brevity
