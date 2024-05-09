package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico1/configs"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico1/internal/internal_error"
	"go.opentelemetry.io/otel"
	"io"
	"net/http"
)

type CepData struct {
	Cep string `json:"cep"`
}

type ViaCep struct {
	City   string  `json:"city"`
	Temp_C float64 ` json:"temp_C,omitempty"`
	Temp_F float64 `json:"temp_F,omitempty"`
	Temp_K float64 `json:"temp_K,omitempty"`
	Erro   bool    `json:"erro"`
}

func FetchViaCep(ctx context.Context, cepData CepData, configData *configs.Conf) (*ViaCep, *internal_error.InternalError) {

	trace := otel.Tracer("FetchViaCep-trace")
	ctx, span := trace.Start(ctx, fmt.Sprintf("Calling:%s", "FetchViaCep"))
	defer span.End()
	// Faz a requisição para a API
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(configData.Service2, cepData.Cep), nil)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para a Api Via cep: %v\n", err)
		return nil, internal_error.NewInternalServerError("Erro ao fazer requisição para a Api Via ")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para a Api Via cep: %v\n", err)
		return nil, internal_error.NewInternalServerError(err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, internal_error.HttpError(resp.StatusCode)
	}
	defer resp.Body.Close()
	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para a Api Via cep: %v\n", err)
		return nil, internal_error.NewInternalServerError("Erro ao fazer requisição para a Api Via ")
	}
	// Converte o JSON em uma estrutura de endereço
	viaCep := &ViaCep{}
	err = json.Unmarshal(body, viaCep)

	if err != nil {
		fmt.Printf("Erro ao decodificar resposta da Api Via cep: %v\n", err)
		return nil, internal_error.HttpError(http.StatusInternalServerError)
	}
	// Envie o resultado para o canal
	return viaCep, nil
}
