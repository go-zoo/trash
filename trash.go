package trash

import (
	"encoding/base64"
	"log"
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

type genErr interface {
	Err
	HTTPErr
}

// Trash defined the Trash data structure
type Trash struct {
	Type   string
	format interface{}
	//func(string, string) Err
	Logger *log.Logger
}

// New create a new Trash with the provided logger and data format
func New(logger *log.Logger, format string) *Trash {
	t := &Trash{Logger: logger}
	switch format {
	case "json":
		t.Type = "json"
		t.format = t.jsonErr
	case "xml":
		t.Type = "xml"
		t.format = t.xmlErr
	}
	return t
}

// NewErr generate a standard Err
func (t *Trash) NewErr(err string, message string) Err {
	er := t.format.(func(string, string) genErr) //(err, message)
	return er(err, message)
}

// NewHTTPErr generate a new HTTPErr
func (t *Trash) NewHTTPErr(err string, message string) HTTPErr {
	er := t.format.(func(string, string) genErr)
	return er(err, message)
}

// NewErr generate a new Err
func (t *Trash) jsonErr(err string, message string) genErr {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	return JsonErr{Logger: t.Logger, errData: errData{checksum, err, message, 0}}
}

// NewHTTPErr generate a new HTTPErr
func (t *Trash) xmlErr(err string, message string) genErr {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	return XmlErr{Logger: t.Logger, errData: errData{checksum, err, message, 0}}
}
