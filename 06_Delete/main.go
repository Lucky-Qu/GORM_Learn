package main

func main() {
	db := ConnectDB()
	if !db.Migrator().HasTable(&User{}) {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
	db.Unscoped().Where("name = ?", "LuckyZ").Delete(&User{})
}
