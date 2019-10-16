package apis

import (
	"context"
	"encoding/json"
	"net/http"
	"social-api/config"
	"social-api/models"
	"social-api/validators"
	"time"

	"github.com/gorilla/mux"
)

// UserEndpoint : Create a new User
func UserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	collection := config.Client.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if r.Method == http.MethodPost {
		// Validation
		valid := true
		vars := mux.Vars(r)

		name := vars["name"]

		valid = validators.IsEmpty(name) && validators.IsChar(name)

		// See if user exists

		// Encrypt Password

		// Return JWT

		// result, _ := collection.InsertOne(ctx, user)
		// json.NewEncoder(w).Encode(result)
	}
}
