package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/temperatura-por-CEP/configs"
	"github.com/paulojr83/Go-Expert/Labs/temperatura-por-CEP/internal/service"
	"github.com/spf13/viper"
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

	result, err := service.FetchViaCep(cep, configData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	weather, err := service.FetchWeather(result.Localidade, configData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, weather)
}
