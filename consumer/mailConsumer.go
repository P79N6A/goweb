package consumer

import (
	"github.com/Shopify/sarama"
	"log"
)

func Consumer(address []string, topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true //接收失败通知
	config.Version = sarama.V0_10_2_1

	// consumer
	//addrs := []string{"localhost:9092"}
	consumer, err := sarama.NewConsumer(address, config)
	if err != nil {
		log.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer consumer.Close()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("try create partitionConsumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()
	log.Printf("consume topic [%s]:\n", topic)
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("消费msg成功{  topic: %s,  offset: %d,  partition: %d,  timestamp: %s,  key: %s,  value: %s}\n",
				msg.Topic, msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Key), string(msg.Value))
			log.Println("执行发送邮件......") //todo 发送邮件，kafka消费记录到es
		case err := <-partitionConsumer.Errors():
			log.Printf("err :%s\n", err.Error())
		}
	}
}
