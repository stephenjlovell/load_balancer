package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
)

var counter int

type Request struct {
	fn func() int // operation to perform
	c  chan int   // channel on which to return result
}

func requester(work chan Request) {
	c := make(chan int, 100)
	go func() {
		for {
			time.Sleep(time.Microsecond * time.Duration(rand.Int63n(10))) // simulate uneven throughput
			work <- Request{do_some_work, c} // send a work request
			result := <-c // wait for answer
			do_something_else(result)
		}
	}()
}

func do_some_work() int {
	counter++
  time.Sleep(time.Microsecond * time.Duration(rand.Int63n(10)))  // simulate some actual work.
	return counter
}

func do_something_else(r int) {
	// do very important work.
}

func main() {
	procs := runtime.NumCPU()
	runtime.GOMAXPROCS(procs)
	fmt.Printf("\nRunning on %d processors.\n", procs)
	nworkers := procs-2

	rand.Seed(8)

	work := make(chan Request, 100)

	balancer := new_balancer(nworkers, work)
	balancer.start()
	balancer.balance(work)

	for i := 1; i <= 100; i++ {
		requester(work)
	}

	go func() {
		for _ = range time.Tick(250 * time.Millisecond) {  
			balancer.print()  // periodically print out the number of pending tasks assigned to each worker.
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Printf("\n %d jobs complete.\n", counter)
}
