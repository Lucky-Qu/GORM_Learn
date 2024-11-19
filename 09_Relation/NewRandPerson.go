package main

import (
	"gorm.io/gorm"
	"math/rand"
)

func NewRandPerson(db *gorm.DB) {
	name := []string{"LuckyQu", "LuckyKai", "LuckyTia", "Lucky", "Name"}
	p := Person{
		Name: name[rand.Intn(5)],
		Age:  rand.Intn(100),
		Purse: Purse{
			Money: rand.Intn(100),
		},
	}
	db.Create(&p)
}
