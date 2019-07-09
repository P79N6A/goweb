package main

import (
	"goweb/apis"
	"goweb/consumer"
	//这里讲db作为go/databases的一个别名，表示数据库连接池
	"goweb/databases"
	"goweb/router"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goweb/docs"
)

// @Title Gin API
// @Version 1.0
// @description API接口文档.
func main() {
	//当整个程序完成之后关闭数据库连接
	defer databases.SqlDB.Close()
	go consumer.Consumer(apis.Address, apis.Topic)
	r := router.InitRouter()
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = r.Run(":8080")
}
