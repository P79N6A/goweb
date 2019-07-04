package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)
import "github.com/Shopify/sarama"

func main() {
	client()
	//proxy()
	/*reader := bufio.NewReader(os.Stdin)
	for {
		str, _, _ := reader.ReadLine()
		fmt.Println(string(str))
	}*/

}
func proxy() {
	// urli := url.URL{}
	// urlproxy, _ := urli.Parse("https://127.0.0.1:9743")
	//os.Setenv("HTTP_PROXY", "http://127.0.0.1:12639")
	//os.Setenv("HTTPS_PROXY", "https://127.0.0.1:12639")
	c := http.Client{
		// Transport: &http.Transport{
		//  Proxy: http.ProxyURL(urlproxy),
		// },
	}
	if resp, err := c.Get("http://118.89.137.42:9092"); err != nil {
		log.Fatalln(err)
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s\n", body)
	}
}

// 客户端可以用来获取消费者和生产者,还可以获取kafka的broker信息和topic信息,以及每个topic中的offset等
func client() {
	os.Setenv("http.proxyHost", "127.0.0.1")
	os.Setenv("http.proxyPort", "12639")
	//os.Setenv("http.proxy","127.0.0.1:12639")
	//os.Setenv("HTTP_PROXY", "http://127.0.0.1:12639")
	//os.Setenv("HTTPS_PROXY", "https://127.0.0.1:12639")
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_1
	client, err := sarama.NewClient([]string{
		"127.0.0.1:9092",
		//"47.103.2.208:9092",
		//"118.89.137.42:9092",
	}, config)
	if err != nil {
		panic("client create error")
	}
	defer client.Close()

	//获取主题的名称集合
	topics, _ := client.Topics()
	for _, e := range topics {
		fmt.Println("topic: ", e)
	}

	//获取broker集合
	brokers := client.Brokers()
	for _, broker := range brokers { //输出每个机器的地址
		fmt.Println("broker: ", broker.ID(), broker.Addr())
	}

	partitions, _ := client.Partitions("java_topic")
	for _, v := range partitions {
		fmt.Println("partition: ", v)
	}
	partitions, _ = client.Partitions("kafka_go_test")
	for _, v := range partitions {
		fmt.Println("partition: ", v)
	}
}
