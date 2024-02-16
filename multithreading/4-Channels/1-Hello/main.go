package main

import "fmt"

/**
* Fazer comunicação entre theads
* Segurança para um thread saber o momento em que ela pode trabalhar com um determinado dado
**/

// Thread 1
func main() {

	channel := make(chan string) // Vazio
	// Thread 2
	go func() { // Está cheio
		channel <- "Hello world!"
	}()

	// Thread 1
	msg := <-channel // Canal esvazia
	fmt.Println(msg)

}
