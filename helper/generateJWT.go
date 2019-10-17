package helper

import (
	"log"
	"social-api/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT : Generate JSON Web Token
func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, e := token.SignedString(constants.JWTKey)
	if e != nil {
		log.Println("Something Went Wrong: ", e.Error())
		return "", e
	}

	return tokenString, nil
}
