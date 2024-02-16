package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BrasilApi struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaoCep struct {
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

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
	Api        string `json:"api"`
}

func main() {
	cBrasilapi := make(chan *BrasilApi)
	cViacep := make(chan *ViaoCep)

	// Faz as requisições concorrentemente
	go fetchBrasilapi("https://brasilapi.com.br/api/cep/v1/01153000", cBrasilapi)
	go fetchViaoCep("http://viacep.com.br/ws/01153000/json/", cViacep)

	// Seleciona o primeiro resultado disponível
	select {
	case result := <-cBrasilapi:
		address := &Address{
			Cep:        result.Cep,
			Logradouro: result.Street,
			Bairro:     result.Neighborhood,
			Localidade: result.City,
			Uf:         result.State,
			Api:        "Brasilapi",
		}
		printAddress(address)
	case result := <-cViacep:
		address := &Address{
			Cep:        result.Cep,
			Logradouro: result.Logradouro,
			Bairro:     result.Bairro,
			Localidade: result.Localidade,
			Uf:         result.Uf,
			Api:        "Viacep",
		}
		printAddress(address)

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Nenhuma resposta recebida dentro do tempo limite.")
	}
}

func fetchBrasilapi(url string, result chan<- *BrasilApi) {
	// Faz a requisição para a API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler resposta da %s: %v\n", url, err)
		return
	}

	// Converte o JSON em uma estrutura de endereço
	brasilApi := &BrasilApi{}
	err = json.Unmarshal(body, brasilApi)

	if err != nil {
		fmt.Printf("Erro ao decodificar resposta da %s: %v\n", url, err)
		return
	}
	// Envie o resultado para o canal
	result <- brasilApi
}

func fetchViaoCep(url string, result chan<- *ViaoCep) {
	// Faz a requisição para a API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer requisição para %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler resposta da %s: %v\n", url, err)
		return
	}

	// Converte o JSON em uma estrutura de endereço
	viaoCep := &ViaoCep{}
	err = json.Unmarshal(body, viaoCep)

	if err != nil {
		fmt.Printf("Erro ao decodificar resposta da %s: %v\n", url, err)
		return
	}
	// Envie o resultado para o canal
	result <- viaoCep
}
func printAddress(address *Address) {
	fmt.Println("Endereço encontrado pela primeira API:")
	fmt.Printf("CEP: %s\n", address.Cep)
	fmt.Printf("Logradouro: %s\n", address.Logradouro)
	fmt.Printf("Bairro: %s\n", address.Bairro)
	fmt.Printf("Localidade: %s\n", address.Localidade)
	fmt.Printf("UF: %s\n", address.Uf)
	fmt.Printf("Api: %s\n", address.Api)
}
