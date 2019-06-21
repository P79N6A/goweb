package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {

	c := cron.New()
	c.AddFunc("0/1 * * * * ? ", func() { fmt.Println("Every second exec") })
	c.AddFunc("@every 1s ", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	time.Sleep(time.Second * 10)
}
func f1() {
	fmt.Println("Every hour on the half hour")
}
