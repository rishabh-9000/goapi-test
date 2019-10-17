package middleware

import (
	"fmt"
	"log"
	"net/http"
	"social-api/constants"

	jwt "github.com/dgrijalva/jwt-go"
)

// IsAuthorized : Authorized route middleware
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["x-auth-token"] != nil {
			token, e := jwt.Parse(r.Header["x-auth-token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return constants.JWTKey, nil
			})

			if e != nil {
				log.Println(w, e.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			log.Println(w, "Not Authorized")
		}
	})
}
