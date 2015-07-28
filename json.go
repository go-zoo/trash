package trash

import (
	"encoding/json"
	"io"
	"net/http"
	"runtime"
)

type JsonErr struct {
	Error `json:"error"`
}

func (j JsonErr) Send(w io.Writer) Err {
	json.NewEncoder(w).Encode(j)
	return j
}

func (j JsonErr) SendHTTP(rw http.ResponseWriter, code int) Err {
	rw.Header().Set("Content-Type", "application/json")
	j.Code = code
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(j)

	return j
}

func (j JsonErr) Log() Err {
	logg.Printf("\x1b[%s%s\x1b[0m \"%s\" ", "41m", j.Type, j.Message)
	return j
}

func (j JsonErr) LogHTTP(req *http.Request) Err {
	if runtime.GOOS != "windows" {
		logg.Printf("\x1b[%s%s\x1b[0m %s (%s %s %s)", "41m", j.Error.Type, j.Error.Message, req.Method, req.RemoteAddr, req.RequestURI)
	} else {
		logg.Printf("!%s! %s (%s %s %s)", j.Error.Type, j.Error.Message, req.Method, req.RemoteAddr, req.RequestURI)
	}
	return j
}

func (j JsonErr) Text() string {
	return j.Message
}
