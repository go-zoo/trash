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
	AlreadyExistErr    = "ALREADY EXIST"
	AuthErr            = "AUTHENTIFICATION ERROR"
	BadRequestErr      = "BAD REQUEST"
	DbErr              = "DATABASE ERROR"
	DeserializationErr = "DESERIALIZATION ERROR"
	FileErr            = "FILE ERROR"
	GenericErr         = "GENERIC ERROR"
	InvalidJSONErr     = "INVALID JSON"
	InvalidDataErr     = "INVALID DATA "
	InternalEr         = "INTERNAL SERVER ERROR"
	NotFoundErr        = "NOT FOUND"
)

var logg = log.New(os.Stdout, "[TRASH] ", 0)

// Err is the default interface for trash
type Err interface {
	Send(io.Writer) Err
	Log() Err
	Text() string
}

// HTTPErr is a upgrade from Err adding HTTP Send and Log
type HTTPErr interface {
	Err
	SendHTTP(http.ResponseWriter, int) HTTPErr
	LogHTTP(*http.Request) HTTPErr
}

// Error is the default error type
type Error struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// NewErr generate a new Err
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

// NewHTTPErr generate a new HTTPErr
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
