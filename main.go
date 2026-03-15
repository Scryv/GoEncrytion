package main

import (
	"crypto/rand"
	_"crypto/sha512"
	_"encoding/hex"
	"fmt"
)


const saltLength = 14 //length salt Const cause needs to be a fixed length

func genRandoSalt(saltLength int) []byte {
	var salt = make([]byte, saltLength)
	rand.Read(salt)

	return salt
}


func main(){
salt := genRandoSalt(saltLength)
fmt.Println(salt)
}
