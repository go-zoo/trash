package trash

import (
	"encoding/base64"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

type XmlErr struct {
	Logger  *log.Logger `xml:"-"`
	errData `xml:"error"`
}

// NewHTTPErr generate a new HTTPErr
func NewXMLPErr(err string, message string) XmlErr {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	return XmlErr{errData: errData{checksum, err, message, 0}}
}

func (x XmlErr) Send(w io.Writer) Err {
	xml.NewEncoder(w).Encode(x)
	return x
}

func (x XmlErr) SendHTTP(rw http.ResponseWriter, code int) HTTPErr {
	rw.Header().Set("Content-Type", "application/xml")
	x.Code = code
	rw.WriteHeader(code)
	xml.NewEncoder(rw).Encode(x)

	return x
}

func (x XmlErr) Log() Err {
	var logger *log.Logger
	if x.Logger != nil {
		logger = x.Logger
	} else {
		logger = logg
	}
	logger.Printf("\x1b[%s%s\x1b[0m %s ", "41m", x.Type, x.Message)
	return x
}

func (x XmlErr) LogHTTP(req *http.Request) HTTPErr {
	var logger *log.Logger
	if x.Logger != nil {
		logger = x.Logger
	} else {
		logger = logg
	}
	if runtime.GOOS != "windows" {
		logger.Printf("\x1b[%s%s\x1b[0m %s (%s %s %s)", "41m", x.errData.Type, x.errData.Message, req.Method, req.RemoteAddr, req.RequestURI)
	} else {
		logger.Printf("!%s! %s (%s %s %s)", x.errData.Type, x.errData.Message, req.Method, req.RemoteAddr, req.RequestURI)
	}
	return x
}

func (x XmlErr) Error() string {
	return x.Message
}
