package model

import (
	"database/sql"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type User struct {
	// TODO: Ideally UUIDs would be used instead of auto-incrementing numeric IDs
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
}

func (u *User) GetUserById(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, first_name, last_name, email, username, password, role FROM user where id=$1",
		u.ID,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Username,
		&u.Password,
		&u.Role)
}

func (u *User) GetUserByUsername(db *sql.DB) error {
	return db.QueryRow(
		"SELECT id, first_name, last_name, email, username, password, role FROM user where username=$1",
		u.Username,
	).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Username,
		&u.Password,
		&u.Role)
}

func GetUsers(db *sql.DB) ([]User, error) {
	// TODO: Implement pagination to handle a database at scale
	rows, err := db.Query("SELECT id, first_name, last_name, email, username, role FROM user")
	if err != nil {
		return nil, errors.Wrap(err, "error fetching users")
	}

	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Role); err != nil {
			return nil, errors.Wrap(err, "error fetching users")
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) CreateUser(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO user(first_name, last_name, email, username, password, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		u.FirstName,
		u.LastName,
		u.Email,
		u.Username,
		u.Password,
		u.Role,
	).Scan(&u.ID)
	if err != nil {
		if sqlErr, ok := err.(sqlite3.Error); ok {
			if sqlErr.Code == sqlite3.ErrConstraint {
				return errors.New("Username or email already exists")
			}
		}
		return errors.Wrap(err, fmt.Sprintf("error creating user: %s", u.Username))
	}
	return nil
}

// TODO: Implement Update/Delete methods. They are not implemented here for the sake of brevity
