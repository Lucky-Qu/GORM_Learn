package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "student:luckyqu717@tcp(127.0.0.1:3306)/GORM_Learn?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "CRUD_",
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}
