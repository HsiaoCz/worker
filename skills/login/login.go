package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	ngrokSkip     = "ngrok-skip-browser-warning"
	ngrokValue    = "69420"
	getMessageUrl = "http://127.0.0.1:9001/api/meetingRoomList"
	// getMessageUrl = "https://mp.weixin.qq.com/api/meetingRoomList"
)

func LoginWithSkip(url string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "1")
	req.Header.Add(ngrokSkip, ngrokValue)
	query := req.URL.Query()
	query.Add("page", "1")
	query.Add("page_size", "10")
	req.URL.RawQuery = query.Encode()
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(req.URL)
	fmt.Println(string(body))
}

func LoginNotWithSkip(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "1")
	query := req.URL.Query()
	query.Add("page", "1")
	query.Add("page_size", "10")
	req.URL.RawQuery = query.Encode()
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(req.URL)
	fmt.Println(string(body))
}

func main() {
	// LoginWithSkip
	// LoginWithSkip(getMessageUrl)
	// LoginNotWithSkip
	LoginNotWithSkip(getMessageUrl)
}
