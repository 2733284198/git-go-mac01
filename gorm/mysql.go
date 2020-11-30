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

	user := User{Name: "q1mi", Age: 18}
	db.Create(&user) // 创建user

	// 查询记录
	//res := db.First(&user)
	//fmt.Println(res.name)
	//
	db.Debug().First(&user) // SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println("根据主键查询第一条记录：", user)

	defer db.Close()
}
