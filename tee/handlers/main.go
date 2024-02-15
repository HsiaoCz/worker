package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("read request body err:", err)
		return
	}
	user := new(User)
	if err := json.Unmarshal(body, user); err != nil {
		slog.Error("json unmarshal err:", err)
		return
	}
	if user.Username != "zhangsan" || user.Password != "12306" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"Message": "登录失败",
			"Code":    http.StatusBadRequest,
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"Message": "登录成功",
		"Code":    http.StatusOK,
	})
}

func main() {
	http.HandleFunc("/user/login", HandleLogin)
	http.ListenAndServe("127.0.0.1:9001", nil)
}
