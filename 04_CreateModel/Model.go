package main

import "gorm.io/gorm"

type User struct {
	gorm.Model `gorm:"comment:测试"`
	Name       string
	Age        int
	Email      string
	Sex        string
}
