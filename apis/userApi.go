package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/models"
	"log"
	"net/http"
	"strconv"
)

func LoginApi(c *gin.Context) {
	user := new(models.User)
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	log.Println(user)
	findUser := models.User{}
	findUser.GetUserByName(user.Username)

	if findUser.Password == user.Password {
		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"msg":      "登录成功",
			"username": user.Username,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "failure",
			"msg":    "账号或者密码错误！",
		})
	}
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	user := &models.User{}
	user.GetUser(intId)
	c.JSON(http.StatusOK, user)
}

func AddUser(c *gin.Context) {
	user := &models.User{}
	if err := c.Bind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println("post参数：", user)
	//username不能重复
	existUser := models.User{}
	existUser.GetUserByName(user.Username)
	if existUser.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "error",
			"msg":    "用户名已存在！",
		})
		return
	}

	user.AddUser()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    fmt.Sprintf("注册用户成功，id为%d", user.Id),
	})
}

func IsExist(c *gin.Context) {
	username := c.Param("username")
	user := models.User{}
	user.GetUserByName(username)
	if user.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "账号已存在！",
			"status": "exist",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "账号未注册",
			"status": "not exist",
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
