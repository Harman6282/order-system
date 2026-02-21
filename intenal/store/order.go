package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type OrderRepo struct {
	db *sql.DB
}

type Orders interface {
	Create(ctx context.Context, id, productName string, price int) (*Order, error)
	Pay(ctx context.Context, id string) (*payRes, error)
	GetStatus(ctx context.Context, id string) (OrderStatus, error)
}

type Storage struct {
	Order Orders
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Order: &OrderRepo{db},
	}
}

type OrderStatus string

const (
	CREATED    OrderStatus = "created"
	PAID       OrderStatus = "paid"
	PROCESSING OrderStatus = "processing"
	COMPLETED  OrderStatus = "completed"
	FAILED     OrderStatus = "failed"
)

type Order struct {
	ID                  string
	ProductName         string
	Price               int
	Status              OrderStatus
	ProcessingBy        *string
	ProcessingStartedAt *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type payRes struct {
	ID          string
	ProductName string
	Price       int
	Status      OrderStatus
}

func (s *OrderRepo) Create(ctx context.Context, id, productName string, price int) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO orders (id, product_name, price)
				VALUES ($1, $2, $3)
				RETURNING id, product_name, price, status, processing_by, processing_started_at, created_at, updated_at`

	var order Order
	err := s.db.QueryRowContext(ctx, query, id, productName, price).Scan(
		&order.ID,
		&order.ProductName,
		&order.Price,
		&order.Status,
		&order.ProcessingBy,
		&order.ProcessingStartedAt,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *OrderRepo) Pay(ctx context.Context, id string) (*payRes, error) {

	var order payRes

	query := `UPDATE orders
			  SET status = 'paid'
			  WHERE id = $1 
			  RETURNING id, product_name, price, status`

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&order.ID,
		&order.ProductName,
		&order.Price,
		&order.Status,
	)

	if err != nil {
		return nil, err
	}

	return &order, nil

}

func (s *OrderRepo) GetStatus(ctx context.Context, id string) (OrderStatus, error) {

	query := `SELECT status FROM orders WHERE id = $1`

	var status OrderStatus

	err := s.db.QueryRowContext(ctx, query, id).Scan(&status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrOrderNotFound
		}
		return "", err
	}

	return status, nil
}
