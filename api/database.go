package api

import (
	"database/sql"
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

// FindUserByEmail finds a user by email
func FindUserByEmail(email string) (*User, error) {
	var user User
	query := "SELECT email, password FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return &user, nil
}

// Add more database-related functions as needed
