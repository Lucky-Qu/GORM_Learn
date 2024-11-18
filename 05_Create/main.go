package main

import "fmt"

func main() {
	db := ConnectDB()
	if !db.Migrator().HasTable(&User{}) {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
	msg := db.Create(&User{
		Name: "LuckyQu",
		Age:  20,
	})
	if msg.Error != nil {
		panic(msg.Error)
	} else {
		fmt.Println("添加数据成功！")
	}
}
