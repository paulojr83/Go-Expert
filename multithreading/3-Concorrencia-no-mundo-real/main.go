package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você teve acesso a essas página %d %s", number, "vezes")))
		time.Sleep(300 * time.Millisecond)
	})
	http.ListenAndServe(":3000", nil)
}
