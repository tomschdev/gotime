package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/tomschdev/gotime/gophercises/ex2"
)

func main() {
	mux := defaultMux() // handles requests with a good old "hello, world"
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/tomschdev/gotime/gophercises/ex2",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	// mapHandler should divert handling of above URLs to urlshort.MapHandler, else to mux
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yaml :=
		`
- path: /urlshort
  url: https://github.com/tomschdev/gotime/gophercises/ex2
- path: /go-pracs
  url: https://github.com/tomschdev/gotime/go-practicals
`
	// yamlHandler should divert handling of above URLS to urlshort.YAMLHandler, else to mapHandler
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	_ = yamlHandler
	json := `
[
	{
		"path":"/jsonpath",
		"url":"https://www.linkedin.com/in/thomas-scholtz-32701/"
	}
]
	`
	jsonHandler, err := urlshort.JSONHandler([]byte(json), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("starting server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
