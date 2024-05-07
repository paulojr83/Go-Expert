package main

import (
	"fmt"
	"github.com/valyala/fastjson"
)

func main() {

	var p fastjson.Parser
	jsonData := `{"foo":"bar", "num": 123, "bool": true, "arr":[11,22,33]}`

	v, err := p.Parse(jsonData)

	if err != nil {
		panic(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("foo"))
	fmt.Printf("Obj=%v\n", v)

	a := v.GetArray("arr")

	for i, value := range a {
		fmt.Printf("Index: %d Value: %s\n", i, value)
	}
}
