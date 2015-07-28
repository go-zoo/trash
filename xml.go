package trash

import (
	"encoding/xml"
	"io"
	"net/http"
	"runtime"
)

type XmlErr struct {
	Error `xml:"error"`
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

func (x XmlErr) LogHTTP(req *http.Request) HTTPErr {
	if runtime.GOOS != "windows" {
		logg.Printf("\x1b[%s%s\x1b[0m %s (%s %s %s)", "41m", x.Error.Type, x.Error.Message, req.Method, req.RemoteAddr, req.RequestURI)
	} else {
		logg.Printf("!%s! %s (%s %s %s)", x.Error.Type, x.Error.Message, req.Method, req.RemoteAddr, req.RequestURI)
	}
	return x
}

func (x XmlErr) Log() Err {
	logg.Printf("\x1b[%s%s\x1b[0m \"%s\" ", "41m", x.Type, x.Message)
	return x
}

func (x XmlErr) Text() string {
	return x.Message
}
