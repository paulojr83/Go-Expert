## Longest Substring Without Repeating Characters
Data uma string, retorne o tamanho da maior substring que não tenha nenhum caracter repetido.

Exemplo 1:

Entrada: abcabcbb
Saida: 3

Resposta: O valor encontrado é "abc", que tem o tamanho 3.

Exemplo 2:

Entrada: zzzabcdzzz

Saida: 5

Resposta: O valor encontrado é "zabcd", que tem o tamanho 5.



1. Inicializa um mapa lastSeen para armazenar a posição do último caractere visto.
2. Inicializa variáveis start para armazenar o início da substring atual e maxLength para armazenar o tamanho máximo da substring sem caracteres repetidos.
3. Percorre a string s usando um loop for e range.
4. Para cada caractere na string, verifica se ele já foi visto e está dentro da janela atual:
   * Se sim, atualiza o início da substring para a posição após a última ocorrência do caractere.
   * Se não, continua verificando o próximo caractere.

5. Atualiza a posição do último caractere visto no mapa lastSeen.
6. Calcula o tamanho da substring sem caracteres repetidos e atualiza o tamanho máximo, se necessário.
7. Retorna o tamanho máximo da substring sem caracteres repetidos.