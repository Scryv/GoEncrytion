package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
type Data struct {
	gorm.Model //auto adds fields like id created/updated/deleted at
	Username string
	Password string
	Salt string `gorm:"uniqueIndex:idx_salt"` //slug needs to be unique or wont work
}

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) //opens db and connects grom

func printUsers() {
	var users []Data  //creates empty list called users []Data cause whole table

	result := db.Find(&users) //loads all rows from db into userSlice
	if result.Error != nil { //result.Error and not err is gorm doesnt return err
		panic(result.Error) //panic cause if broken state
	}

	csvData := "" //creates variable called csvData

	for _, user := range users { //loops trough every row
		csvData += fmt.Sprintf( //appends all these looped to csvData
			"%s,%s,%s\n",  //Sprintf cause can be saved in var 
			user.Username,
			user.Password,
			user.Salt,
		)
	}

	fmt.Println(csvData) //prints csv data
}

func main(){
	db.AutoMigrate(&Data{})
	printUsers()
}
