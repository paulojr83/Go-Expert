package main

import (
	"fmt"
	"net/http"
	"time"
)

var number1 uint64 = 0

// go run -race face-main.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number1++
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essas página %d %s", number1, "vezes")))
		time.Sleep(300 * time.Millisecond)
	})
	http.ListenAndServe(":3000", nil)
}
