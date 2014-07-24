package main

import (
// "fmt"
)

type Worker struct {
	requests chan Request // work to do (a buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

func (w *Worker) work(done chan *Worker) {
	go func() {
		for {
			req := <-w.requests // get requests from load balancer
			// fmt.Printf("W: task received.  ")
			req.c <- req.fn() // do the work and send the answer back to the requestor
			done <- w         // tell load balancer a task has been completed by worker w.
			// fmt.Printf("W: done.  ")
		}
	}()
}
