package internal

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// LoadTester contém a configuração do teste de carga
type LoadTester struct {
	url         string
	requests    int
	concurrency int
	results     chan *http.Response
}

// NewLoadTester cria uma nova instância do LoadTester
func NewLoadTester(url string, requests, concurrency int) *LoadTester {
	return &LoadTester{
		url:         url,
		requests:    requests,
		concurrency: concurrency,
		results:     make(chan *http.Response, requests),
	}
}

// Run executa os testes de carga
func (lt *LoadTester) Run() {
	var wg sync.WaitGroup

	// Limitando a quantidade de goroutines
	semaphore := make(chan struct{}, lt.concurrency)

	for i := 0; i < lt.requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- struct{}{}
			lt.makeRequest()
			<-semaphore
		}()
	}

	wg.Wait()
	close(lt.results)
}

// makeRequest faz uma requisição HTTP e envia o resultado para o canal de resultados
func (lt *LoadTester) makeRequest() {
	resp, err := http.Get(lt.url)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	lt.results <- resp
}

// Report gera o relatório do teste de carga
func (lt *LoadTester) Report(duration time.Duration) {
	totalRequests := 0
	status200 := 0
	statusCodes := make(map[int]int)

	for resp := range lt.results {
		totalRequests++
		statusCodes[resp.StatusCode]++
		fmt.Sprintf("StatusCode: ", resp.StatusCode)
		if resp.StatusCode == http.StatusOK {
			status200++
		}
		resp.Body.Close()
	}

	fmt.Println("Relatório do Teste de Carga")
	fmt.Println("--------------------------")
	fmt.Printf("URL: %s\n", lt.url)
	fmt.Printf("Total de Requisições: %d\n", totalRequests)
	fmt.Printf("Tempo Total: %v\n", duration)
	fmt.Printf("Status 200: %d\n", status200)
	fmt.Println("Outros Status Codes:")
	for code, count := range statusCodes {
		if code != http.StatusOK {
			fmt.Printf("Status %d: %d\n", code, count)
		}
	}
}
