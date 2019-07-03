package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "goweb/utils"
)

//因为我们需要在其他地方使用SqlDB这个变量，所以需要大写代表public
var SqlDB *sql.DB

//初始化方法
func init() {
	var err error
	//SqlDB, err = sql.Open("mysql", "root:zxcvbnm123@tcp(gz-cdb-ngh86yed.sql.tencentcdb.com:61077)/test?parseTime=true")
	SqlDB, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/test?parseTime=true")
	CheckErr(err)
	//连接检测
	err = SqlDB.Ping()
	CheckErr(err)
	SqlDB.SetMaxIdleConns(5)  //最大空闲连接
	SqlDB.SetMaxOpenConns(10) //最大打开连接
}
