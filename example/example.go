package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-zoo/trash"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// Err with config
		t := trash.New(log.New(os.Stdout, "[$ TEST $] ", 0), "xml")
		t.NewHTTPErr(trash.GenericErr, "Test Error").SendHTTP(rw, 404).LogHTTP(req)

		// standalone Err
		trash.NewJSONErr(trash.BadRequestErr, "err002").SendHTTP(rw, 502).LogHTTP(req)
	})
	http.ListenAndServe(":8080", nil)
}
