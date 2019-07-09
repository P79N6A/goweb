package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	. "goweb/apis"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func InitRouter() *gin.Engine {
	//gin.DisableConsoleColor()//禁用控制台颜色
	f, _ := os.Create("gin.log")
	//f, _ := os.Open("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) //将日志同时写入文件和控制台
	//gin.SetMode(gin.ReleaseMode)// 正式版发布

	router := gin.Default() // 默认启动方式，包含 Logger、Recovery 中间件
	//router := gin.New()//无中间件启动

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	//redis存session
	//store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//router.Use(sessions.Sessions("session", store))

	// HTML渲染
	router.LoadHTMLGlob("templates/*")

	// 使用中间件
	//router.Use(Logger())

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

	router.GET("/temp", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
			"name":  "易爽",
		})
	})

	// http重定向
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://google.com")
	})
	// 路由重定向，浏览器地址不变化
	router.GET("/redirect1", func(c *gin.Context) {
		c.Request.URL.Path = "/persons"
		router.HandleContext(c)
	})

	//分组路由
	v2 := router.Group("v2")
	{
		v2.GET("login", LoginEndpoint)
		v21 := v2.Group("v21") //内嵌分组路由
		{
			v21.GET("login", LoginEndpoint)
		}
	}
	v3 := router.Group("user")
	{
		v3.POST("/login", LoginApi)
		v3.GET("/hot", GetHot)
		v3.GET("/keyword", GetKeyword)
		v3.GET("/get/:id", GetUser)
		v3.POST("/register", AddUser)
		v3.GET("/exist/:username", IsExist)
	}
	msgGroup := router.Group("/message")
	{
		msgGroup.POST("mail", MailPush)
	}

	// simulate some private data
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}
	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) { // 使用BasicAuth()（验证）中间件
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 在中间件或处理程序中启动新的Goroutines时，
	// 你不应该使用其中的原始上下文，你必须使用只读副本（c.Copy()）
	router.GET("/long_async", func(c *gin.Context) {
		// 创建要在goroutine中使用的副本
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)
			// 这里使用你创建的副本
			log.Println("Done! in path " + cCp.Request.URL.Path)
			log.Println("Done! in path " + c.Request.URL.Path) // 可以运行，但不应该
		}()
	})

	router.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)
		// 这里没有使用goroutine，所以不用使用副本
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	return router
}

// 自定义中间件
func Logger() gin.HandlerFunc { //打印请求耗时
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		//log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println("耗时：", latency, "状态码：", status)
	}
}