package lion

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Req  *http.Request
	Resp http.ResponseWriter
}

func (c *Context) String(code int, value string) {
	c.Resp.WriteHeader(code)
	c.Resp.Write([]byte(value))
}

func (c *Context) JSON(code int, value any) {
	c.Resp.WriteHeader(code)
	c.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Resp).Encode(value)
}
