package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msq1 := <-c1:
			println(msq1)
		case msq2 := <-c2:
			println(msq2)
		case <-time.After(time.Second * 3):
			println("Timeout")
		default:
			println("default")
		}
	}
}
