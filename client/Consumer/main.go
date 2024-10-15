package client 


import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)


func main() {

	topics := []string{ 
		"test"
	}
	c, err := NewKafkaConsumer(
		"kafka_example:8080",
		"app-example-group",
		"app-example-consumer-id",
		true,
		topics
	)

	if err != nil {
		panic(fmt.Sprintf("cannot create consumer - NewKafkaConsumer error: %w", err.Error()))
	}

	for {
		msg, err := c.Consume()
		if err != nil {
			log.Error("cannot consume msg")
		} else{
			log.Info(fmt.Sprintf("Consumed Message: %s /n topic partition: %s", msg.Value, msg.TopicPartition))
		}
	}

}
