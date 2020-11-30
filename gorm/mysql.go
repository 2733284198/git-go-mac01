package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string
	Age  int64
}

func main() {
	db, _ := gorm.Open("mysql", "root:root@/go?charset=utf8&parseTime=True&loc=Local")

	user := User{Name: "q1mi", Age: 18}

	db.NewRecord(user) // 主键为空返回`true`
	db.Create(&user)   // 创建user
	db.NewRecord(user) // 创建`user`后返回`false`

	defer db.Close()
}
