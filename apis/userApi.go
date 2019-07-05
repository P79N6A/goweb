package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string)

func init() {
	users["ys"] = "123"
}

func LoginApi(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//username := c.Request.FormValue("username") //url写first_name=a&last_name=b
	//password := c.Request.FormValue("password")
	//fmt.Println(username, password)

	user := new(User)
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(user)
	value, ok := users[user.Username]
	if ok && value == user.Password {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"msg":    "登录成功",
			"data": gin.H{
				"username": user.Username,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "failure",
			"msg":    "账号或者密码错误！",
		})
	}
}

func GetHot(c *gin.Context) {
	strs := []string{"企业微信", "办公网", "VPN", "邮箱", "wifi", "Outlook", "网络安全"}
	c.JSON(http.StatusOK, strs)
}
func GetKeyword(c *gin.Context) {
	strs := []string{"分机号码发送流程", "分开发送", "分屏功能", "分区", "分割线",
		"在外办公", "在MAC机上安装", "企业微信", "企业云盘",
		"邮箱服务器", "邮箱查看", "电话会议", "电话无显示"}
	c.JSON(http.StatusOK, strs)
}
