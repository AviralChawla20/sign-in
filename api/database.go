package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // Database connection object

// User represents a user record in the database
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// InitializeDB initializes the database connection
func InitializeDB(dbConnStr string) error {
	var err error
	db, err = sql.Open("mysql", dbConnStr)
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// ErrorResponse represents an error response in JSON format
type ErrorResponse struct {
	Message string `json:"message"`
}

// FindUserByEmail finds a user by email
func FindUserByEmail(email string) (*User, error) {
	var user User
	query := "SELECT email, password FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		// If the user is not found, return an error response in JSON format
		errorResponse := ErrorResponse{Message: "User not found"}
		errJSON, _ := json.Marshal(errorResponse)
		return nil, errors.New(string(errJSON))
	}
	return &user, nil
}

// Add more database-related functions as needed
