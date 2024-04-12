package main

import (
	"database/sql"
	"fmt"
	"github.com/paulojr83/Go-Expert/CleanArch/configs"
	"github.com/paulojr83/Go-Expert/CleanArch/internal/event/handler"
	"github.com/paulojr83/Go-Expert/CleanArch/internal/infra/graph"
	"github.com/paulojr83/Go-Expert/CleanArch/internal/infra/grpc/pb"
	"github.com/paulojr83/Go-Expert/CleanArch/internal/infra/grpc/service"
	"github.com/paulojr83/Go-Expert/CleanArch/internal/infra/web/webserver"
	"github.com/paulojr83/Go-Expert/CleanArch/pkg/events"

	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// go run .\main.go .\wire_gen.go
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db)

	webserver := webserver.NewWebServer(config.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webserver.AddHandler("POST", "/order", webOrderHandler.Create)
	webserver.AddHandler("GET", "/order", webOrderHandler.List)

	fmt.Println("Starting web server on port", config.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", config.GraphQLServerPort)
	http.ListenAndServe(":"+config.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	amqpConfigDial := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.RabbitMQUser, config.RabbitMQPassword, config.RabbitMQHost, config.RabbitMQPort)
	conn, err := amqp.Dial(amqpConfigDial)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
