package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	. "goweb/apis"
	"io"
	"os"
)

func InitRouter() *gin.Engine {
	//gin.DisableConsoleColor()//禁用控制台颜色
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //将日志同时写入文件和控制台

	// 默认启动方式，包含 Logger、Recovery 中间件
	router := gin.Default()
	//router := gin.New()//无中间件启动

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// redis存session
	//store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//router.Use(sessions.Sessions("session", store))

	router.GET("/", IndexApi) //IndexApi为一个Handler
	router.POST("/person", AddPersonApi)
	router.GET("/persons", GetPersonsApi)
	router.GET("/person/:id", GetPersonApi)
	router.PUT("/person/:id", ModPersonApi)
	router.DELETE("/person/:id", DelPersonApi)

	router.GET("/session/:key/:value", SetSession)
	router.GET("/getsession/:key", GetSession)

	router.GET("/cookie", GetCookie)
	router.GET("/setcookie", SetCookie)
	return router
}
