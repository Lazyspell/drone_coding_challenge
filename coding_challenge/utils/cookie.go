package utils

import (
	"log"
	"net/http"
	"time"
)

func GenerateStateJwtCookie(w http.ResponseWriter) string {

	var expiration = time.Now().Add(60 * time.Minute)

	state, err := jwtToken()
	if err != nil {
		log.Println(err)
		return state
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    state,
		HttpOnly: true,
		Expires:  expiration,
	}
	http.SetCookie(w, &cookie)

	return state

}
