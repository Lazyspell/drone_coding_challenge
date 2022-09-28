package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lazyspell/drone_test/utils"
)

//when drone has login a jwt token will be sent out allowing them to access endpoint.
func LoginUser(w http.ResponseWriter, r *http.Request) {

	utils.GenerateStateJwtCookie(w)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("drone auth")
}
