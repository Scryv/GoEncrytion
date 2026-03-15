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

func hashPasswd(passwd string, salt []byte) string{
	var passwdBytes = []byte(passwd) //creates byte slice of the passwd str
	passwdBytes = append(passwdBytes, salt...) //appends and the ... is for since salt is a slice
	
}

func main(){
salt := genRandoSalt(saltLength)
fmt.Println(salt)
}
