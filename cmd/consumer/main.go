package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	KafkaServer  = "localhost:9092"
	KafkaTopic   = "messages-v1-topic"
	KafkaGroupId = "support-service"
)

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
		"group.id":          KafkaGroupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	topic := KafkaTopic
	c.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var message Message
			err := json.Unmarshal(msg.Value, &message)
			if err != nil {
				fmt.Printf("Error decoding message: %v\n", err)
				continue
			}

			fmt.Printf("Received message: %+v\n", message)
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
