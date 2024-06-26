package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	/*//Create a new product repository
	repository := product.NewProductRepository(db)

	//Create a new product usecase
	usecase := product.NewProductUseCase(repository)
	*/
	usecase := NewProductUseCase(db)
	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
