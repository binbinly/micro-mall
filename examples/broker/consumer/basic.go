package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
	"log"

	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
)

var (
	topic = "go.micro.topic.foo"
)

// Example of a shared subscription which receives a subset of messages
func sharedSub(b broker.Broker) {
	_, err := b.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("consumer"))
	if err != nil {
		fmt.Println(err)
	}
}

// Example of a subscription which receives all the messages
func sub(b broker.Broker) {
	_, err := b.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()

	b := rabbitmq.NewBroker(broker.Addrs("amqp://guest:guest@localhost:5672"))

	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	sharedSub(b)

	fmt.Println("start success")
	select {}
}
