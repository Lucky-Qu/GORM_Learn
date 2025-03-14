package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "student:luckyqu717@tcp(127.0.0.1:3306)/GORM_Learn?charset=utf8mb4&parseTime=True&loc=local)",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "Learn_",
			SingularTable: false,
		},
	})
	if err != nil {
		panic(err)
	}
	M := db.Migrator()
	err = M.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	fmt.Println(M.HasTable(&User{}))
}
