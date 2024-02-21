package userip

import (
	"context"
	"net"
	"net/http"
)

// 非导出类型，避免和其他包的 key 冲突

type key int

const userIPKey key = 0

// 将ip写入context，返回生成的 child context

func NewContext(ctx context.Context, userIP net.IP) context.Context {

	return context.WithValue(ctx, userIPKey, userIP)

}

func FromContext(ctx context.Context) (net.IP, bool) {

	userIP, ok := ctx.Value(userIPKey).(net.IP)

	return userIP, ok

}

// 获取ip地址
func FromRequest(r *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	return net.ParseIP(ip), err
}
