package main

import (
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func HandleFilm(w http.ResponseWriter, r *http.Request) {
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

func HandleAddFilm(w http.ResponseWriter, r *http.Request) {
	log.Print("HTMX request received")
	log.Print(r.Header.Get("HX-Request"))
}

func main() {
	http.HandleFunc("/", HandleFilm)
	http.HandleFunc("/add-film/", HandleAddFilm)
	http.ListenAndServe("192.168.206.1:9001", nil)
}
