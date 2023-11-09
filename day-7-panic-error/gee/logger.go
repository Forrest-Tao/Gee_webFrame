package gee

import (
	"log"
	"time"
)

func Logger() func(c *Context) {
	return func(c *Context) {
		//start time
		t := time.Now()
		// Process request
		c.Next()
		//// Calculate resolution time
		log.Printf("[%d] %s %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
