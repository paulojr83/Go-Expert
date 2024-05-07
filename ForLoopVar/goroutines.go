package main

import (
	"fmt"
	"sync"
	"time"
)

func showMessage(message string, ch chan int) {
	for i := range 5 {
		ch <- i
		time.Sleep(time.Second * 1)
		fmt.Println(i, message)
	}
	close(ch)
}
func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go showMessage("Go is a great programming language", ch)

	go showMessage("Welcome, Gophers!", ch)
	go reader(ch, &wg)

	wg.Wait()
}
func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
		wg.Done()
	}
}
