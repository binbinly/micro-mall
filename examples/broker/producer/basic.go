package main

import (
	"fmt"
	"log"
	"time"

	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
	// To enable rabbitmq plugin uncomment
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
)

var (
	topic = "go.micro.topic.foo"
)

func pub(b broker.Broker) {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := b.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
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

	pub(b)
}
