package main

import (
	"awesomeProject1/api"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Configure the database connection
	dbConnStr := "sql8645907:N5jq6Eb4Su@tcp(sql8.freemysqlhosting.net:3306)/sql8645907"

	if err := api.InitializeDB(dbConnStr); err != nil {
		panic(err)
	}
	defer api.CloseDB()

	r := mux.NewRouter()

	r.HandleFunc("/signin", api.SigninHandler).Methods("POST")

	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	http.ListenAndServe(port, r)
}
