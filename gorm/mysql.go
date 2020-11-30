package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64  `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Name string `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	db, _ := gorm.Open("mysql", "root:root@/go?charset=utf8&parseTime=True&loc=Local")

	/*	user := User{Name: "q1mi", Age: 18}
		db.Create(&user) // 创建user

		db.Debug().First(&user)
		fmt.Println("插入记录：", user)*/

	var user []User
	//var user User

	db.Where("id > ?", "3").Find(&user)
	// `deleted_at` IS NULL AND ((name = 'jinzhu' AND age >= '20'))
	fmt.Println("记录3：", user)

	defer db.Close()
}
