// Key Moments
// 1. Parsing values from url
// 2. Extract Query Parameters
// 3. Separate Parameters in Query
// 4. Construct Custom URL

package main

import (
	"fmt"
	"net/url"
)

const myUrl = "http://lco.dev:3030/learn?coursename=reactjs&paymentid=gibberish123"

func main() {
	fmt.Println("Handling URLs")
	fmt.Println(myUrl)

	// Parsing (means extracting information from the url)
	result, _ := url.Parse(myUrl) //treats the string as an actual url
	fmt.Println(result.Scheme)    //Extracts Protocol Scheme
	fmt.Println(result.Host)      //Extracts Domain Name
	fmt.Println(result.Port())    //Extracts Port
	fmt.Println(result.Path)      //Extracts Path
	fmt.Println(result.RawQuery)  //Extracts Query (parameters after ?)

	//Queries Parameters in URL
	queryParameters := result.Query()
	fmt.Printf("Type of Queries: %T\n", queryParameters)
	fmt.Println("Queries: ", queryParameters)
	fmt.Println("Individual Parameter Value: ", queryParameters["coursename"])

	//Separate the Parameters of the Query
	for _, val := range queryParameters {
		fmt.Println(val)
	}

	//Construct a URL
	constructUrl := &url.URL{
		Scheme:  "http",
		Host:    "youtube.com",
		Path:    "/subscriptions",
		RawPath: "user=srikarthick",
	}

	fmt.Println(constructUrl.String())
}
