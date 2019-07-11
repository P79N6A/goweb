package main

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	/*client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client.Info())
	fmt.Println(elasticsearch.Version)*/

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(es.Info())
}
