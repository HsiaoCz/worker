package wss

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/worker/binance/wssubstream/data"
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

func GetAllBitCoinsInfo(url string) string {
	postData := &data.PostGetAll{
		ReceWindow: 1,
		TimeStamp:  time.Now().UTC().UnixMilli(),
	}
	pd, err := json.Marshal(postData)
	if err != nil {
		log.Println("marshal data err:", err)
		return ""
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(pd))
	if err != nil {
		log.Println("create new request err:", err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(xma, apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client exec request err:", err)
		return ""
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("read resp.Body err:", err)
		return ""
	}
	return string(result)
}
