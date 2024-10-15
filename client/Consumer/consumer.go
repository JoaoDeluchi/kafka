package client 

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer interface {
	consume func() string
}

type consumer struct {
	kafkaConsumer kafka.Consumer
}

func (c *consumer) subscribeTopcs(topics []string) error {
	err := c.kafkaConsumer.SubscribeTopics(topics, nil)

	return err 
}

func (c *consumer) Consume() (kafka.Message, error) {
		msg, err := c.kafkaConsumer.ReadMessage(-1)
		if err == nil {
			log.Info(msg, nil
		}
	}
}

func NewKafkaConsumer(bootstrapServers, groupID, clientID,  shouldStoreOffsets kafka.ConfigValue, topics []string) (*consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"broker.address.family": "v4",
		"client.id": clientID
		"group.id":              groupID,
		"session.timeout.ms":    6000,
		"auto.offset.reset": "earliest",
		"enable.auto.offset.store": shouldStoreOffsets,
	})

	if err != nil {
		return nil, err
	}

	err := c.subscribeTopcs(topics)

	if err != nil {
		return fmt.Errorf("error: cannot subscribe to topics - %w", err)
	}

	return consumer{
		kafkaConsumer: c,
	}, nil 
}