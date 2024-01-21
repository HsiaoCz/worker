package wss

import (
	"io"
	"log"
	"net/http"
)

var (
	xma    = "X-MBX-APIKEY"
	apiKey = "x1z91yZRcvSbHgySqjwAeKxRZ5cneq8hgaA8c9BXgzoXdsKDAccBXd29v752tu6k"
)

func GetListenKey(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(xma, apiKey)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(body)
}
