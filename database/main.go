package main

import (
	"gorm.io/gorm" //lets me use go structs in plaats van sql
	"github.com/glebarez/sqlite"
	"fmt"
)

type Data struct {
	//gorm.Model auto adds fields like id created/updated/deleted 
	Username string //be sure to if scan dont forget id
	Password string
	Salt string `gorm:"uniqueIndex:idx_salt"` //slug needs to be unique or wont work
}

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) //opens db and connects grom

func (d Data) String() string { //prints post in 1 line not sm rando bs
	return fmt.Sprintf("Username: %s, Password: %s, Salt: %s", d.Username, d.Password, d.Salt)
}

func createPost(username string, passwd string, salt string) Data { //func for creating post and also returns it
	newPost := Data{Username: username, Password: passwd, Salt: salt} //new post with TitleandSlug your input 
	if res := db.Create(&newPost); res.Error != nil { //var of the create func res if res error
	panic(res.Error) //not nil or duplicate it wil give error
}
return newPost
}

func main(){
	db.AutoMigrate(&Data{}) //autocreates tables and updates schema
	freshPost := createPost("esrr", "5b0a277cc82b772e88eee648a8d9f69be687646426b7b9b1b5f8fbca4243025d8bea619b30a1457d5e26881dd9a70caeea7dd858a8690aaeadb9f257", "3c68fad118676accaddb3") //calls func and saves row in db
	fmt.Println(freshPost) //normally gives the struct u appended but cause of string Post func it returns clean
}
