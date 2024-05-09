package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/Tracing-Distribuido-Span/servico2/configs"
	"go.opentelemetry.io/otel"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro"`
}

func FetchViaCep(ctx context.Context, cep string, configData *configs.Conf) (*ViaCep, error) {
	trace := otel.Tracer("FetchViaCep-trace")
	ctx, span := trace.Start(ctx, fmt.Sprintf("Calling:%s", "FetchViaCep"))
	defer span.End()

	// Faz a requisição para a API
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(configData.ViaCepApi, cep), nil)
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
	viaCep := &ViaCep{}
	err = json.Unmarshal(body, viaCep)

	if err != nil {
		fmt.Printf("Erro ao decodificar resposta da Api Via cep: %v\n", err)
		return nil, err
	}
	// Envie o resultado para o canal
	return viaCep, nil
}
