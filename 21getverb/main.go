package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web requests : GET")
	performGetRequest("http://localhost:8080")
}

func performGetRequest(myUrl string) {
	response, err := http.Get(myUrl)
	if err != nil {
		print(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	//Reading the contents of the webpage
	content, _ := ioutil.ReadAll(response.Body)
	// {
	fmt.Println(string(content))
	// }

	//or

	// {
	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)
	fmt.Println("Byte Count: ", bytecount)
	fmt.Println(responseString.String())
	// }
}
