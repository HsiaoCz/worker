package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello My man")
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "Jhon Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)
	http.ListenAndServe("127.0.0.1:9001", nil)
}
