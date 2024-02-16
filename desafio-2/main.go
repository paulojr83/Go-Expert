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
type API uint8

const (
	Brasilapi API = iota
	Viacep
)

func main() {
	// Canal para receber o resultado
	result := make(chan *Address, 2)

	// Faz as requisições concorrentemente
	go fetchAddress("https://brasilapi.com.br/api/cep/v1/01153000", result, Brasilapi)
	go fetchAddress("http://viacep.com.br/ws/01153000/json/", result, Viacep)

	// Seleciona o primeiro resultado disponível
	select {
	case address := <-result:
		fmt.Println("Endereço encontrado pela primeira API:")
		printAddress(address)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Nenhuma resposta recebida dentro do tempo limite.")
	}
}

func fetchAddress(url string, result chan<- *Address, api API) {
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

	address := &Address{}
	// Converte o JSON em uma estrutura de endereço
	if api == Viacep {
		viaoCep := &ViaoCep{}
		err = json.Unmarshal(body, viaoCep)
		address = &Address{
			Cep:        viaoCep.Cep,
			Logradouro: viaoCep.Logradouro,
			Bairro:     viaoCep.Bairro,
			Localidade: viaoCep.Localidade,
			Uf:         viaoCep.Uf,
			Api:        "Viacep",
		}
	}
	if api == Brasilapi {
		brasilApi := &BrasilApi{}
		err = json.Unmarshal(body, brasilApi)

		address = &Address{
			Cep:        brasilApi.Cep,
			Logradouro: brasilApi.Street,
			Bairro:     brasilApi.Neighborhood,
			Localidade: brasilApi.City,
			Uf:         brasilApi.State,
			Api:        "Brasilapi",
		}
	}

	if err != nil {
		fmt.Printf("Erro ao decodificar resposta da %s: %v\n", url, err)
		return
	}
	// Envie o resultado para o canal
	result <- address
}

func printAddress(address *Address) {
	fmt.Printf("CEP: %s\n", address.Cep)
	fmt.Printf("Logradouro: %s\n", address.Logradouro)
	fmt.Printf("Bairro: %s\n", address.Bairro)
	fmt.Printf("Localidade: %s\n", address.Localidade)
	fmt.Printf("UF: %s\n", address.Uf)
	fmt.Printf("Api: %s\n", address.Api)
}
