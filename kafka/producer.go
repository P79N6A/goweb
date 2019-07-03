package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"time"
)

func main() {
	producer()
}

// 生产者
func producer() {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//config.Version = sarama.V0_11_0_2
	config.Version = sarama.V2_0_0_0

	//使用配置,新建一个异步生产者
	producer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}
	defer producer.AsyncClose()

	// 发送的消息,主题,key
	msg := &sarama.ProducerMessage{
		Topic: "kafka_go_test",
		Key:   sarama.StringEncoder("go_test"),
	}

	//value := "this is message"
	reader := bufio.NewReader(os.Stdin)
	for {
		//设置发送的真正内容
		//fmt.Scanln(&value)
		strBytes, _, _ := reader.ReadLine()
		value := string(strBytes)
		//将字符串转化为字节数组
		msg.Value = sarama.ByteEncoder(value)
		msg.Timestamp = time.Now()
		fmt.Printf("input [%s]\n", value)

		// 使用通道发送
		producer.Input() <- msg

		// 循环判断哪个通道发送过来数据.
		select {
		case suc := <-producer.Successes():
			fmt.Printf("msg {  offset: %d,    partition: %d,     timestamp: %s,     value: %s  }\n", suc.Offset, suc.Partition, suc.Timestamp.String(), suc.Value)
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}
