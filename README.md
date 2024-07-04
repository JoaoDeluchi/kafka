# Kafka 

## Topic 

A topic in Kafka is a communication channel used to organize messages. Each topic has a unique name across the entire Kafka cluster. Messages are sent to and read from specific topics.

It is possible to have multiple consumers for a single topic, each reading different messages.

## Partitions 

Partitions are immutable and ordered sequences of messages. Each message receives an ID, known as an offset.

Topics can be divided into different partitions to increase throughput.

Messages are consumed sequentially within a partition. Kafka ensures the order of messages within a single partition, so if message ordering is crucial for your application, it is recommended to use only one partition.

### Keys 

Keys are used to ensure that related messages are placed in the same partition when necessary (e.g., when consumption order is mandatory).

If message order does not matter, keys can be omitted.

Messages with the same key are always added to the same partition.

### Replication Factor 

Partitions can be replicated across different brokers to guarantee resilience.

Replication has a cost; use a minimum of 2 but avoid using more than needed.

## Producer 

### Message Delivery Guarantees 

- **Ack 0**: No guarantee. Use when losing a message will not affect your business.
- **Ack 1**: Less performance compared to Ack 0. The leader partition guarantees the message but not the follower partitions.
- **Ack -1 (All)**: Lower performance. The leader partition waits for responses from follower partitions, ensuring the message is saved by all. Use when message loss is unacceptable.

### Delivery Semantics 

- **At most once**: Better performance but some messages may be lost.
- **At least once**: Moderate performance but may result in duplicate messages.
- **Exactly once**: Worst performance but ensures no message is lost or duplicated.

### Idempotent Producer

Kafka discards duplicate messages and guarantees sequence within a partition, but this decreases performance.

## Consumer 

A topic can have a single consumer. However, even with multiple partitions, the processing time will not decrease.

### Consumer Group

Having more consumers can increase throughput.

If the number of consumers exceeds the number of partitions, some consumers will be idle. It is best to match the number of partitions with the number of consumers.
[not actively consuming any messages because there are no partitions assigned]


# Command Line 

Go to the kafka container using the command: 
[docker exec -it kafka-kafka-1 bash] 

## Create topic 

[kafka-topics --create --topic=topicname --bootstrap-server=localhost:9092 --partitions=3]

## List topics

[kafka-topics --list --bootstrap-server=localhost:9092]

## Get Details of the topic

[kafka-topics --bootstrap-server=localhost:9092 --topic=topicname --describe]


