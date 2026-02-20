package store

import (
	"context"
	"database/sql"
	"time"
)

type OrderRepo struct {
	db *sql.DB
}

type Orders interface {
	Create(ctx context.Context, ProductName string, price int) error
}

type Storage struct {
	Order Orders
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Order: &OrderRepo{db},
	}
}

type PaidStatus string

const (
	Pending PaidStatus = "pending"
	Paid    PaidStatus = "paid"
	Failed  PaidStatus = "failed"
)

type Order struct {
	ID            string
	ProductName      string
	Price         int
	PaymentStatus PaidStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (o *OrderRepo) Create(ctx context.Context, ProductName string, price int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return nil
}
