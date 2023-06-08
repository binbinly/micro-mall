package main

import (
	"context"
	"github.com/go-micro/plugins/v4/broker/rabbitmq"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
	"log"
	"time"
)

type TestEvent struct {
	Name string    `json:"name"`
	Age  int       `json:"age"`
	Time time.Time `json:"time"`
}

type Example struct {
}

func (e *Example) Handler(ctx context.Context, r *TestEvent) error {
	log.Printf("handler data:%#v\n", r)
	return nil
}

func main() {
	b := rabbitmq.NewBroker(broker.Addrs("amqp://guest:guest@localhost:5672"),
		rabbitmq.ExchangeName("exchange.text"))
	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	s := server.NewServer(server.Broker(b))

	service := micro.NewService(
		micro.Server(s),
		micro.Broker(b),
	)
	brkrSub := broker.NewSubscribeOptions(
		broker.Queue("topic.exchange.queue"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)
	// Register a subscriber
	err := micro.RegisterSubscriber(
		"my-test-topic",
		service.Server(),
		&Example{},
		server.SubscriberContext(brkrSub.Context),
		server.SubscriberQueue("topic.exchange.queue"),
	)
	if err != nil {
		log.Fatal(err)
	}

	e := micro.NewEvent("my-test-topic", service.Client())
	go func() {
		time.Sleep(5 * time.Second)
		log.Print("pub event")
		//jsonData, _ := json.Marshal(&TestEvent{
		//	Name: "test",
		//	Age:  16,
		//})
		err := e.Publish(context.Background(), &TestEvent{
			Name: "test",
			Age:  16,
		}, client.WithExchange("exchange.text"))
		if err != nil {
			log.Fatal(err)
		}
	}()

	// service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func handler(ctx context.Context, evt *TestEvent) error {
	log.Printf("receive event: %+v", evt)
	return nil
}
