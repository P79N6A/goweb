package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer()
}

func consumer() {
	fmt.Println("consumer_test")

	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_0_0_0

	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer consumer.Close()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition("kafka_go_test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partitionConsumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg{    topic: %s,    offset: %d,    partition: %d,    timestamp: %s,    key: %s,    value: %s}\n",
				msg.Topic, msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Key), string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
