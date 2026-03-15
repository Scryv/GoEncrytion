package main

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
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
	hash := sha512.Sum512(passwdBytes) //hashes the slice using sha512
	return hex.EncodeToString(hash[:]) //encodes to readable and [:] to change [64]byte to []byte
}

func main(){
var passwd string //just var for passwd
fmt.Println("What password do you want to hash: ") //prompt
fmt.Scanln(&passwd) //scans answer does stop by space tho also & so it can overwrite var
salt := genRandoSalt(saltLength) //call and assign genSalt
hashedpasswd := hashPasswd(passwd, salt) //call and asign hashPasswd
fmt.Println("Salt used", hex.EncodeToString(salt)) //Prints salt used
fmt.Printf("Your hashed passwd: %v\n", hashedpasswd) //prints the hashed salt+passwd
}
