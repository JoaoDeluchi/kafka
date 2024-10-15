package client 

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	kafkaTopicEnv = "EXAMPLE_KAFKA_TOPIC"
)

func main() {
	deliveryChan := make(chan kafka.Event)

	producer, err := NewKafkaProducer(os.GetEnv(kafkaTopicEnv), 3000, deliveryChan)

	if err != nil {
		panic(fmt.Sprintf("cannot build application - NewKafkaProducer Error: %w", err))
	}
	// providing a partition key, you ensure that all messages will be published in the same partition
	err = producer.Publish("Message", []byte("example-partition-key"))

	if err != nil {
		panic(fmt.Sprintf("cannot publish message - producer.Publish Error: %w", err))
	}

	go producer.DeliveryReport()
}