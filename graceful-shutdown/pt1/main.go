package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	server := &http.Server{
		Addr: ":3000",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(4 * time.Second)
		writer.Write([]byte("Hello World!"))
	})

	go func() {
		fmt.Printf("Server is running at %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not listem on %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully sShutting the server: %v\n", err)
	}
	fmt.Println("Server stopped.")
}
