package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("request read err:", err)
		return
	}
	userL := new(UserLogin)
	if err := json.Unmarshal(body, userL); err != nil {
		slog.Error("json unmarshal err:", err)
		return
	}
	user := &UserLogin{
		Username: "bob",
		Password: "12334a",
	}
	if userL.Username != user.Username || userL.Password != user.Password {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{
			"Message": "请检查用户名或密码",
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
	http.HandleFunc("/login", Login)
	http.ListenAndServe("127.0.0.1:9001", nil)
}
