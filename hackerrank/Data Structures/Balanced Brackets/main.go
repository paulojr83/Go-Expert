package main

import (
	"fmt"
)

func isBalanced(s string) string {
	// Create a stack to store opening brackets
	var stack []rune

	// Define a mapping of opening to closing brackets
	bracketMap := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	// Iterate through the string
	for _, char := range s {
		// If the current character is an opening bracket, push it onto the stack
		if char == '{' || char == '[' || char == '(' {
			stack = append(stack, char)
		} else {
			// If the current character is a closing bracket
			// Check if the stack is empty or the top element does not match the closing bracket
			if len(stack) == 0 || bracketMap[stack[len(stack)-1]] != char {
				return "NO"
			}
			// Pop the top element from the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is empty, all brackets are matched
	if len(stack) == 0 {
		return "YES"
	}
	// If the stack is not empty, there are unmatched opening brackets
	return "NO"
}

func main() {
	// Sample input strings
	input := []string{"{[()]}", "{[(])}", "{{[[(())]]}}"}

	// Check if each string is balanced and print the result
	for _, s := range input {
		fmt.Println(isBalanced(s))
	}
}
