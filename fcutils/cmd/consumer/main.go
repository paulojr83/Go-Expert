package main

import (
	"fmt"
	"github.com/paulojr83/Go-Expert/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	queue := "minhafila"
	defer ch.Close()
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, msgs, queue)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
