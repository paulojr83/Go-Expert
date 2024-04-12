package main

import "sync"

func main() {
	ch := make(chan string, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch <- "Hello 1"
	wg.Done()
	ch <- "Hello 2"
	wg.Done()
	println(<-ch)
	println(<-ch)
	wg.Wait()
}
