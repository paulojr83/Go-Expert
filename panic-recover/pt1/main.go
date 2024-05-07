package main

import (
	"fmt"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			if r == "Something went wrong! Panic1" {
				fmt.Println("Recovered in main: ", r)
			}
			if r == "Something went wrong! Panic2" {
				fmt.Println("Recovered in main: ", r)
			}

		}
	}()
	myPanic1()
	myPanic2()
}

func myPanic1() {
	panic("Something went wrong! Panic1")
}

func myPanic2() {
	panic("Something went wrong! Panic2")
}
