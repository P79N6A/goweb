package main

import (
	"fmt"

	"github.com/Gre-Z/common/email"
)

func main() {
	username := "yishuang150@163.com"
	password := "y1sh5angys537519"
	mail := email.New163mail(username, password)
	rec := []string{"164497083@qq.com"}
	err := mail.Info("标题", "作者", rec).SendText("文本内容")
	fmt.Println(err)
}
