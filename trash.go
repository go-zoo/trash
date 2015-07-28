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

type Trash struct {
	Type   string
	Format func(string, string) Err
	Logger *log.Logger
}

func New(logger *log.Logger, format string) *Trash {
	t := &Trash{Logger: logger}
	switch format {
	case "json":
		t.Type = "json"
		t.Format = t.jsonErr
	case "xml":
		t.Type = "xml"
		t.Format = t.xmlErr
	}
	return t
}

func (t *Trash) NewErr(err string, message string) Err {
	er := t.Format(err, message)
	return er
}

func (t *Trash) NewHTTPErr(err string, message string) HTTPErr {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	switch t.Type {
	case "json":
		return JsonErr{Logger: t.Logger, Error: Error{checksum, err, message, 0}}
	case "xml":
		return XmlErr{Logger: t.Logger, Error: Error{checksum, err, message, 0}}
	}
	return nil
}

// NewErr generate a new Err
func (t *Trash) jsonErr(err string, message string) Err {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	return JsonErr{Logger: t.Logger, Error: Error{checksum, err, message, 0}}
}

// NewHTTPErr generate a new HTTPErr
func (t *Trash) xmlErr(err string, message string) Err {
	checksum := base64.StdEncoding.EncodeToString([]byte(time.Now().String()))
	return XmlErr{Logger: t.Logger, Error: Error{checksum, err, message, 0}}
}
