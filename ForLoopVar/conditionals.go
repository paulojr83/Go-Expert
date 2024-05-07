package main

import "fmt"

func main() {
	n1 := 10

	if n1 < 0 {
		fmt.Println(n1, "is negative")
	} else if n1 < 100 {
		fmt.Println(n1, "is positive")
	} else {
		fmt.Println(n1, "is positive and is a large number!")
	}
}
