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
		DSN: "student:luckyqu717@/GORM_Learn",
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "Learn_",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")
	M := db.Migrator()
	if M.HasTable(&User{}) {
		err = M.DropTable(&User{})
		if err != nil {
			panic(err)
		}
		fmt.Println("我成功删除了一个表！")
	} else {
		err = M.CreateTable(&User{})
		if err != nil {
			panic(err)
		}
		fmt.Println("我成功创建了一个表！")
	}
}
