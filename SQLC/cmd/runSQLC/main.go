package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // _ black identify, ignore no use import
	"github.com/paulojr83/Go-Expert/SQLC/internal/db"
)

func main() {

	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	/*err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:   uuid.New().String(),
		Name: "Backend",
		Description: sql.NullString{
			String: "Teste",
			Valid:  true,
		},
	})

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}*/

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name: "Backend updated",
		Description: sql.NullString{
			String: "updated",
			Valid:  true,
		},
		ID: "9b007e9f-1463-45ef-aa25-9a196afe5507",
	})
	if err != nil {
		panic(err)
	}

	category, err := queries.GetCategory(ctx, "9b007e9f-1463-45ef-aa25-9a196afe5507")
	println(category.ID, category.Name, category.Description.String)

	err = queries.DeleteCategory(ctx, "9b007e9f-1463-45ef-aa25-9a196afe5507")
	if err != nil {
		panic(err)
	}
}
