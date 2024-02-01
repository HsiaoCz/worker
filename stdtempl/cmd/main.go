package main

import (
	"net/http"

	"github.com/HsiaoCz/worker/stdtempl/handler"
)

func main() {
	userCase := handler.Usercase{}
	http.HandleFunc("/user", userCase.Show)
	http.ListenAndServe("192.168.206.1:9008", nil)
}
