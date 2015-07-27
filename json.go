package trash

import (
	"encoding/json"
	"io"
)

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
