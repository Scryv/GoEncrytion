package main

import (
	"gorm.io/gorm" //lets me use go structs in plaats van sql
	"github.com/glebarez/sqlite"
	"fmt"
)

type Post struct {
	gorm.Model //auto adds fields like id created/updated/deleted at
	Title string
	Slug string `gorm:"uniqueIndex:idx_slug"` //slug needs to be unique or wont work
	Likes uint
}

var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) //opens db and connects grom

func (p Post) String() string { //prints post in 1 line not sm rando bs
	return fmt.Sprintf("Post Title: %s, Slug: %s", p.Title, p.Slug)
}

func createPost(title string, slug string) Post { //func for creating post and also returns it
	newPost := Post{Title: title, Slug: slug} //new post with TitleandSlug your input 
	if res := db.Create(&newPost); res.Error != nil { //var of the create func res if res error
	panic(res.Error) //not nil or duplicate it wil give error
}
return newPost
}

func main(){
	db.AutoMigrate(&Post{}) //autocreates tables and updates schema
	freshPost := createPost("New Post Title", "new-slug") //calls func and saves row in db
	fmt.Println(freshPost) //normally gives the struct u appended but cause of string Post func it returns clean
}
