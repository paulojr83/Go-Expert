package main

func solution(s string) int {
	// Inicializa um mapa para armazenar a posição do último caractere visto
	lastSeen := make(map[rune]int)
	// Inicializa variáveis para armazenar o início da substring e o tamanho máximo
	start := 0
	maxLength := 0

	// Percorre a string
	for i, char := range s {
		// Se o caractere já foi visto e está dentro da janela atual,
		// atualiza o início da substring para a posição após a última ocorrência do caractere
		if idx, ok := lastSeen[char]; ok && idx >= start {
			start = idx + 1
		}
		// Atualiza a posição do último caractere visto
		lastSeen[char] = i
		// Atualiza o tamanho máximo da substring sem caracteres repetidos
		if length := i - start + 1; length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func main() {
	// Testa a função com exemplos
	examples := []string{"abcabcbb", "zzzabcdzzz"}
	for _, ex := range examples {
		result := solution(ex)
		println("Entrada:", ex, "\nSaída:", result, "\n")
	}
}
