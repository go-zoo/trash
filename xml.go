package trash

import (
	"encoding/xml"
	"io"
)

type XmlErr struct {
	Error `xml:"error"`
}

func (x XmlErr) Send(w io.Writer) Err {
	xml.NewEncoder(w).Encode(x)
	return x
}

func (x XmlErr) Log() Err {
	logg.Printf("\x1b[%s%s\x1b[0m \"%s\" ", "41m", x.Type, x.Message)
	return x
}

func (x XmlErr) Text() string {
	return x.Message
}
