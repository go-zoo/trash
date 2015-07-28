package main

import (
	"net/http"

	"github.com/go-zoo/trash"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		trash.NewHTTPErr(trash.GenericErr, "Test Error", "json").SendHTTP(rw, 404).LogHTTP(req)
	})
	http.ListenAndServe(":8080", nil)
}
