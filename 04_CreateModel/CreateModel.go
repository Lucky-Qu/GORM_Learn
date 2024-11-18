package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "student:luckyqu717@/GORM_Learn",
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "CM_",
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

}
