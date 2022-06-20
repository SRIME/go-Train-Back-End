// Key Moments
// 1. make a fake request body using json
// 2. post the request and store it in response
// 3. convert the response to readable format

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("-----How to make post requeset using fake json data-----")

	const myurl = "http://localhost:8080/post"
	performPostJsonRequest(myurl)

}

func performPostJsonRequest(myurl string) {
	//1. create fake json payload
	requestBody := strings.NewReader(`
		{
			"name" : "Sri karthick",
			"age" : 20,
			"prog. language" : "go"
		}
	`)

	// 2. Post the request and store the response
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 3. Read and Convert the response to readable format
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
