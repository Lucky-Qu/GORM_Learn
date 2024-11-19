package main

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name  string
	Age   int
	Purse Purse
}
type Purse struct {
	gorm.Model
	Money    int
	PersonID int
}
