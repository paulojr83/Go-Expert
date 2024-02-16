package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}

}

// Thread 1
func main() {

	// Thread 2
	go task("A")

	// Thread 3
	go task("B")

	// Thread
	go func() {
		for i := 1; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	//Waiting 15 sec to leave the program
	time.Sleep(15 * time.Second)

}
