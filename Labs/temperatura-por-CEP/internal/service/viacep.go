package service

import (
	"encoding/json"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/temperatura-por-CEP/configs"
	"io/ioutil"
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
}

func FetchViaCep(cep string, configData *configs.Conf) (*ViaCep, error) {

	// Faz a requisição para a API
	resp, err := http.Get(fmt.Sprintf(configData.ViaCepApi, cep))
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
	viaCep := &ViaCep{}
	err = json.Unmarshal(body, viaCep)

	if err != nil {
		fmt.Printf("Erro ao decodificar resposta da Api Via cep: %v\n", err)
		return nil, err
	}
	// Envie o resultado para o canal
	return viaCep, nil
}
