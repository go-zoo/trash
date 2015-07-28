package trash

import (
	"io"
	"net/http"
)

// Err is the default interface for trash
type Err interface {
	Send(io.Writer) Err
	Log() Err
	Error() string
}

// HTTPErr is a upgrade from Err adding HTTP Send and Log
type HTTPErr interface {
	Err
	SendHTTP(http.ResponseWriter, int) HTTPErr
	LogHTTP(*http.Request) HTTPErr
}

// Error is the default error type
type errData struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
