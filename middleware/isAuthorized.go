package middleware

import (
	"fmt"
	"log"
	"net/http"
	"social-api/constants"

	jwt "github.com/dgrijalva/jwt-go"
)

// IsAuthorized : Authorized route middleware
func IsAuthorized(endpoint http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["X-Auth-Token"] != nil {
			token, e := jwt.Parse(r.Header["X-Auth-Token"][0], func(token *jwt.Token) (interface{}, error) {
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
