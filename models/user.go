package models

import (
	"goweb/databases"
	"log"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db = databases.GOrmDB

func init() {
	//var db = databases.GOrmDB
	if !db.HasTable(&User{}) {
		// 创建表时添加表后缀
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
		log.Println("创建表user")
	}
}

func (user *User) AddUser() {
	db.Save(user)
}

func (user *User) GetUser(id int) {
	db.Find(user, id)
}

func (user *User) GetUserByName(username string) {
	db.Where("username=?", username).Find(user)
}
