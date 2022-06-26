// Creates a fixed template which is used by all pages and adds some logic

package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	fmt.Println(tpl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Successfully started at : localhost:" + port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandlers)
	http.ListenAndServe(":"+port, mux)
}

func indexHandlers(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

//go install github.com/cosmtrek/air@latest
// install above repo and run $ air to automatically update the changes done in both go as well as html files
