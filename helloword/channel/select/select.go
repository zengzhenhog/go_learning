package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker id %d, Received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	// var c1, c2 chan int // c1 and c2 = nil
	var c1, c2 = generator(), generator()
	worker := createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeOut")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}

// func main() {
// 	// var c1, c2 chan int // c1 and c2 = nil
// 	var c1, c2 = generator(), generator()
// 	worker := createWorker(0)

// 	n := 0
// 	hasValue := false
// 	for {
// 		var activeWorker chan<- int
// 		if hasValue {
// 			activeWorker = worker
// 		}
// 		select {
// 		case n = <-c1:
// 			hasValue = true
// 		case n = <-c2:
// 			hasValue = true
// 		case activeWorker <- n:
// 			hasValue = false
// 		}
// 	}

// 	// 接收一次
// 	// select {
// 	// case n := <-c1:
// 	// 	fmt.Println("Received from c1:", n)
// 	// case n := <-c2:
// 	// 	fmt.Println("Received from c2:", n)
// 	// default:
// 	// 	fmt.Println("No value received")
// 	// }
// }
