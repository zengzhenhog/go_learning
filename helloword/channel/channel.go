package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		// n := <-c
		// n, ok := <-c
		// if !ok {
		// 	break
		// }
		fmt.Printf("Worker id %d, Received %c\n", id, n)
		// fmt.Printf("Worker id %d, Received %c\n", id, <-c)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedAndCloseChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	bufferedAndCloseChannel()
	// chanDemo()
	// fmt.Println('a', "a") // 97 a
}

// func worker(id int, c chan int) {
// 	for {
// 		// n := <-c
// 		fmt.Printf("Worker id %d, Received %d\n", id, <-c)
// 	}
// }

// func chanDemo() {
// 	c := make(chan int)
// 	go worker(1, c)

// 	c <- 1
// 	c <- 2
// 	time.Sleep(time.Millisecond)
// }

// func worker(id int, c chan int) {
// 	for {
// 		// n := <-c
// 		fmt.Printf("Worker id %d, Received %c\n", id, <-c)
// 	}
// }

// func chanDemo() {
// 	var channels [10]chan int

// 	for i := 0; i < 10; i++ {
// 		channels[i] = make(chan int)
// 		go worker(i, channels[i])
// 	}

// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'a' + i
// 	}
// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'A' + i
// 	}

// 	time.Sleep(time.Millisecond)
// }
