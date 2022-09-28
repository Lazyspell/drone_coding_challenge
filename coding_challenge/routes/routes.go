package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/lazyspell/drone_test/config"
	"github.com/lazyspell/drone_test/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Group(func(drone chi.Router) {
		drone.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte("Authorization"), nil)))
		drone.Use(jwtauth.Authenticator)

		drone.Get("/drone/destination", handlers.NewDestination)

	})

	mux.Post("/login", handlers.LoginUser)

	return mux
}
