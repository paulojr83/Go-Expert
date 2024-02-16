package main

import "fmt"

// Thread 1
func main() {
	hello := make(chan string)

	go recebe("Hello", hello)
	ler(hello)

}

// Invalid operation: hello <- nome (send to the receive-only type <-chan chan string)
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// Invalid operation: <-data (receive from the send-only type chan<- string)
func ler(data chan string) {
	fmt.Println(<-data)
}
