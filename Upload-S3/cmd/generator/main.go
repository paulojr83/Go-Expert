package main

import (
	"fmt"
	"os"
)

func main() {
	qtd := 1000
	for i := 0; i <= qtd; i++ {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("Hello, World!")
	}
}
