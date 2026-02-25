package store

import (
	"context"
	"log"
	"time"
)

// type ProcessorRepo interface {
// 	ClaimOrder(ctx context.Context, orderID string) error
// 	CompleteOrder(ctx context.Context, orderID string) error
// }

type Processor struct {
	repo Orders
}

func NewProcessor(r Orders) *Processor {
	return &Processor{repo: r}
}

func (p *Processor) ProcessOrder(ctx context.Context, orderID string) error {
	// claim order and lock row

	err := p.repo.ClaimOrder(ctx, orderID)
	if err != nil {
		return err
	}

	log.Println("processing order: ", orderID)

	// simulate work
	time.Sleep(2 * time.Second)

	// complete order

	return p.repo.CompleteOrder(ctx, orderID)
}
