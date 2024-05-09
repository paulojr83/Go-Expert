package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico2/configs"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico2/utils"
	"go.opentelemetry.io/otel"
	"io"
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
	City   string  `json:"city"`
}

func FetchWeather(ctx context.Context, city string, configData *configs.Conf) (*WeatherResult, error) {
	trace := otel.Tracer("FetchWeather-trace")
	ctx, span := trace.Start(ctx, fmt.Sprintf("Calling:%s", "FetchWeather"))
	defer span.End()

	locate := utils.ClearSpecialCharacter(city)
	// Substituir espaços em branco por underscores
	locate = strings.ReplaceAll(locate, " ", "_")

	url := fmt.Sprintf(configData.WeatherApi, locate, configData.WeatherApiKey)
	// Faz a requisição para a API
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para a Api Via cep: %v\n", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
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
		City:   city,
	}
	return &result, nil
}
