package main

import (
	"fmt"
	"github.com/paulojr83/Go-Expert/packaging/math"
)

func main() {
	m := math.NewMath(10, 10)
	fmt.Println(m.Add())
}
