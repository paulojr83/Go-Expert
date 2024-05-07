package web

import (
	"embed"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"html/template"
	"io"
	"net/http"
	"time"
)

//go:embed template/*
var templateContent embed.FS

type WebServer struct {
	TemplateData *TemplateData
}

func NewServer(template *TemplateData) *WebServer {
	return &WebServer{TemplateData: template}
}

func (we *WebServer) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))

	// promhttp
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/", we.HandleRequest)

	return router
}

type TemplateData struct {
	Title              string
	BackgroundColor    string
	ResponseTime       time.Duration
	ExternalCallUrl    string
	ExternalCallMethod string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
	Content            string
}

func (we *WebServer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()

	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInit := we.TemplateData.OTELTracer.Start(ctx, "SPAN_INITIAL: "+we.TemplateData.RequestNameOTEL)
	time.Sleep(time.Second)
	spanInit.End()

	ctx, span := we.TemplateData.OTELTracer.Start(ctx, "Chama externa: "+we.TemplateData.RequestNameOTEL)
	defer span.End()

	time.Sleep(time.Millisecond * we.TemplateData.ResponseTime)

	if we.TemplateData.ExternalCallUrl != "" {
		var req *http.Request
		var err error

		if we.TemplateData.ExternalCallMethod == "GET" {
			req, err = http.NewRequestWithContext(ctx, "GET", we.TemplateData.ExternalCallUrl, nil)
		} else if we.TemplateData.ExternalCallMethod == "POST" {
			req, err = http.NewRequestWithContext(ctx, "POST", we.TemplateData.ExternalCallUrl, nil)
		} else {
			http.Error(w, "Invalid External call method", http.StatusInternalServerError)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		we.TemplateData.Content = string(bodyBytes)
	}

	tpl := template.Must(template.New("index.html").ParseFS(templateContent, "template/index.html"))
	err := tpl.Execute(w, we.TemplateData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
