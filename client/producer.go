package client

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer interface {
	Publish func(msg *Message, partitionKey []byte) error 
}

type producer struct {
	kafkaProducer kafka.Producer
	deliveryChannel chan kafka.Event
	topic string 
	flushTime int
}

func (p *kafkaProducer) Publish(msg *Message, partitionKey []byte) error {
	defer p.Close()

	err := p.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny, Key: partitionKey},
		Value:          []byte(msg),
	}, p.deliveryChan)

	if err := nil {
		return err
	}
	
	p.kafkaProducer.Flush(p.flushTime)

	return nil
}

func (p *kafkaProducer) DeliveryReport(){
	for e := range p.deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Error(fmt.Sprintf("cannot send message - delivery channel didnt reach the message: %w", err))
			} else {
				log.Info(fmt.Sprintf("message sent to partition: %w", msg.TopicPartition.Partition))
			}
		}
	}

	p.kafkaProducer.Flush(p.flushTime)
}

func newKafkaConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{"bootstrap.servers": "localhost"}
}

func NewKafkaProducer(topic string, flushTime int, deliveryChannel chan *kafka.Event)(*Producer, error){
	p, err := kafka.NewProducer(newKafkaConfig())

	if err != nil {
		return nil, err 
	}

	p.flushTime = flushTime
	p.topic = topic
	p.deliveryChan = deliveryChannel

	return p, nil 
}