package worker

import (
	"context"
	"sync"
)

type Job struct {
	OrderID string
}

type Processor interface {
	ProcessOrder(ctx context.Context, orderID string) error
}

type Pool struct {
	jobsQueue chan Job
	workers int
	processor Processor
	wg sync.WaitGroup
}

func NewPool(workerCount int, processor Processor) *Pool {
	return &Pool{
		jobsQueue: make(chan Job, 20),
		workers: workerCount,
		processor: processor,
	}
}


func (p *Pool) Start(ctx context.Context) {
	for i := 1; i <= p.workers; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}
}

func (p *Pool) Stop() {
	close(p.jobsQueue)
	p.wg.Wait()
}

func (p *Pool) Enqueue(orderID string) {
	p.jobsQueue <- Job{OrderID: orderID}
}