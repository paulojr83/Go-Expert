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
    $ migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
    $ migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down


### [SQLX](https://pkg.go.dev/github.com/jmoiron/sqlx#readme-install)
* https://pkg.go.dev/github.com/jmoiron/sqlx#readme-install
* https://github.com/jmoiron/sqlx

### [SQLC](https://sqlc.dev)
* https://docs.sqlc.dev/en/stable/overview/install.html
* https://github.com/sqlc-dev/sqlc

      $ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    