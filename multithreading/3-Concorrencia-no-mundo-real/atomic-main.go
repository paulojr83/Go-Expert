package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number2 uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number2++
		atomic.AddUint64(&number2, 1)
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essas página %d %s", number2, "vezes")))
		time.Sleep(300 * time.Millisecond)
	})
	http.ListenAndServe(":3000", nil)
}
