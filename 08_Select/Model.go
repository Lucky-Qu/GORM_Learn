package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"comment:名字"`
	Age  int    `gorm:"comment:年龄"`
}

type getUser struct {
	Name string
	Age  int
}
