package trash

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

type JsonErr struct {
	Logger  *log.Logger `json:"-"`
	errData `json:"error"`
}

// NewErr generate a new Err
func NewJSONErr(err string, message string) JsonErr {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	return JsonErr{errData: errData{checksum, err, message, 0}}
}

func (j JsonErr) Send(w io.Writer) Err {
	json.NewEncoder(w).Encode(j)
	return j
}

func (j JsonErr) SendHTTP(rw http.ResponseWriter, code int) HTTPErr {
	rw.Header().Set("Content-Type", "application/json")
	j.Code = code
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(j)

	return j
}

func (j JsonErr) Log() Err {
	var logger *log.Logger
	if j.Logger != nil {
		logger = j.Logger
	} else {
		logger = logg
	}
	logger.Printf("\x1b[%s%s\x1b[0m %s ", "41m", j.Type, j.Message)
	return j
}

func (j JsonErr) LogHTTP(req *http.Request) HTTPErr {
	var logger *log.Logger
	if j.Logger != nil {
		logger = j.Logger
	} else {
		logger = logg
	}
	if runtime.GOOS != "windows" {
		logger.Printf("\x1b[%s%s\x1b[0m %s (%s %s %s)", "41m", j.errData.Type, j.errData.Message, req.Method, req.RemoteAddr, req.RequestURI)
	} else {
		logger.Printf("!%s! %s (%s %s %s)", j.errData.Type, j.errData.Message, req.Method, req.RemoteAddr, req.RequestURI)
	}
	return j
}

func (j JsonErr) Error() string {
	return j.Message
}
