package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	Users []User `json:"users"` //Remember NOSPACES
}

type User struct {
	Name  string `json:"name"`
	Phone string `json:"PHONE"` //Case insensitive matches json key
	Email string `json:"email"`
}

func main() {
	fmt.Println("Accessing JSON data from GO")

	// Opening the json file
	jsonFile, err := os.Open("users.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfuly Opened JSON File")

	defer jsonFile.Close()

	byteJson, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteJson, &users)

	fmt.Println(users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("Name: ", users.Users[i].Name)
		fmt.Println("Phone: ", users.Users[i].Phone)
		fmt.Println("Email: ", users.Users[i].Email)
	}
}
