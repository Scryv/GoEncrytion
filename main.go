package main

import (
	"crypto/rand"
	_"crypto/sha512"
	_"encoding/hex"
	"fmt"
)


const saltLength = 14 //length salt Const cause needs to be a fixed length

func genRandoSalt(saltLength int) []byte {  //func for creating random salt
	var salt = make([]byte, saltLength) // makes a byte slice variable called salt
	rand.Read(salt) //reads the slice and fully changes it and ads its own rando value

	return salt //returns salts
}


func main(){
salt := genRandoSalt(saltLength)
fmt.Println(salt)
}
