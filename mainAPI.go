package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type author struct {
	Name string
	FavoriteBook string
}

var favoriteAuthors [6]author = [6]author{
	{ Name: "N. K. Jemisin", FavoriteBook: "The Fifth Season"},
	{ Name: "Noam Chomsky", FavoriteBook: "Who Rules the World?"},
	{ Name: "Nnedi Okarafor", FavoriteBook: "Who Fears Death"},
	{ Name: "Robert Jordan", FavoriteBook: "The Dragon Reborn"},
	{ Name: "Frederick Forsyth", FavoriteBook: "The Day of the Jackal"},
	{ Name: "Octavia Butler", FavoriteBook: "Kindred"},
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/*", fs)

	http.HandleFunc("/", apiHandler)

	http.HandleFunc("/authors", authorHandler)

	http.ListenAndServe(":8080", nil)
}

func authorHandler(w http.ResponseWriter, req *http.Request) {
	favoriteAuthorsJSON, err := json.Marshal(favoriteAuthors)

	if err != nil {
		panic("Could not marshal JSON")
	}

	fmt.Fprintf(w, string(favoriteAuthorsJSON))
}

func apiHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/index.html")
}