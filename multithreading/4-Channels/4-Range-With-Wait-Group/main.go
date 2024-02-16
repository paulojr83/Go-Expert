package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	// Thread 2
	go publish(ch)

	// Thread 3
	go reader(ch, &wg)

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
		wg.Done()
	}
}
func publish(ch chan int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}
	close(ch) // fechar o canal
}
