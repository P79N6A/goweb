package databases

import (
	"github.com/jinzhu/gorm"
	"goweb/utils"
)

var GOrmDB *gorm.DB

func init() {
	var err error
	GOrmDB, err = gorm.Open("mysql", "root:admin@/test?charset=utf8&parseTime=True&loc=Local")
	utils.CheckErr(err)
	GOrmDB.SingularTable(true)
}
