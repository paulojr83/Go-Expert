package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico1/configs"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico1/internal/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type WebServer struct {
	TemplateData *TemplateData
}

func NewServer(template *TemplateData) *WebServer {
	return &WebServer{TemplateData: template}
}

type WebServerInterface interface {
	FetchViaCep(cepData service.CepData, configData *configs.Conf) (*service.ViaCep, error)
}

func (we *WebServer) CreateServer() *gin.Engine {
	// Configuração do middleware de injeção de contexto OpenTelemetry
	injector := otel.GetTextMapPropagator()
	router := gin.New()
	router.Use(func(c *gin.Context) {
		ctx := c.Request.Context()
		injector.Inject(ctx, propagation.HeaderCarrier(c.Request.Header))
		c.Next()
	})

	router.Use(gin.Recovery()) // Recuperação de pânico
	router.Use(gin.Logger())
	//router.Use(we.HandleRequest)
	// promhttp
	http.Handle("/metrics", promhttp.Handler())
	return router
}

type TemplateData struct {
	OTELTracer      trace.Tracer
	Content         string
	RequestNameOTEL string
}

func (we *WebServer) HandleRequest(c *gin.Context) {
	ctx := context.Background()
	// Acesse o propagador de texto do OpenTelemetry
	propagator := otel.GetTextMapPropagator()
	// Injete o contexto OpenTelemetry nos cabeçalhos HTTP da requisição
	propagator.Inject(ctx, propagation.HeaderCarrier(c.Request.Header))
	// Continue com o próximo middleware ou manipulador
	c.Next()
}
