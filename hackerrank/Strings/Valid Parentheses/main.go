package main

//This function iterates through the characters of the input string 'toValidate',
//using a stack to keep track of opening brackets. For each character:
// If it's an opening bracket, it's pushed onto the stack.
// If it's a closing bracket, it checks if the stack is empty or if the corresponding opening bracket at the top of the stack matches the current closing bracket. If not, the string is invalid.
// If the stack is empty after processing all characters, the string is valid.
// The function returns true if the string is valid and false otherwise.

func solution(toValidate string) bool {
	stack := []rune{} // Using rune for characters to handle unicode characters
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range toValidate {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {

	// Test the function with examples
	examples := []string{"[]", "[()]", "[", "[(", ")[(", "{[]}"}
	for _, ex := range examples {
		result := solution(ex)
		println(result)
	}
}
