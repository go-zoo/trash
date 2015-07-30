package trash

import (
	"io"
	"net/http"
)

const (
	AlreadyExistErr       = "ALREADY EXIST"
	AuthErr               = "AUTHENTIFICATION ERROR"
	BadRequestErr         = "BAD REQUEST"
	DbErr                 = "DATABASE ERROR"
	DeserializationErr    = "DESERIALIZATION ERROR"
	FileErr               = "FILE ERROR"
	GenericErr            = "GENERIC ERROR"
	InvalidJSONErr        = "INVALID JSON"
	InvalidXMLErr         = "INVALID XML ERROR"
	InvalidDataErr        = "INVALID DATA "
	InternalErr           = "INTERNAL SERVER ERROR"
	NotFoundErr           = "NOT FOUND"
	UnauthorizedAccessErr = "UNAUTHORIZED ACCESS ERROR"
)

type genErr interface {
	Err
	HTTPErr
	FmtErr
}

// Err is the default interface for trash
type Err interface {
	Send(io.Writer) Err
	Log() Err
	Error() string
}

// HTTPErr is a upgrade from Err adding HTTP Send and Log
type HTTPErr interface {
	SendHTTP(http.ResponseWriter, int) HTTPErr
	LogHTTP(*http.Request) HTTPErr
}

// FmtErr is the Err formatting interface{}
type FmtErr interface {
	FormatErr() string
}

// Error is the default error type
type errData struct {
	ID      string `json:"id" xml:"id"`
	Type    string `json:"type" xml:"type"`
	Message string `json:"message" xml:"message"`
	Code    int    `json:"code" xml:"code"`
}
