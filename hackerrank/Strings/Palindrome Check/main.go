package main

import (
	"fmt"
)

// solution verifica se a string é uma palindrome
func solution(s string) bool {
	// Inicializa os índices de início e fim
	start := 0
	end := len(s) - 1

	// Percorre a string até os índices se cruzarem
	for start < end {
		// Se os caracteres nas posições start e end forem diferentes, a string não é uma palindrome
		if s[start] != s[end] {
			return false
		}
		// Avança start e retrocede end
		start++
		end--
	}
	// Se o loop terminar sem retornar falso, a string é uma palindrome
	return true
}

func main() {
	examples := []string{"algomania", "abccba", "hello", "radar", "level"}
	for _, ex := range examples {
		result := solution(ex)
		fmt.Printf("'%s' é uma palindrome? %t\n", ex, result)
	}
}
