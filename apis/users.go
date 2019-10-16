package apis

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"social-api/config"
	"social-api/models"
	"social-api/validators"
	"time"
)

// UserEndpoint : Create a new User
func UserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	collection := config.Client.Database("test").Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if r.Method == http.MethodPost {
		var user models.User
		var err models.Err
		var errors []models.Err

		_ = json.NewDecoder(r.Body).Decode(&user)

		// Validation
		name := user.Name
		if !validators.IsEmpty(name) {
			err.Message = "Name is Empty."
			errors = append(errors, err)
		}
		if !validators.IsChar(name) {
			err.Message = "Invalid Characters in Name."
			errors = append(errors, err)
		}

		email := user.Email
		if !validators.IsEmail(email) {
			err.Message = "Invalid Email."
			errors = append(errors, err)
		}

		password := user.Password
		if !validators.IsEmpty(password) {
			err.Message = "Password cannot be empty."
			errors = append(errors, err)
		}

		// See if user exists

		// Encrypt Password

		// Return JWT
		log.Println(user)
		if len(errors) != 0 {
			result, e := collection.InsertOne(ctx, user)
			if e != nil {
				log.Fatal(e)
				json.NewEncoder(w).Encode(e)
			}
			json.NewEncoder(w).Encode(result)
		} else {
			result := "Internal Server Error"
			json.NewEncoder(w).Encode(result)
		}
	}
}
