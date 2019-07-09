package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer()
}

func consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true //接收失败通知
	config.Version = sarama.V0_10_2_1

	// consumer
	addrs := []string{"localhost:9092"}
	consumer, err := sarama.NewConsumer(addrs, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer consumer.Close()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	//topic := "java_topic"
	//topic := "kafka_go_test"
	topic := "go_topic"
	//topic := "__consumer_offsets"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partitionConsumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()
	fmt.Printf("consume topic [%s]:\n", topic)
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg{  topic: %s,  offset: %d,  partition: %d,  timestamp: %s,  key: %s,  value: %s}\n",
				msg.Topic, msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Key), string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
