package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fastjson"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var p fastjson.Parser
	jsonData := `{"user": {"name": "Json Doe", "age": 30}}`

	v, err := p.Parse(jsonData)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Obj=%v\n", v)

	userObj := v.GetObject("user")

	fmt.Printf("user=%v\n", userObj)
	fmt.Printf("User name: %s\n", userObj.Get("name"))
	fmt.Printf("User age: %s\n", userObj.Get("age"))

	userJSON := v.Get("user").String()

	var user User

	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}
	fmt.Printf("User: %v", user)

	var sc fastjson.Scanner

	sc.Init(`   {"foo":  "bar"  }[  ]
		12345"xyz" true false null    `)

	for sc.Next() {
		fmt.Printf("%s\n", sc.Value())
	}
	if err := sc.Error(); err != nil {
		log.Fatalf("unexpected error: %s", err)
	}

}
