package main

import "net/http"

// for single page apps (REACT, Angular)
func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/*", fs)

	http.HandleFunc("/", spaHandler);

	http.ListenAndServe(":8080", nil)
}

func spaHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/index.html")
}