package main

import (
	"fmt"
	"net/http"
)

// for simple/static pages
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World")
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./static/hello.html")
}

