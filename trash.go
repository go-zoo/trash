package trash

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ALREADY_EXIST_ERR   = "ALREADY EXIST"
	AUTH_ERR            = "AUTHENTIFICATION ERROR"
	BAD_REQUEST_ERR     = "BAD REQUEST"
	DB_ERR              = "DATABASE ERROR"
	DESERIALIZATION_ERR = "DESERIALIZATION ERROR"
	FILE_ERR            = "FILE ERROR"
	GENERIC_ERR         = "GENERIC ERROR"
	INVALID_JSON_ERR    = "INVALID JSON"
	INVALID_DATA_ERR    = "INVALID DATA "
	INTERNAL_ERR        = "INTERNAL SERVER ERROR"
	NOT_FOUND_ERR       = "NOT FOUND"
)

var logg = log.New(os.Stdout, "[TRASH] ", 0)

type Err interface {
	Send(io.Writer) Err
	Log() Err
	Text() string
}

type HTTPErr interface {
	Err
	SendHTTP(http.ResponseWriter, int) HTTPErr
	LogHTTP(*http.Request) HTTPErr
}

type Error struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewErr(err string, message string, format string) Err {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	switch format {
	case "json":
		return JsonErr{Error: Error{checksum, err, message, 0}}
	case "xml":
		return XmlErr{Error: Error{checksum, err, message, 0}}
	default:
		return nil
	}
}

func NewHTTPErr(err string, message string, format string) HTTPErr {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	switch format {
	case "json":
		return JsonErr{Error: Error{checksum, err, message, 0}}
	case "xml":
		return XmlErr{Error: Error{checksum, err, message, 0}}
	default:
		return nil
	}
}
