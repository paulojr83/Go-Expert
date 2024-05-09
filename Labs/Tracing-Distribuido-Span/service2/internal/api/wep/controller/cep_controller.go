package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico2/configs"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico2/internal/service"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"net/http"
)

func FindCep(c *gin.Context) {
	cep := c.Param("cep")

	if cep == "" || len(cep) != 8 {
		c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	configData := &configs.Conf{
		ViaCepApi:     viper.GetString("VIA_CEP_API"),
		WeatherApiKey: viper.GetString("WEATHER_API_KEY"),
		WeatherApi:    viper.GetString("WEATHER_API"),
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
	weather, err := service.FetchWeather(ctx, result.Localidade, configData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, weather)
}
