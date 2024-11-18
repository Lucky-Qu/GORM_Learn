package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := ConnectDB()
	if !db.Migrator().HasTable(&User{}) {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
	var u User
	db.Model(&User{}).Find(&u)
	fmt.Println(u)
}
