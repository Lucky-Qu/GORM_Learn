package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"comment:名字"`
	Age  int    `gorm:"comment:年龄"`
}
