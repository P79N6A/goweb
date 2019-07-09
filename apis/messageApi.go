package apis

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type mailmsg struct {
	Msg       string    `json:"msg"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Timestamp time.Time `json:"timestamp"`
}

var producer sarama.AsyncProducer

const Topic = "mail_topic"
const Key = "mail_key"

var Address = []string{"127.0.0.1:9092"}

func init() {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_2_1
	var err error
	producer, err = sarama.NewAsyncProducer(Address, config)
	if err != nil {
		log.Fatalf("producer_test create producer error :%s\n", err.Error())
		return
	}
}

func MailPush(c *gin.Context) {
	var mail mailmsg
	if err := c.BindJSON(&mail); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	mail.Timestamp = time.Now()
	log.Println("解析post参数：", mail)
	bytes, _ := json.Marshal(mail)
	msg := &sarama.ProducerMessage{
		Topic: Topic,
		Key:   sarama.StringEncoder(Key),
		Value: sarama.ByteEncoder(bytes),
	}
	producer.Input() <- msg
	// 循环判断哪个通道发送过来数据.
	select {
	case suc := <-producer.Successes():
		log.Printf("发送msg成功 {  offset: %d,    partition: %d,     timestamp: %s,     value: %s  }\n", suc.Offset, suc.Partition, suc.Timestamp.String(), suc.Value)
		//todo kafka生产记录到es
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	case fail := <-producer.Errors():
		log.Printf("err: %s\n", fail.Err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
		})
	}

}
