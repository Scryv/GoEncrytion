package main

import (
	"encoding/csv"
	"os"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)
const saltLength = 14 //length salt Const cause needs to be a fixed length

func hashPasswd(passwd string, salt []byte) string{
	var passwdBytes = []byte(passwd) //creates byte slice of the passwd str
	passwdBytes = append(passwdBytes, salt...) //appends and the ... is for since salt is a slice
	hash := sha512.Sum512(passwdBytes) //hashes the slice using sha512
	return hex.EncodeToString(hash[:]) //encodes to readable and [:] to change [64]byte to []byte
}

func getUser(username string) (string, string, bool) {
	file, err := os.Open("users.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	dataSheet, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, info := range dataSheet {
		if info[0] == username {
			return info[1], info[2], true
		}
	}

	return "", "", false
}

func doPasswdMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = hashPasswd(currPassword, salt)

	return hashedPassword == currPasswordHash
}

func main() {
    var username string
    var passwd string

    fmt.Println("Login: ")
    fmt.Scanln(&username)
    fmt.Println("Password: ")
    fmt.Scanln(&passwd)
    
	storedHash, storedSalt, found := getUser(username)
	saltBytes, err := hex.DecodeString(storedSalt)
	if !found {
		fmt.Println("User wasnt found ")
		return
	}
	
    if err != nil {
       fmt.Println("Salt decode failed:", err)
       return
    }

    match := doPasswdMatch(storedHash, passwd, saltBytes)
    if match {
       fmt.Println("Logged in")
    } else {
       fmt.Println("Invalid passwd")
    }
    
}
