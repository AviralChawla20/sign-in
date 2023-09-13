package api

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// SigninHandler handles user sign-in
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var inputUser User
	err := json.NewDecoder(r.Body).Decode(&inputUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query the database to find the user by email
	user, err := FindUserByEmail(inputUser.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Respond with success message
	response := map[string]interface{}{"message": "Sign in successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Define other route handlers as needed
