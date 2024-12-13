package main

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"math/rand"
)

const (
	KafkaServer = "localhost:9092"
	KafkaTopic  = "messages-v1-topic"
)

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func producer() (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
	})

	if err != nil {
		return nil, err
	}

	return p, nil
}

// ProduceN Generate and push N messages to kafka
func ProduceN(num int) {
	p, _ := producer()
	defer p.Close()

	for i := 0; i < num; i++ {
		produce(p)
	}
}

func produce(p *kafka.Producer) {
	topic := KafkaTopic
	message := Message{
		ID:      uuid.New().String(),
		Content: uuid.New().String(),
		UserID:  rand.Intn(1000) + 1,
	}

	value, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		panic(err)
	}
}
