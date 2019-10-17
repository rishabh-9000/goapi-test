package main

import (
	"fmt"
	"log"
	"net/http"
	"social-api/apis"
	"social-api/config"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting the application...")

	config.DbConnect()

	router := mux.NewRouter()

	// Test Route
	router.HandleFunc(
		"/",
		Test).Methods("GET")

	// Users Route
	router.HandleFunc(
		"/api/users",
		apis.UserEndpoint).Methods("GET", "POST")

	// Login Route
	router.HandleFunc(
		"/api/login",
		apis.UserLogin).Methods("POST")

	// // Profile Route
	// router.HandleFunc(
	// 	"/api/profile",
	// 	apis.CreatePersonEndpoint).Methods("POST")

	// // Posts Route
	// router.HandleFunc(
	// 	"/api/posts",
	// 	apis.CreatePersonEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}

// Test : Endpoint to test API is running
func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Running. And method is: ")
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, r.Method)
	}
}
