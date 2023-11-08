package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// PostForm returns the first value for the named component of the query
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query the first value associated with the given key
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status set StatusCode
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(c.StatusCode)
}

func (c *Context) SetHeader(key, value string) {
	c.Req.Header.Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "test/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, data string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(data))
}
