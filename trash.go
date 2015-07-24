package trash

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"
	"time"
)

const (
	DB_ERR              = "DATABASE ERROR"
	AUTH_ERR            = "AUTHENTIFICATION ERROR"
	BAD_REQUEST_ERR     = "BAD REQUEST"
	INVALID_JSON_ERR    = "INVALID JSON"
	INVALID_DATA_ERR    = "INVALID DATA "
	INTERNAL_ERR        = "INTERNAL SERVER ERROR"
	ALREADY_EXIST_ERR   = "ALREADY EXIST"
	NOT_FOUND_ERR       = "NOT FOUND"
	FILE_ERR            = "FILE ERROR"
	DESERIALIZATION_ERR = "DESERIALIZATION ERROR"
)

var logg = log.New(os.Stdout, "[TRASH] ", 0)

type Err interface {
	Send(io.Writer) Err
	Log() Err
	Text() string
}

type Error struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type JsonErr struct {
	Error `json:"error"`
}

func (j JsonErr) Send(w io.Writer) Err {
	json.NewEncoder(w).Encode(j)
	return j
}

func (j JsonErr) Log() Err {
	logg.Printf("\x1b[%s%s\x1b[0m \"%s\" ", "41m", j.Type, j.Message)
	return j
}

func (j JsonErr) Text() string {
	return j.Message
}

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

func NewErr(err string, message string, format string) Err {
	checksum := base64.StdEncoding.EncodeToString([]byte(err + time.Now().String()))
	switch format {
	case "json":
		return JsonErr{Error: Error{checksum, err, message, 0}}
	case "xml":
		return XmlErr{Error: Error{checksum, err, message, 0}}
	default:
		return nil
	}
}
