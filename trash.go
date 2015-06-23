package trash

import (
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
)

const (
	DB_ERR            = "DB ERROR"
	AUTH_ERR          = "AUTHENTIFICATION ERROR"
	BAD_REQUEST_ERR   = "BAD REQUEST"
	INVALID_JSON_ERR  = "INVALID JSON"
	INTERNAL_ERR      = "INTERNAL SERVER ERROR"
	ALREADY_EXIST_ERR = "ALREADY EXIST ERROR"
	NOT_FOUND_ERR     = "NOT FOUND"
)

type Err interface {
	Send(io.Writer) Err
	Log() Err
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
	log.Printf("\x1b[%s%s\x1b[0m \"%s\" ", "41m", j.Type, j.Message)
	return j
}

type XmlErr struct {
	Error `xml:"error"`
}

func (x XmlErr) Send(w io.Writer) Err {
	xml.NewEncoder(w).Encode(x)
	return x
}

func (x XmlErr) Log() Err {
	log.Printf("\x1b[%s%s\x1b[0m \"%s\" ", "41m", x.Type, x.Message)
	return x
}

func NewErr(err string, message string, format string) Err {
	checksum := md5.Sum([]byte(err + message + format))
	switch format {
	case "json":
		return JsonErr{Error: Error{string(checksum[:]), err, message, 0}}
	case "xml":
		return XmlErr{Error: Error{string(checksum[:]), err, message, 0}}
	default:
		return nil
	}
}
