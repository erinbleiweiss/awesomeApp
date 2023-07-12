package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/erinbleiweiss/awesomeApp/backend/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (app *AppHandler) register(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		msg := fmt.Sprintf("Error creating user: invalid input")
		ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if newUser.FirstName == "" ||
		newUser.LastName == "" ||
		newUser.Username == "" ||
		newUser.Email == "" ||
		newUser.Password == "" {
		ErrorResponse(w, "Fields cannot be blank", http.StatusBadRequest)
		return
	}

	// TODO: Server-side field validation, including valid email format
	// TODO: Ideally, password requirements should be enforced by a separate validator function
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 8)
	if err != nil {
		ErrorResponse(w, "Invalid password", http.StatusBadRequest)
		return
	}
	newUser.Password = string(hashedPassword)

	// Set all roles to "user" by default (we're assuming there are other privileged roles like "admin"
	newUser.Role = "user"

	if err := newUser.CreateUser(app.db); err != nil {
		msg := fmt.Sprintf("Error creating user: %v", err)
		ErrorResponse(w, msg, http.StatusInternalServerError)
		return
	}

	// Don't include hashed password in json response
	newUser.Password = ""

	JSONResponse(w, newUser)
}

func (app *AppHandler) login(w http.ResponseWriter, r *http.Request) {
	// In this simple app, only admins are able to login
	// Any requests to users without an "admin" role will be denied.

	// TODO: A real app would implement something like basic auth and issue session tokens
	type loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	data := loginData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		msg := fmt.Sprintf("login error: invalid input")
		ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	user := model.User{Username: data.Username}
	if err := user.GetUserByUsername(app.db); err != nil {
		switch err {
		case sql.ErrNoRows:
			msg := fmt.Sprintf("username not found")
			ErrorResponse(w, msg, http.StatusNotFound)
			return
		default:
			msg := fmt.Sprintf("Error querying user %s: %v", data.Username, err)
			ErrorResponse(w, msg, http.StatusInternalServerError)
			return
		}
	}

	if user.Role != "admin" {
		ErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	authError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if authError != nil {
		ErrorResponse(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	response := map[string]string{
		"status": "success",
	}
	JSONResponse(w, response)
}
