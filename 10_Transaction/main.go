package main

import "gorm.io/gorm"

const (
	DSN = "student:luckyqu717@tcp(127.0.0.1:3306)/GORM_Learn?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	db := ConnectDB(DSN)
	if !db.Migrator().HasTable(&User{}) {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
	//事务方法
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		tx.Create(&User{
			Name: "LuckyQu",
			Age:  20,
		})
		return err
	})
	if err != nil {
		panic(err)
	}

	tx := db.Begin()
	tx.Create(&User{
		Name: "Lucky",
		Age:  22,
	})
	tx.Commit()
	tx = db.Begin()
	tx.Create(&User{
		Name: "LuckyKai",
		Age:  23,
	})
	tx.Rollback()

	tx.Commit()
	tx = db.Begin()
	tx.Create(&User{
		Name: "Name",
		Age:  18,
	})
	tx.SavePoint("SavePoint1")
	tx.Create(&User{
		Name: "Here",
		Age:  28,
	})
	tx.RollbackTo("SavePoint1")
	tx.Commit()
}
