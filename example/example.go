package main

import (
	"net/http"

	"github.com/go-zoo/trash"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		trash.NewErr("TEST", "test error", "xml").Send(rw).Log()
	})
	http.ListenAndServe(":8080", nil)
}
