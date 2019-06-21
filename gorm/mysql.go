package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goweb/utils"
)

//表名是结构体名称的复数形式
//列名是字段名的蛇形小写
// delete_at有索引
/*type User struct {
	ID uint             // 列名为 `id`
	Name string         // 列名为 `name`
	Birthday time.Time  // 列名为 `birthday`
	CreatedAt time.Time // 列名为 `created_at`
}
// 重设列名
type Animal struct {
    AnimalId    int64     `gorm:"column:beast_id"`         // 设置列名为`beast_id`
    Birthday    time.Time `gorm:"column:day_of_the_beast"` // 设置列名为`day_of_the_beast`
    Age         int64     `gorm:"column:age_of_the_beast"` // 设置列名为`age_of_the_beast`
}

字段ID为主键
type User struct {
  ID   uint  // 字段`ID`为默认主键
  Name string
}

// 使用tag`primary_key`用来设置主键
type Animal struct {
  AnimalId int64 `gorm:"primary_key"` // 设置AnimalId为主键
  Name     string
  Age      int64
}
*/
type User struct {
	// 默认表名是`users`
	gorm.Model
	Name string
}

var db, err = gorm.Open("mysql", "root:admin@/test?charset=utf8&parseTime=True&loc=Local")

func main() {

	utils.CheckErr(err)
	defer db.Close()
	//b := db.HasTable("user")
	//db.AutoMigrate(&User{})
	//db.CreateTable(&User{})
	//db.DropTable(&User{})

	//db.Create(&User{Name:"易爽"})
	db.LogMode(true)

	var users []User
	db.Find(&users) //获取所有记录
	fmt.Println("所有记录", users)

	user := User{}
	db.First(&user) //第一条，按主键排序

	fmt.Println("第一条", user)

	user = User{}
	db.Last(&user) //最后一条
	fmt.Println("最后一条", user)

	user = User{}
	db.First(&user, 4) //使用主键获取记录
	fmt.Println("主键查询", user)

	// where条件查询
	user = User{}
	db.Where("name=?", "杨显").First(&user)
	fmt.Println("where查询", user)

	//insert()
	user = User{}
	//如果模型有DeletedAt字段，它将自动获得软删除功能！
	//那么在调用Delete时不会从数据库中永久删除，而是只将字段DeletedAt的值设置为当前时间。
	//db.Delete(&user)
}

func insert() {
	user := User{Name: "杨显"}
	db.Save(&user)
	fmt.Println(user)
	user = User{Name: "周杨"}
	db.Create(&user)
	fmt.Println(user)
	fmt.Println("插入数据完成")
}
