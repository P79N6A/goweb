package main

import (
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"time"
)

type Producer struct {
	BrokerServers string               `json:"broker_servers"`
	Config        *sarama.Config       `json:"config"`
	SyncProducer  sarama.SyncProducer  `json:"sync_producer"`
	AsyncProducer sarama.AsyncProducer `json:"async_producer"`
}

type Consumer struct {
	BrokerServers string         `json:"broker_servers"`
	Config        *sarama.Config `json:"config"`
}

//获取生产者实例
func (p *Producer) GetSysProducer() error {
	var err error
	p.SyncProducer, err = sarama.NewSyncProducer(strings.Split(p.BrokerServers, ","), p.Config)
	if err != nil {
		log.Printf("Create producer failed %s \n", err.Error())
		return err
	}
	return nil
}

//批量发送数据到kafka
func (p *Producer) BathProduceMsg(topic, key string, messages []string) error {
	if p.SyncProducer == nil {
		if e := p.GetSysProducer(); e != nil {
			log.Fatalf("Create kafka producer failed: %s \n", e.Error())
		}
	}
	defer p.SyncProducer.Close()
	msgs := make([]*sarama.ProducerMessage, len(messages))
	for i, m := range messages {
		msg := sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Key = sarama.StringEncoder(key)
		msg.Value = sarama.StringEncoder(m)
		msg.Timestamp = time.Now()
		msgs[i] = &msg
	}
	return p.SyncProducer.SendMessages(msgs)
}

//单条发送数据到kafka
func (p *Producer) ProducerMsg(topic, key string, msg string) (partition int32, offset int64, err error) {
	if p.SyncProducer == nil {
		if e := p.GetSysProducer(); e != nil {
			log.Fatalf("Create kafka producer failed: %s \n", e.Error())
		}
	}
	defer p.SyncProducer.Close()
	m := &sarama.ProducerMessage{}
	m.Topic = topic
	m.Key = sarama.StringEncoder(key)
	m.Value = sarama.StringEncoder(msg)
	m.Timestamp = time.Now()
	return p.SyncProducer.SendMessage(m)
}

const topicSign = "kafka_go"

//获取符合条件的topic以及对应的分区数
func (c *Consumer) GetTopicsAndPartitions() (map[string]int, error) {
	consumer, e := sarama.NewConsumer(strings.Split(c.BrokerServers, ","), nil)
	defer consumer.Close()
	if e != nil {
		log.Printf("Create kafka consumer failed: %s \n", e.Error())
		return nil, e
	}
	topics, _ := consumer.Topics()
	whTopic := make(map[string]int)
	for _, t := range topics {
		if strings.HasPrefix(t, topicSign) {
			if p, e := consumer.Partitions(t); e == nil {
				whTopic[t] = len(p)
			}
		}
	}
	return whTopic, nil
}

//消费kafka集群下满足条件的所有topic的数据
func (c *Consumer) Consume(done chan bool, consumeMethod func(msg *sarama.ConsumerMessage)) {
	whTopic, e := c.GetTopicsAndPartitions()
	if e != nil {
		return
	}
	for t := range whTopic {
		consumer, e := sarama.NewConsumer(strings.Split(c.BrokerServers, ","), nil)
		if e != nil {
			log.Printf("Create consumer for %s topic failed: %s \n", t, e.Error())
		}
		partitions, e := consumer.Partitions(t)
		if e != nil {
			log.Printf("Get %s topic's paritions failed: %s \n", t, e.Error())
		}
		for _, p := range partitions {
			pc, e := consumer.ConsumePartition(t, p, sarama.OffsetOldest)
			if e != nil {
				log.Printf("Start consumer for partition failed %d: %s \n", p, t)
				return
			}
			go func(pc sarama.PartitionConsumer) {
				for msg := range pc.Messages() {
					consumeMethod(msg)
				}
				done <- true
			}(pc)
		}
	}
}
