package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gentildpinto/codelivery/simulator/application/kafka"
	kfkinfra "github.com/gentildpinto/codelivery/simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kfkinfra.NewConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		go kafka.Produce(msg)
		fmt.Println(string(msg.Value))
	}
}
