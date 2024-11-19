package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB(DSN string) (DB *gorm.DB) {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   "Relation_",
		},
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic(err)
	}
	return DB
}
