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

set GOOS=linux
go build main.go