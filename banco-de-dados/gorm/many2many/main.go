package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories"`
}
type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	/*
		category := Category{Name: "Cozinha"}
		db.Create(&category)

		category2 := Category{Name: "Eletronicos"}
		db.Create(&category2)

		p := Product{
			Name:       "Panela Eletrica",
			Price:      356.0,
			Categories: []Category{category, category2},
		}

		db.Create(&p)
		fmt.Println(p)
	*/
	fmt.Println("/********************* Categories with products ************************/")
	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println("*", category.Name)
		for _, product := range category.Products {
			fmt.Println("-", product.Name)
		}
	}
	fmt.Println("/********************* Products with category ************************/")
	var products []Product
	err = db.Model(&Product{}).Preload("Categories").Find(&products).Error
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println("*", product.Name)
		for _, category := range product.Categories {
			fmt.Println("-", category.Name)
		}
	}
}
