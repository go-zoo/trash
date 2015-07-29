package main

import (
	"errors"
	"net/http"

	"github.com/go-zoo/trash"
)

func main() {
	// Create a new trash
	t := trash.New(trash.DefaultLogger, "xml")

	http.HandleFunc("/xml", func(rw http.ResponseWriter, req *http.Request) {
		// Err with config
		t.NewHTTPErr(trash.GenericErr, "Test Error").SendHTTP(rw, 404).LogHTTP(req)
	})

	http.HandleFunc("/json", func(rw http.ResponseWriter, req *http.Request) {
		// Standalone JsonErr
		trash.NewJSONErr(trash.BadRequestErr, errors.New("Json Err !!")).SendHTTP(rw, 401).LogHTTP(req)
	})

	http.HandleFunc("/local", func(rw http.ResponseWriter, req *http.Request) {
		t.NewErr(trash.BadRequestErr, "local error").Log()
	})

	http.ListenAndServe(":8080", nil)
}
