package tee

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router map[string]Handlefunc
	srv    http.Server
}

func New() *Engine {
	return &Engine{
		router: make(map[string]Handlefunc),
	}
}

func (e *Engine) addRoute(method string, pattern string, handler Handlefunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) GET(pattern string, handler Handlefunc) {
	e.addRoute("GET", pattern, handler)
}
func (e *Engine) POST(pattern string, handler Handlefunc) {
	e.addRoute("GET", pattern, handler)
}
func (e *Engine) PUT(pattern string, handler Handlefunc) {
	e.addRoute("GET", pattern, handler)
}
func (e *Engine) DELETE(pattern string, handler Handlefunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) Start(addr string) error {
	e.srv = http.Server{
		Addr:    addr,
		Handler: e,
	}
	return e.srv.ListenAndServe()
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
