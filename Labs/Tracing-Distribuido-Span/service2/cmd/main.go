package main

import (
	"context"
	"fmt"
	web "github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico2/internal"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico2/internal/api/wep/controller"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrade "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	viper.AutomaticEnv()
	/*viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}*/
}

func main() {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initProvider(viper.GetString("OTEL_SERVICE_NAME"), viper.GetString("OTEL_EXPORTER_OTLP_ENPOINT"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatalf("Failed to shutdown TraceProvider %w", err)
		}
	}()

	trace := otel.Tracer("microservice-trace")
	templateData := &web.TemplateData{
		OTELTracer: trace,
	}
	server := web.NewServer(templateData)
	router := server.CreateServer()

	router.GET("/cep/:cep", controller.FindCep)

	go func() {
		log.Println("Starting server on port: ", viper.GetString("HTTP_PORT"))
		if err := http.ListenAndServe(viper.GetString("HTTP_PORT"), router); err != nil {
			log.Fatal(err)
		}
	}()

	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due to other reason...")
	}

	// Create a timeout context for the graceful shutdown
	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
}

func initProvider(serviceName, collectorURL string) (func(ctx context.Context) error, error) {
	ctx := context.Background()

	// Configuração do exportador Zipkin
	exporter, err := zipkin.New(
		viper.GetString("ZIPKIN_ENDPOINT"),
	)

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to create resouce: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, collectorURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to create gRPC connection to collector: %w", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("Failed to create trace exporter: %w", err)
	}

	bsp := sdktrade.NewBatchSpanProcessor(traceExporter)
	traceProvider := sdktrade.NewTracerProvider(
		sdktrade.WithBatcher(exporter),
		sdktrade.WithSampler(sdktrade.AlwaysSample()),
		sdktrade.WithResource(res),
		sdktrade.WithSpanProcessor(bsp),
	)

	otel.SetTracerProvider(traceProvider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return traceProvider.Shutdown, nil
}
