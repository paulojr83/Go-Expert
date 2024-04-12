## Desafio 3
### Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
  Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.

Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.


    $ go run .cmd\ordersystem\main.go .\wire_gen.go



## Migrate
[Github](https://github.com/golang-migrate/migrate?tab=readme-ov-file#cli-usage)

### install windows
    $ scoop install migrate

### Installing Golang Migrate on Ubuntu

1. Let us setup the repository to install the migrate package.

        $ curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
2. Update the system by executing the following command.

        $ sudo apt-get update
3. Execute the following command in the terminal to install migrate.

       $ sudo apt-get install migrate


##
    $ migrate create -ext=sql -dir=sql/migrations -seq init
    $ migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
    $ migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down


### GRPC

[Doc](https://grpc.io/docs/languages/go/quickstart/)

    $ protoc --go_out=. --go-grpc_out=. .\internal\infra\grpc\protofiles\order.proto

Now we run go generate to execute wire:

    $ go generate


## GraphQL

     go run github.com/99designs/gqlgen init
     go run github.com/99designs/gqlgen generate
