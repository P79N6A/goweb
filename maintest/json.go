package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		Username string `json:"name"`
		Password string `json:"password"`
	}
	jsonstr := `[
  {"name":"李一", "password":"123456"},
  {"name":"李二", "password":"8484949"},
  {"name":"李三", "password":"fq885485"},
  {"name":"李四", "password":"fdi659"},
  {"name":"李五", "password":"18fedws"},
  {"name":"李六", "password":"127875few3456"},
  {"name":"李七", "password":"12KOjoi56"},
  {"name":"李八", "password":"12feef56"},
  {"name":"李九", "password":"ew88ss"},
  {"name":"李十", "password":"148ew62"},
  {"name":"李十一", "password":"853ewf4"}
]`

	users := new([]User)
	_ = json.Unmarshal([]byte(jsonstr), &users)
	fmt.Println(users)

	fmt.Println(fmt.Sprintf("新增用户id为%d", 4))
}
