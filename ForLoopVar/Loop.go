package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) []string {
	var result []string

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func reduce(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum = number + sum
	}

	return sum
}
func main() {

	fmt.Println(fizzbuzz(15))
	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}
