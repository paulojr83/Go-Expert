package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/paulojr83/Go-Expert/gRPC/internal/database"
	"github.com/paulojr83/Go-Expert/gRPC/internal/pb"
	"github.com/paulojr83/Go-Expert/gRPC/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("failed to open database: %v", err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcService := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcService, categoryService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcService.Serve(lis); err != nil {
		panic(err)
	}
}
