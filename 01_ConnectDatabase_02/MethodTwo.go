package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "student:luckyqu717@tcp(127.0.0.1:3306)/GORM_Learn?charset=utf8mb4&parseTime=True&loc=local",
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}
