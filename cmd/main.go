package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	r.Get("/", app.health)

	fmt.Println("server started at :8080")
	err := http.ListenAndServe(app.config.addr, r)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}
