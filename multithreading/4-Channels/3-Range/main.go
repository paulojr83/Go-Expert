package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
	}
}
func publish(ch chan int) {
	for i := 1; i < 11; i++ {
		ch <- i
	}
	close(ch) // fechar o canal
}
