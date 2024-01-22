package main

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "quickstart-events"
	for i := 0; i < 5; i++ {
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(strconv.Itoa(i)),
			Key:   []byte(fmt.Sprintf("Hello %d", i)),
		}

		err = producer.Produce(msg, nil)
		if err != nil {
			panic(err)
		}
	}
	producer.Flush(5 * 1000)
}
