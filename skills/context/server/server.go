package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/HsiaoCz/worker/skills/context/google"
	"github.com/HsiaoCz/worker/skills/context/userip"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var cancel context.CancelFunc

	timeout, err := time.ParseDuration(r.FormValue("timeout"))
	if err != nil {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	}
	defer cancel()

	query := r.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
		return
	}
	// 解析ip
	userIP, err := userip.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 将ip地址存储到ctx中
	ctx = userip.NewContext(ctx, userIP)
	// 调用搜索api获取查询结果
	start := time.Now()
	results, err := google.Search(ctx, query)
	if err != nil {
		return
	}
	elapsed := time.Since(start)
	json.NewEncoder(w).Encode(map[string]any{
		"Results": results,
		"Time":    elapsed,
	})
}
