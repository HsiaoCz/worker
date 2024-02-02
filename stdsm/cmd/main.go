package main

import (
	"net/http"

	"github.com/HsiaoCz/worker/stdsm/controllers"
)

func main() {
	http.HandleFunc("/user", controllers.HandleUserShow)
	http.ListenAndServe("192.168.206.1:9009", nil)
}
