package main

import (
	"flag"
	"fmt"
	"github.com/paulojr83/Go-Expert/Desafios-tcnicos/Sistema-Stress-test/internal"
	"os"
	"time"
)

func main() {

	// Definindo parâmetros de entrada
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requisições")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL do serviço é obrigatória.")
		flag.Usage()
		os.Exit(1)
	}

	// Inicializando o LoadTester
	loadTester := internal.NewLoadTester(*url, *requests, *concurrency)
	start := time.Now()

	// Executando o teste de carga
	loadTester.Run()

	// Calculando o tempo total gasto
	duration := time.Since(start)

	// Gerando o relatório
	loadTester.Report(duration)
}
