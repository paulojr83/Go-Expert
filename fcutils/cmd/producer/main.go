package main

import (
	"fmt"
	"github.com/paulojr83/Go-Expert/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	for i := 0; i < 10; i++ {
		rabbitmq.Publish(ch, fmt.Sprintf("{\"name\":\"test-%d\"}", i), "amq.direct")
	}

}
