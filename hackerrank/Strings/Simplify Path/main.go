package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func solution(path string) string {
	// Remove espaços em branco extras e divide o caminho em seus componentes
	components := strings.Split(path, "/")
	var stack []string

	for _, comp := range components {
		// Se o componente for '.', não faz nada
		if comp == "." || comp == "" {
			continue
		}
		// Se o componente for '..', remove o diretório anterior do stack
		if comp == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			// Caso contrário, empilha o componente
			stack = append(stack, comp)
		}
	}

	// Reconstrói o caminho simplificado a partir dos componentes do stack
	simplifiedPath := "/" + filepath.Join(stack...)
	return simplifiedPath
}

func main() {
	// Testa a função com exemplos
	examples := []string{"/home/", "/home/../", "/home/../home/./", "/home/../home"}
	for _, ex := range examples {
		result := solution(ex)
		fmt.Printf("Entrada: %s\nSaída: %s\n\n", ex, result)
	}
}
