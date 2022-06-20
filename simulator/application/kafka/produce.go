package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gentildpinto/codelivery/simulator/application/route"
	"github.com/gentildpinto/codelivery/simulator/infra/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewProducer()
	r := route.NewRoute()
	json.Unmarshal(msg.Value, &r)

	r.LoadPositions()
	positions, err := r.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
