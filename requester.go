package main

import (
	"fmt"
	"math/rand"
	"time"
	// "runtime"
)

var counter, other_counter int

type Request struct {
	fn func() int // operation to perform
	c  chan int   // channel on which to return result
}

func requester(work chan Request) {
	c := make(chan int, 100)
	// fmt.Printf("R.")
	go func() {
		for {
			time.Sleep(time.Microsecond * time.Duration(rand.Int63n(10))) // simulate uneven throughput

			work <- Request{do_some_work, c} // send a work request
			// fmt.Println("R: work requested.")
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
	other_counter++
	// do very important work.
}

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(8)
	nworkers := 8
	work := make(chan Request, 100)

	balancer := new_balancer(nworkers, work)
	balancer.start()
	balancer.balance(work)

	for i := 1; i <= 100; i++ {
		requester(work)
	}

	go func() {
		for _ = range time.Tick(250 * time.Millisecond) {
			balancer.print()
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Printf("\n %d/%d jobs complete.\n", counter, other_counter)
}
