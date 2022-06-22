package main

import (
	"encoding/json"
	"fmt"
)

type userDetails struct {
	Name                string   `json:"user name"` //Aliases doesn't work in decoding
	Email               string   `json:"email"`
	Password            string   `json:"-"`
	InterestsInLanguage []string `json:"Interest in languages,omitempty"`
}

func main() {
	fmt.Println("Decoding the JSON data Locally(Intro to decode from web)")
	decodeJsonLocal()
}

func decodeJsonLocal() {
	jsonData := []byte(`
	{
		"user name": "Sri karthick",
		"email": "srikarthick.kanda@gmail.com",
		"Interest in languages": ["react","go","java","script"]
	}
	`)

	var UserDetailsDecoded userDetails

	isValid := json.Valid(jsonData)

	if isValid {
		fmt.Println("The JSON data was Valid")
		json.Unmarshal(jsonData, &UserDetailsDecoded)
		fmt.Printf("%#v\n", UserDetailsDecoded)
		// fmt.Println(UserDetailsDecoded)
	} else {
		fmt.Println("JSON Data Was not valid")
	}

	// Treat the values as key value pair

	var keyValueFormat map[string]interface{}

	json.Unmarshal(jsonData, &keyValueFormat)
	fmt.Printf("%#v\n", keyValueFormat)

	for key, value := range keyValueFormat {
		fmt.Printf("Key is %v, Value is %v and DataType is %T\n", key, value, value)
	}
}
