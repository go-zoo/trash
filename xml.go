package trash

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

// XmlErr data structure
type XmlErr struct {
	Logger  *log.Logger `xml:"-"`
	errData `xml:"error"`
}

// NewXMLErr generate a new HTTPErr
func NewXMLErr(err string, message interface{}) XmlErr {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	return XmlErr{errData: errData{checksum, err, extractMessage(message), 0}}
}

// Send the Err to the provided io.Writer
func (x XmlErr) Send(w io.Writer) Err {
	xml.NewEncoder(w).Encode(x)
	return x
}

// SendHTTP send the Err to the HTTP Client
func (x XmlErr) SendHTTP(rw http.ResponseWriter, code int) HTTPErr {
	rw.Header().Set("Content-Type", "application/xml")
	x.Code = code
	rw.WriteHeader(code)
	xml.NewEncoder(rw).Encode(x)

	return x
}

// Log the default infos
func (x XmlErr) Log() Err {
	var logger *log.Logger
	if x.Logger != nil {
		logger = x.Logger
	} else {
		logger = DefaultLogger
	}
	if runtime.GOOS != "windows" {
		logger.Printf("\x1b[%s%s\x1b[0m %s ", "41m", x.Type, x.Message)
	} else {
		logger.Printf("[%s] %s", x.errData.Type, x.errData.Message)
	}
	return x
}

// LogHTTP log the default infos plus the request
func (x XmlErr) LogHTTP(req *http.Request) HTTPErr {
	var logger *log.Logger
	if x.Logger != nil {
		logger = x.Logger
	} else {
		logger = DefaultLogger
	}
	if runtime.GOOS != "windows" {
		logger.Printf("\x1b[%s%s\x1b[0m %s (%s %s %s)", "41m", x.errData.Type, x.errData.Message, req.Method, req.RemoteAddr, req.RequestURI)
	} else {
		logger.Printf("[%s] %s (%s %s %s)", x.errData.Type, x.errData.Message, req.Method, req.RemoteAddr, req.RequestURI)
	}
	return x
}

// Error return the Err message
func (x XmlErr) Error() string {
	return x.Message
}

func (x XmlErr) FormatErr() string {
	return fmt.Sprintf("%s [%s] %s %s \n", time.Now().String(), x.Type, x.Message, "xml")
}
