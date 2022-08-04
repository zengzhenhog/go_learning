package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, c chan int, done func()) {
	for n := range c {
		fmt.Printf("Worker id %d, Received %c\n", id, n)
		done()
	}
}

// func doWorker(id int, c chan int, done chan bool) {
// 	for n := range c {
// 		fmt.Printf("Worker id %d, Received %c\n", id, n)
// 		// go func() {done <- true}()
// 		done <- true
// 	}
// }

type Worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var channels [10]Worker

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range channels {
		worker.in <- 'a' + i
	}
	for i, worker := range channels {
		worker.in <- 'A' + i
	}

	wg.Wait()

	// for i, worker := range channels {
	// 	worker.in <- 'a' + i
	// }
	// for _, worker := range channels {
	// 	<-worker.done
	// }
	// for i, worker := range channels {
	// 	worker.in <- 'A' + i
	// }
	// for _, worker := range channels {
	// 	<-worker.done
	// }

	// go func() {done <- true}() doWorker里结合并发发送，或者按上面的方式先接收小写再接收大写
	// for _, worker := range channels {
	// 	<-worker.done
	// 	<-worker.done
	// }
}

func main() {
	chanDemo()
}
