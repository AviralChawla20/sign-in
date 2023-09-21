package main

import (
	"awesomeProject1/api"
	"fmt"
	_ "github.com/SparkPost/gosparkpost"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Configure the database connection
	//apiKey := "84000f42afe2ecf004f7769ccd491209145bdc42"
	dbConnStr := "sql8645907:N5jq6Eb4Su@tcp(sql8.freemysqlhosting.net:3306)/sql8645907"

	if err := api.InitializeDB(dbConnStr); err != nil {
		panic(err)
	}
	defer api.CloseDB()
	// Define CORS options to allow any origin, methods, and headers
	//corsOptions := handlers.AllowedOrigins([]string{"*"})

	r := mux.NewRouter()

	// Define CORS options to allow specific origins, methods, and headers
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Apply the CORS middleware to your router with the defined options
	cors := handlers.CORS(
		allowedOrigins,
		allowedMethods,
		allowedHeaders,
		handlers.ExposedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Define your API routes
	// r.HandleFunc("/send-email", sendEmailHandler).Methods("POST")
	r.HandleFunc("/signin", api.SigninHandler).Methods("POST")

	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	http.ListenAndServe(port, cors)

}
