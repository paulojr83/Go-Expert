package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // _ black identify, ignore no use import
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name =?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ? ")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	var p Product
	err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	//err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	product := NewProduct("Notebook 2", 1990.90)
	println(product)

	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 3000.90
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	p, err := selectProduct(ctx, db, "9996927b-da11-43d4-af33-6bc7654e4760")
	fmt.Printf("Product: %v, possui o preço de R$%.2f", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de R$%.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, "9996927b-da11-43d4-af33-6bc7654e4760")
	if err != nil {
		panic(err)
	}

}
