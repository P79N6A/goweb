package apis

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SetSession(c *gin.Context) {
	key, value := c.Param("key"), c.Param("value")
	session := sessions.Default(c)
	session.Set(key, value)
	session.Save()
	c.String(http.StatusOK, "set session: %s : %s", key, value)
}

func GetSession(c *gin.Context) {
	key := c.Param("key")
	session := sessions.Default(c)
	value := session.Get(key)
	c.String(http.StatusOK, "get session : %s : %s", key, value)
}

func GetCookie(c *gin.Context) { //读取cookie
	cookie, err := c.Cookie("mysession")
	if err != nil {
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		log.Fatalln(err)
	}
	log.Printf("Cookie value: %s \n", cookie)
}
func SetCookie(c *gin.Context) {
	c.SetCookie("gin_cookie", "myValue", 3600, "/", "localhost", false, true)
	c.String(http.StatusOK, "set cookies")
}

func LoginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "login")
}
