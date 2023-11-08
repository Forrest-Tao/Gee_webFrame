package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/patten", pattenHandler)
	http.ListenAndServe(":9090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path=%q\n", r.URL.Path)
}

func pattenHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "header[%q]=%q\n", k, v)
	}
}
