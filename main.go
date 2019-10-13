package main

import (
	"fmt"
	"log"
	"net/http"
	"social-api/apis/user"
	"social-api/config"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the application...")

	config.DbConnect()

	router := mux.NewRouter()

	router.HandleFunc("/", Test).Methods("GET")
	router.HandleFunc("/person", user.CreatePersonEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}

// Test : Endpoint to test API is running
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Running.")
}
