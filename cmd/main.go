package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Harman6282/order-system/intenal/database"
	"github.com/Harman6282/order-system/intenal/store"
	"github.com/Harman6282/order-system/intenal/worker"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  *store.Storage
	pool *worker.Pool
}

type config struct {
	addr string
}

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := database.NewPostgresPool(ctx)
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cfg := config{
		addr: ":8080",
	}

	storage := store.NewStorage(db)

	processor := store.NewProcessor(storage.Order)
    pool := worker.NewPool(5, processor)

	
	app := &application{
		config: cfg,
		store:  store.NewStorage(db),
		pool: pool,
	}
	
	pool.Start(ctx)
	
	r.Get("/", app.health)
	r.Post("/create-order", app.createOrder)
	r.Patch("/pay/{id}", app.payOrder)
	r.Get("/status/{id}", app.checkStatus)

	fmt.Println("server started at :8080")

	err = http.ListenAndServe(app.config.addr, r)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}
