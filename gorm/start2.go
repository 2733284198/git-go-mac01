package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	Code  string
	Price uint
}

/*type User struct {
	id int `gorm:"primary_key"`
	name  string
	age 	uint
}*/

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	//db.AutoMigrate(&User{})

	// 创建
	db.Create(&User{name: "name1", age: 20})
	db.Create(&User{name: "name2", age: 20})

	// 读取
	//var User user
	//db.First(&user, 1)                   // 查询id为1的product

	//db.First(&user, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	//db.Delete(&product)
}
