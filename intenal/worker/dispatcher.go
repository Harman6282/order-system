package worker

type Dispatcher struct {
	pool *Pool
}

func NewDispatcher(pool *Pool) *Dispatcher {
	return &Dispatcher{pool: pool}
}


func (d *Dispatcher) Enqueue(orderId string) {
	d.pool.Enqueue(orderId)
}