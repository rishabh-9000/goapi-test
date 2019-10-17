package helper

import (
	"log"
	"social-api/constants"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT : Generate JSON Web Token
func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email

	tokenString, e := token.SignedString(constants.JWTKey)
	if e != nil {
		log.Println("Something Went Wrong: ", e.Error())
		return "", e
	}

	return tokenString, nil
}
