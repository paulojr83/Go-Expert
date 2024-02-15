package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/paulojr83/Go-Expert/packaging/math"
)

func main() {
	m := math.NewMath(10, 10)
	fmt.Println(m.Add())

	uid := uuid.New()
	fmt.Println(uid)
}
