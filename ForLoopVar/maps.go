package main

import (
	"fmt"
	"strings"
)

func printHasPrefix(prefix string, courses map[uint16]string) {
	for index, course := range courses {
		if strings.HasPrefix(course, prefix) {
			fmt.Println(index, course)
		}
	}
}

func main() {
	courses := map[uint16]string{
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		7: "English",
		8: "Cantonese",
	}

	printHasPrefix("C", courses)
	courses[4] = "Algorithms"
	courses[9] = "Spanish"
	delete(courses, 1)
	fmt.Println(courses)
}
