package api

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"sync"
)

// SigninHandler handles user sign-in
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

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
			// If the user is not found, return an error response in JSON format
			errorResponse := ErrorResponse{Message: "User not found"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		// Compare passwords
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password))
		if err != nil {
			// If the password is invalid, return an error response in JSON format
			errorResponse := ErrorResponse{Message: "Invalid password"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		// Respond with success message
		response := map[string]interface{}{"message": "Sign in successful"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}()

	wg.Wait()
}
