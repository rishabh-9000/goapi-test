package apis

import (
	"context"
	"encoding/json"
	"net/http"
	"social-api/config"
	"social-api/models"
	"time"
)

// CreatePersonEndpoint : Create a new Person
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var person models.User

	_ = json.NewDecoder(r.Body).Decode(&person)
	collection := config.Client.Database("GoTest").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)

	json.NewEncoder(w).Encode(result)
}
