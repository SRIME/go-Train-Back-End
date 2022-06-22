package main

import (
	"encoding/json"
	"fmt"
)

type userDetails struct {
	Name                string   `json:"user name"`
	Email               string   `json:"email"`
	Password            string   `json:"-"`
	InterestsInLanguage []string `json:"Interest in languages,omitempty"` //prints value if it is not nil and omits the field if its nil
}

func main() {
	fmt.Println("Create JSON Data in go")
	encodeJson()
}

func encodeJson() {
	UserDetailsEncoded := []userDetails{
		{"Sri karthick", "srikarthick.kanda@gmail.co,m", "qwerty123", []string{"react", "go", "java", "script"}},
		{"abc", "abc123@gmail.com", "qwsef12", nil},
		{"xyz", "xyz@gmail.com", "gakjk123", []string{"c++", "ruby", "rust"}},
	}

	// finalData, _ := json.Marshal(UserDetailsEncoded)
	finalData, _ := json.MarshalIndent(UserDetailsEncoded, "", "\t")
	fmt.Println(string(finalData))
}
