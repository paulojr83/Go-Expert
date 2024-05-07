package main

import (
	"fmt"
	"strings"
)

func add(n1, n2 float64) float64 {
	return n1 + n2
}

func sayHello(name string) string {
	return "Hello " + name
}

func sayLoudly(phrase string) string {
	return strings.ToUpper(phrase)
}

func getRectangleArea(width, length int) string {
	product := width * length

	if product < 50 {
		return fmt.Sprintf("The area is %d, which is less than 50", product)
	} else {
		return fmt.Sprintf("The area is %d, which is greater than or equal to 50", product)
	}
}
func main() {

	fmt.Println(add(2.2, 54.3))
	fmt.Println(sayHello("Paulo"))
	fmt.Println(sayLoudly(sayHello("Paulo")))

	fmt.Println(getRectangleArea(5, 15))
}
