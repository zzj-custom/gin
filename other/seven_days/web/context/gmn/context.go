package gmn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ContentTypeString = "text/plain"
	ContentTypeJson   = "application/json"
	ContentTypeHtml   = "text/html"
)

type H map[string]any

type Context struct {
	// 原始数据
	Write http.ResponseWriter
	Req   *http.Request

	// 请求数据
	Method string
	Path   string

	// 返回数据
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Write:  w,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Write.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Write.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...any) {
	c.SetHeader("Content-Type", ContentTypeString)
	c.Write.WriteHeader(code)
	c.Write.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Json(code int, obj any) {
	c.SetHeader("Content-Type", ContentTypeJson)
	c.Status(code)
	encoder := json.NewEncoder(c.Write)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Write, err.Error(), code)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Write.Write(data)
}

func (c *Context) Html(code int, html string) {
	c.SetHeader("Content-Type", ContentTypeHtml)
	c.Status(code)
	c.Write.Write([]byte(html))
}
