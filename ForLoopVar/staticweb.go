package main

import (
	"fmt"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("D:\\workspace estudo\\Go-Expert\\ForLoopVar\\static"))

	// Definir o handler raiz para servir arquivos estáticos
	http.Handle("/", fileServer)

	fmt.Println("Server running...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
