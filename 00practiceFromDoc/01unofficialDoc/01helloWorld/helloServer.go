// Reference for this practice set: https://freshman.tech/web-development-with-go/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World")) //Writes to response body
}

func main() {
	// {
	err := godotenv.Load() // default loading of port from env file(optional)
	if err != nil {
		log.Println("Error loading env file")
	}
	// }

	//main program
	port := os.Getenv("PORT") //gets the default free port available in the system
	if port == "" {           //if it fails to get default port
		port = "3000"
	}
	fmt.Println("sucesfullly started at port ", port)
	mux := http.NewServeMux() //create new request multiplexer

	mux.HandleFunc("/", indexHandler)  //register http request by using HandleFunc
	http.ListenAndServe(":"+port, mux) //starts the server on port
}
