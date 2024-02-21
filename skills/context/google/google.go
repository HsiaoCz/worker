package google

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/worker/skills/context/userip"
)

type Result struct {
	Title string
	URL   string
}

type Results []Result

func Search(ctx context.Context, query string) (Results, error) {

	// Prepare the Google Search API request.

	req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)

	if err != nil {

		return Results{}, err

	}

	q := req.URL.Query()

	q.Set("q", query)

	// 从 ctx 解出 user ip

	if userIP, ok := userip.FromContext(ctx); ok {

		q.Set("userip", userIP.String())

	}

	req.URL.RawQuery = q.Encode()

	var results Results

	// 请求搜索 api

	err = httpDo(ctx, req, func(resp *http.Response, err error) error {

		if err != nil {

			return err

		}

		defer resp.Body.Close()

		// 解析 response

		// https://developers.google.com/web-search/docs/#fonje

		var data struct {
			ResponseData struct {
				Results []struct {
					TitleNoFormatting string

					URL string
				}
			}
		}

		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {

			return err

		}

		for _, res := range data.ResponseData.Results {

			results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})

		}

		return nil

	})

	// httpDo 会等待提供的 closure 返回

	return results, err

}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {

	// 单起goroutine发送请求

	c := make(chan error, 1)

	req = req.WithContext(ctx)

	go func() { c <- f(http.DefaultClient.Do(req)) }()

	select {

	case <-ctx.Done():

		<-c // 等待 f 返回

		return ctx.Err()

	case err := <-c:

		return err

	}

}
