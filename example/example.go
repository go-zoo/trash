package main

import (
	"net/http"

	"github.com/go-zoo/trash"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		trash.NewErr(trash.NOT_FOUND_ERR, "Test Error", "json").SendHTTP(rw, 404).Log()
	})
	http.ListenAndServe(":8080", nil)
}
