package utils

import (
	"log"

	"github.com/go-chi/jwtauth"
)

func jwtToken() (string, error) {
	tokenAuth := jwtauth.New("HS256", []byte("Authorization"), nil)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{})
	if err != nil {
		log.Println("issue generating the token")
	}

	return tokenString, nil

}
