package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	/*
		m := math.NewMath(10, 10)
		fmt.Println(m.Add())
	*/
	uuid := uuid.New()
	fmt.Println(uuid.String())
}
