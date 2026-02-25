package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const DB_URL = "postgres://admin:adminpassword@localhost:5432/orderDB?sslmode=disable"

func NewPostgresPool(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to open connection pool: %v", err)
	}

	db.SetMaxOpenConns(60)
	db.SetMaxIdleConns(60)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to connect to Database: %v", err)
	}

	log.Println("connection pool established")

	return db, nil

}
