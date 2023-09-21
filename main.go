package main

import (
	"awesomeProject1/api"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Configure the database connection
	dbConnStr := "sql12647981:XM51KVKzDA@tcp(sql12.freemysqlhosting.net:3306)/sql12647981"

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
