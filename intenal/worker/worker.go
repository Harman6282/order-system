package worker

import (
	"context"
	"log"
)
// this is the worker for the go routines
func (p *Pool) worker(ctx context.Context, id int) {
	defer p.wg.Done()

	log.Printf("worker-%d started\n", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker-%d shutting down\n", id)
			return

		case job, ok := <-p.jobsQueue:
			if !ok {
				return
			}
			log.Printf("worker-%d processing order %s\n", id, job.OrderID)

			err := p.processor.ProcessOrder(ctx, job.OrderID)
			if err != nil {
				log.Println("processing failed:", err)
			}
		}
	}
}
