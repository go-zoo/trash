package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/go-zoo/trash"
)

func main() {
	// Create a new trash
	file, _ := os.Create("err_log.txt")
	t := trash.New(trash.DefaultLogger, "xml")
	trash.NewDump(file, t)

	http.HandleFunc("/xml", func(rw http.ResponseWriter, req *http.Request) {
		// Err with config
		t.NewHTTPErr(trash.GenericErr, "Test Error").SendHTTP(rw, 404).LogHTTP(req)
	})

	http.HandleFunc("/json", func(rw http.ResponseWriter, req *http.Request) {
		// Standalone JsonErr
		trash.NewJSONErr(trash.BadRequestErr, errors.New("Json Err !!")).SendHTTP(rw, 401).LogHTTP(req)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		createError("error interface from" + r.RemoteAddr)
	})

	http.HandleFunc("/local", func(rw http.ResponseWriter, req *http.Request) {
		t.NewErr(trash.BadRequestErr, "local error").Log()
	})

	http.HandleFunc("/short", func(rw http.ResponseWriter, req *http.Request) {
		t.NewShortErr(trash.GenericErr, "short Err", rw)
	})

	http.ListenAndServe(":8080", nil)
}

func createError(m string) error {
	err := trash.NewJSONErr(trash.GenericErr, m)
	err.Log()
	return err
}
