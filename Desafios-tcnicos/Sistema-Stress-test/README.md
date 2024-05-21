# Load Tester CLI

O Load Tester CLI é uma ferramenta desenvolvida em Go para realizar testes de carga em serviços web. Ele permite que você forneça uma URL, o número total de requisições e a quantidade de chamadas simultâneas, e então gera um relatório detalhado após a execução dos testes.

## Funcionalidades

- Realiza requisições HTTP para a URL especificada.
- Distribui as requisições de acordo com o nível de concorrência definido.
- Gera um relatório com:
    - Tempo total gasto na execução.
    - Quantidade total de requisições realizadas.
    - Quantidade de requisições com status HTTP 200.
    - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Pré-requisitos

- Docker instalado.
- Go (opcional, caso queira rodar localmente sem Docker).

## Como usar sem o Docker
* Code 200

   ```sh
   go run .\cmd\main.go --url=http://google.com --requests=10 --concurrency=10    
   ```
![img_2.png](img_2.png)

* Code 404
   ```sh
   go run .\cmd\main.go --url=http://google.com/404 --requests=10 --concurrency=10    
   ```
![img.png](img.png)

### Usando Docker

1. **Construa a imagem Docker:**

   ```sh
   docker build -t load-tester .
   ```

2. Execute o contêiner Docker:

    ```sh
    docker run --rm load-tester --url=http://google.com --requests=10 --concurrency=10
    ```
![img_1.png](img_1.png)