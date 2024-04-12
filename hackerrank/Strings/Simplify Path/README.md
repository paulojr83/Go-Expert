## Simplify Path
Simplifique o path absoluto para um arquivo (no estilo Unix). Em outras palavras, converta o mesmo para o modo canonico.

Neste modo, o '.' se refere ao diretório atual. '..' se refere ao diretório acima.

Lembre-se que o path neste formato deve sempre começar com '/' e sempre devera ter um '/' único entre os diretórios.

Exemplo 1:

Entrada: "/home/"
Saída: "/home"

Exemplo 2:

Entrada: "/home/../"
Saída: "/"

Exemplo 3:

Entrada: "/home/../home/./"
Saída: "/home"

Exemplo 4:

Entrada: "/home/../home"
Saída: "/home"



1. Remove espaços em branco extras da entrada e divide o caminho em seus componentes.
2. Cria um slice vazio stack para manter os componentes do caminho simplificado.
3. Percorre os componentes do caminho:
   * Se o componente for '.', não faz nada.
   * Se o componente for '..', remove o diretório anterior do stack.
   * Caso contrário, empilha o componente no stack.
4. Reconstrói o caminho simplificado a partir dos componentes do stack.
5. Retorna o caminho simplificado.