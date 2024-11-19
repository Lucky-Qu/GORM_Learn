package main

import "fmt"

const (
	DSN = "student:luckyqu717@tcp(127.0.0.1:3306)/GORM_Learn?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	//定义错误
	var err error
	//连接数据库
	db := ConnectDB(DSN)
	//若没表则建表
	if !db.Migrator().HasTable(&Person{}) {
		err = db.Migrator().CreateTable(&Person{})
		if err != nil {
			panic(err)
		}
		err = db.Migrator().CreateTable(&Purse{})
		if err != nil {
			panic(err)
		}
	}
	//新建用户
	for range 10 {
		NewRandPerson(db)
	}
	//查询用户
	result := Person{}
	db.Preload("Purse").Where("name = ?", "LuckyQu").Find(&result)
	fmt.Println(result)
}
