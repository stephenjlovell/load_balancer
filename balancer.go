package main

import (
	"container/heap"
	"fmt"
)

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) start() {
	for _, worker := range b.pool {
		worker.work(b.done)
	}
}

func (b *Balancer) balance(work chan Request) {

	go func() {
		for {
			// time.Sleep(time.Millisecond * time.Duration(rand.Int63n(100)))
			select {
			case req := <-work: // request received
				b.dispatch(req) // send request to a worker
			case w := <-b.done: // worker finished with a request
				b.completed(w) //
			}
		}
	}()

}

// route the request to the most lightly loaded worker in the priority queue.

func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Worker)
	w.requests <- req
	w.pending++
	heap.Push(&b.pool, w)
}

func (b *Balancer) completed(w *Worker) {
	w.pending--
	heap.Remove(&b.pool, w.index)
	heap.Push(&b.pool, w)
}

func (b *Balancer) print() {
	fmt.Printf("\n")
	total_pending := 0
	for _, worker := range b.pool {
		pending := worker.pending
		fmt.Printf("%d  ", pending)
		total_pending += pending
	}
	fmt.Printf("| %d  ", total_pending)
}

func new_balancer(nworker int, work chan Request) *Balancer {
	b := &Balancer{
		done: make(chan *Worker, 100),
		pool: make(Pool, nworker),
	}
	for i := 0; i < nworker; i++ {
		b.pool[i] = &Worker{
			requests: make(chan Request, 100), // each worker needs its own channel on which to receive work
			index:    i,                       // from the load balancer.
		}
	}
	heap.Init(&b.pool)

	return b
}
