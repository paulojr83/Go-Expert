package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product //has many
}
type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category     //Belongs to
	SerialNumber SerialNumber // Has one
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/courses?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	/*
				// create
				db.Create(&Product{
					Name:  "Notebook",
					Price: 3000.0,
				})

				// create batch
				products := []Product{
					{Name: "Product 1", Price: 1},
					{Name: "Product 2", Price: 2},
					{Name: "Product 3", Price: 3},
					{Name: "Product 4", Price: 4},
					{Name: "Product 5", Price: 5},
				}
				db.Create(&products)

				var product Product
				db.First(&product, 1)
				db.First(&product, "name = ? ", "Product 1")
				fmt.Println(product)

				var products []Product
				db.Limit(2).Offset(2).Find(&products)
				for _, product := range products {
					fmt.Println(product)
				}


			//where
			var products []Product
			//db.Where("price > ?", 100).Find(&products)
			db.Where("name LIKE ?", "%book%").Find(&products)
			for _, product := range products {
				fmt.Println(product)
			}

		var p Product
		result := db.First(&p, 2)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// O registro não foi encontrado
				// Faça algo aqui para lidar com o caso em que o registro não existe
				// Por exemplo, exibir uma mensagem de erro ou retornar um erro personalizado
				panic("Registro não encontrado")
			} else {
				// Outro erro ocorreu durante a consulta
				panic("Erro ao consultar o banco de dados: " + result.Error.Error())
			}
		}

		if p.ID != 0 {
			fmt.Printf(p.Name)
			p.Name = "New Mouse"
			db.Save(&p)

			fmt.Printf(p.Name)
			db.Delete(&p)
		}

	*/

	/*
					category := Category{Name: "Eletronicos"}
					db.Create(&category)

					var category Category
					db.First(&category, 1)
					db.Create(&Product{
						Name:       "Notebook",
						Price:      1000.00,
						CategoryID: category.ID,
					})

				var products []Product
				db.Preload("Category").Find(&products)
				for _, product := range products {
					fmt.Println(product.Name, product.Category.Name)
				}

			category := Category{Name: "Comidas"}
			db.Create(&category)

			p := Product{
				Name:       "Balas de goma",
				Price:      10.00,
				CategoryID: category.ID,
			}
			db.Create(&p)

			db.Create(&SerialNumber{
				Number:    "123456",
				ProductID: p.ID,
			})



		var products []Product
		db.Preload("Category").Preload("SerialNumber").Find(&products)
		for _, product := range products {
			fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
		}

	*/

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println("-", product.Name, " Serial number:", product.SerialNumber.Number)
		}
	}

}
