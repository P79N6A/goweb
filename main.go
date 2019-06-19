package main

import (
	//这里讲db作为go/databases的一个别名，表示数据库连接池
	"goweb/databases"
	"goweb/router"
)

func main() {
	//当整个程序完成之后关闭数据库连接
	defer databases.SqlDB.Close()
	initRouter := router.InitRouter()
	_ = initRouter.Run("127.0.0.1:8081")
}
