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
```` Bash
docker exec -it kafka-kafka-1 bash
```` 

## Create topic 

```` Bash
kafka-topics --create --topic=topicname --bootstrap-server=localhost:9092 --partitions=3
````

## List topics

```` Bash
kafka-topics --list --bootstrap-server=localhost:9092
````

## Get Details of the topic

```` Bash
kafka-topics --bootstrap-server=localhost:9092 --topic=topicname --describe
````

## Start consumer assigned to a topic 

```` Bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=topicname
````
can use parameters 
    --group=groupname
    --from-beginning

## Start producer 

```` Bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=topicname
````

## Get details from the consumer group 

```` Bash
kafka-consumer-groups --bootstrap-server=localhost:9092 --group=groupname --describe
````

# Confluent Control Center (interface web)

access localhost:9021 with the container running


## Global Configuration Properties 


Property                          | Importance | Description
----------------------------------|------------|--------------------------------------------------------------------------
bootstrap.servers                 | HIGH       | One or more broker host servers
client.id                         | low        | Identifier, type string 
delivery.timeout.ms               | HIGH       | Limits time that a produce message wait for successfull delivery, type int
acks or request.required.acks	  | HIGH       | 0, 1 or -1 description in the producer session / Message Delivery Guarantees.
enable.idempotence                | HIGH       | does the messages be produced exactly once and in the original produce order? 

### more information - https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md


# Kafka Connect 

## Connectors 

### Source 

Connectors that allow you to get data from databases or any data source. Ex: Oracle, GCP Pub/Sub, Amazon SQS and etc.

### Sink

Connectors that allow you to save data in databases, queues and any data source. Ex: Azure Data Lake Storage, GCP Functions, ServiceNow, Datadog 

### Transform

How th


#### Link to discover Connectors - https://www.confluent.io/hub/?utm_medium=sem&utm_source=google&utm_campaign=ch.sem_br.nonbrand_tp.prs_tgt.kafka-connectors_mt.xct_rgn.latam_sbrgn.brazil_lng.eng_dv.all_con.connectors_con.kafka_connect&utm_term=kafka%20connect%20connectors&creative=&device=c&placement=&gad_source=1&gclid=Cj0KCQjwyL24BhCtARIsALo0fSA5dmvUPQBOpfmmZOxrav95cD35ElvwmzS_JpHCjLhoHMjrpz5h0SYaAoicEALw_wcB



Can use the control center to help to work with Kafka Connect 

## Setup a connector 

can be created based on configs that confluent provides or created from a file .properties, examples: 
- /connectors/mysql.properties 
- /connector/mongodb.properties 

In the mongodb connector file, you can see a simple example with the transform connector

# Confluent Cloud 

## Can choose a cloud provider to create an cluster (AWS, AZURE, GCP)

After create a cluster confluent cloud will provides the CLI documentation and everything you need to work with kafka in the cloud provider chosen. 

Can be created a free cluster to study, but be carefull. The enterprise / professional signature is expansive. 

### AWS - Amazon MSK / Managed Streaming for Apache Kafka 

AWS provides an cheapier way to work with kafka streams in AWS. Its just for streams, and dont provide any help or support to others Confluent Services.
