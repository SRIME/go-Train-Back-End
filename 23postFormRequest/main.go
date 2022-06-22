package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Perform Post Form Request")
	const myurl = "http://localhost:8080/postform"
	performPostFormRequest(myurl)
}

func performPostFormRequest(myurl string) {
	data := url.Values{}
	data.Add("Name", "Sri karthick")
	data.Add("Ph no.", "9632207512")
	data.Add("Fav Technology", "React and go")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
