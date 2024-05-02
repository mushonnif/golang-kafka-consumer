package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	err = consumer.Subscribe("helloworld", nil)

	if err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(1 * time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		}
	}
}
