package main

func main() {
	db := ConnectDB()
	if !db.Migrator().HasTable(&User{}) {
		err := db.Migrator().CreateTable(&User{})
		if err != nil {
			panic(err)
		}
	}
	db.Model(&User{}).Where("name = ?", "LuckyQu").Update("name", "LuckyZ")

}
