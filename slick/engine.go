package slick

import (
	"fmt"
	"net/http"
)

type Slick struct {
	router map[string]Handlefunc
	srv    http.Server
}

func New() *Slick {
	return &Slick{
		router: make(map[string]Handlefunc),
	}
}

func (s *Slick) addRoute(method string, pattern string, handler Handlefunc) {
	key := method + "-" + pattern
	s.router[key] = handler
}

func (s *Slick) GET(pattern string, handler Handlefunc) {
	s.addRoute("GET", pattern, handler)
}

func (s *Slick) POST(pattern string, handler Handlefunc) {
	s.addRoute("POST", pattern, handler)
}

func (s *Slick) Start(addr string) (err error) {
	s.srv = http.Server{
		Handler: s,
		Addr:    addr,
	}
	return s.srv.ListenAndServe()
}

func (s *Slick) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", r.URL)
	}
}
