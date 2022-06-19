// Key moments:
// 1. get response
// 2. read the response and web contents
// 3. close response

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("-----Handling web Request-----")

	response, err := http.Get(url) //1. get the response from web(it brings the address of the web response)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //3. Callers responsibility for closing the response

	fmt.Printf("Response Type %T\n", response)
	data, err := ioutil.ReadAll(response.Body) //2. using ioutil read the contents of the website
	if err != nil {
		panic(err)
	}
	fmt.Println("Response: ", string(data))
}
