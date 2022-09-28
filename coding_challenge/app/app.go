package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lazyspell/drone_test/config"
	r "github.com/lazyspell/drone_test/routes"
)

var app config.AppConfig

const portNumber = ":8080"

func Start() {
	fmt.Println("Application has started")

	fmt.Println(fmt.Sprintf("Starting application on http://localhost:8080"))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: r.Routes(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {

	}
	log.Fatal(err)
}
