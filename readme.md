main.go 入口

router注册路由

apis编写Handler方法，调用models读写数据库

（databases注册数据库连接池）

c.Param("key") //获取路径参数

c.DefaultQuery("key","defaultValue") //获取GET url参数
c.Query("key")

c.PostForm("key") //获取POST参数
c.DefaultPostForm("key","defaultValue")

c.Bind(&p)//将请求主体绑定到结构体中,目前支持JSON、XML、YAML和标准表单值(foo=bar&boo=baz)的绑定

编译到linux下运行
set GOOS=linux
go build main.go

swagger文档
go get -u github.com/swaggo/swag/cmd/swag
http://127.0.0.1:8080/swagger/index.html
swag init 生成doc文件夹
go run main.go启动项目

go-elasticsearch: Elastic 官方的 Go 语言客户端
https://github.com/elastic/go-elasticsearch#go-elasticsearch
https://www.infoq.cn/article/HVzMNKuYYmCkRTK-oZdp   中文版
