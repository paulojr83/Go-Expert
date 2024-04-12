package main

import (
	"strconv"
	"strings"
)

func solution(toDecode string) string {
	stack := []string{}
	for i := 0; i < len(toDecode); i++ {
		if toDecode[i] == ']' {
			decoded := ""
			// Desempilhe até encontrar '['
			for stack[len(stack)-1] != "[" {
				decoded = stack[len(stack)-1] + decoded
				stack = stack[:len(stack)-1]
			}
			// Remova '['
			stack = stack[:len(stack)-1]
			k := ""
			// Obtenha o número de repetições
			for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				k = stack[len(stack)-1] + k
				stack = stack[:len(stack)-1]
			}
			// Converta k para inteiro
			numRepeats, _ := strconv.Atoi(k)
			// Repita a string decodificada e empilhe de volta
			decodedRepeated := strings.Repeat(decoded, numRepeats)
			stack = append(stack, decodedRepeated)
		} else {
			// Empilhe caracteres normais
			stack = append(stack, string(toDecode[i]))
		}
	}
	// Concatene todas as strings na pilha
	return strings.Join(stack, "")
}

func main() {
	// Teste da função com exemplos
	examples := []string{"2[a]3[bc]", "3[a2[c]]", "2[abc]3[cd]ef"}
	for _, ex := range examples {
		result := solution(ex)
		println(result)
	}
}
