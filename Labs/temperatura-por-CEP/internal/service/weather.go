package service

import (
	"encoding/json"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/temperatura-por-CEP/configs"
	"github.com/paulojr83/Go-Expert/Labs/temperatura-por-CEP/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

type CurrentWeather struct {
	Temp_c float64 ` json:"temp_C,omitempty"`
	Temp_f float64 `json:"temp_F,omitempty"`
}

type Weather struct {
	Current CurrentWeather `json:"current"`
}

type WeatherResult struct {
	Temp_C float64 ` json:"temp_C,omitempty"`
	Temp_F float64 `json:"temp_F,omitempty"`
	Temp_K float64 `json:"temp_K,omitempty"`
}

func FetchWeather(city string, configData *configs.Conf) (*WeatherResult, error) {

	city = utils.ClearSpecialCharacter(city)
	// Substituir espaços em branco por underscores
	city = strings.ReplaceAll(city, " ", "_")

	url := fmt.Sprintf(configData.WeatherApi, city, configData.WeatherApiKey)
	// Faz a requisição para a API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para a Api Via cep: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler resposta da Api Via cep: %v\n", err)
		return nil, err
	}

	// Converte o JSON em uma estrutura de endereço
	weather := &Weather{}
	err = json.Unmarshal(body, weather)

	if err != nil {
		fmt.Printf("Erro ao decodificar resposta da Api Via cep: %v\n", err)
		return nil, err
	}

	kelvin := utils.CelsiusToKelvin(weather.Current.Temp_c)
	result := WeatherResult{
		Temp_C: weather.Current.Temp_c,
		Temp_F: weather.Current.Temp_f,
		Temp_K: kelvin,
	}
	return &result, nil
}
