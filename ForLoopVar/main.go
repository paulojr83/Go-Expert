package main

import "fmt"

func main() {

	var language = "Go"
	const released = 2009
	isAProgrammingLanguage := true

	fmt.Printf("language: %s\nreleased: %d\nisAProgrammingLanguage:%v\n", language, released, isAProgrammingLanguage)

	for i := range 10 {
		fmt.Println(i)
	}
	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, value := range values {
		go func() {
			fmt.Println(value)
			done <- true
		}()
	}

	for range values {
		<-done
	}

	number := 8
	numberText := "eight"

	if number == 8 && numberText == "eight" {
		fmt.Println("Success!")
	} else {
		fmt.Println("Fail!")
	}
}
