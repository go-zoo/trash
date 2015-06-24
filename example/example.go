package main

import (
	"errors"
	"net/http"

	"github.com/go-zoo/trash"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		dump := trash.NewDump(rw, "json")
		err := errors.New("test")

		dump.Catch(err.Error())

		errr := errors.New("BLABLABLA")
		dump.Catch(errr.Error())
	})
	http.ListenAndServe(":8080", nil)
}
