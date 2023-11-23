package main

import (
	"awesomeProject1/api"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func initializeRouter() {
	r := mux.NewRouter()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	r.HandleFunc("/signin", api.SigninHandler).Methods("POST")

	http.Handle("/", corsMiddleware(r))
}

func main() {
	dbConnStr := "sql12647981:XM51KVKzDA@tcp(sql12.freemysqlhosting.net:3306)/sql12647981"

	if err := api.InitializeDB(dbConnStr); err != nil {
		panic(err)
	}
	defer api.CloseDB()

	initializeRouter()

	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	http.ListenAndServe(port, nil)
}
