package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func firstTest(db *gorm.DB, u *User) {
	result := db.Model(&User{}).First(&u, 2)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(u)
}
func takeTest(db *gorm.DB, u *User) {
	result := db.Take(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(u)
}
func lastTest(db *gorm.DB, u *User) {
	result := db.Last(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println(u)
}

func main() {
	//连接数据库
	db := ConnectDB()
	//检测表，如果没有就创建
	if !db.Migrator().HasTable(&User{}) {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
	//声明接收结构体u
	u := User{}
	firstTest(db, &u)
	//takeTest(db, &u)
	//lastTest(db, &u)
	gu := getUser{}
	db.Model(&User{}).Where("name = ?", "LuckyQu").Find(&gu)
	fmt.Println(gu)
}
