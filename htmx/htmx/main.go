package main

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
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
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	log.Print(title)
	log.Print(director)
	htmlstr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'> %s  - %s  </li>", title, director)
	tmpl, err := template.New("t").Parse(htmlstr)
	if err != nil {
		slog.Error("template parse fialed:", err)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", HandleFilm)
	http.HandleFunc("/add-film/", HandleAddFilm)
	http.ListenAndServe("192.168.206.1:9001", nil)
}
