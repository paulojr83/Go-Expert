package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico1/configs"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico1/internal/service"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
)

type ViaCepResponse struct {
	City   string  `json:"city"`
	Temp_C float64 ` json:"temp_C,omitempty"`
	Temp_F float64 `json:"temp_F,omitempty"`
	Temp_K float64 `json:"temp_K,omitempty"`
}

func FindCepMux(w http.ResponseWriter, r *http.Request) {
	var cep service.CepData
	if err := json.NewDecoder(r.Body).Decode(&cep); err != nil {
		if cep.Cep == "" || len(cep.Cep) != 8 {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		}
		return
	}
	configData := &configs.Conf{
		Service2: viper.GetString("SERVICE2_API"),
	}
	ctx := r.Context()
	result, err := service.FetchViaCep(ctx, cep, configData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.Erro {
		http.Error(w, "Can not find zipcode", http.StatusNotFound)
		return
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ViaCepResponse{
		City:   result.City,
		Temp_C: result.Temp_C,
		Temp_F: result.Temp_F,
		Temp_K: result.Temp_K,
	})
}

func FindCep(c *gin.Context) {
	var cep service.CepData
	if err := c.ShouldBindJSON(&cep); err != nil {
		if cep.Cep == "" || len(cep.Cep) != 8 {
			c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		}
		return
	}
	configData := &configs.Conf{
		Service2: viper.GetString("SERVICE2_API"),
	}
	trace := otel.Tracer("FindCep-trace")
	ctx, span := trace.Start(c.Request.Context(), c.Request.URL.Path)
	defer span.End()

	result, err := service.FetchViaCep(ctx, cep, configData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if result.Erro {
		c.JSON(http.StatusNotFound, "Can not find zipcode")
		return
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Request.Header))
	c.JSON(http.StatusOK, &ViaCepResponse{
		City:   result.City,
		Temp_C: result.Temp_C,
		Temp_F: result.Temp_F,
		Temp_K: result.Temp_K,
	})
}
