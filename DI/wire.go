//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/paulojr83/Go-Expert/DI/product"
)

var serRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		serRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
