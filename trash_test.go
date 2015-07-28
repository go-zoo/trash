package trash

import (
	"bytes"
	"log"
	"testing"
)

func TestTrash(t *testing.T) {
	buffer := bytes.NewBufferString("")
	trh := New(log.New(buffer, "", 0), "json")
	err := trh.NewErr(GenericErr, "test")
	err.Log()
	if buffer.Len() < 1 {
		t.Fail()
	}
}

func TestTrashWeirdTypo(t *testing.T) {
	buffer := bytes.NewBufferString("")
	trh := New(log.New(buffer, "", 0), "JsON")
	err := trh.NewErr(GenericErr, "test")
	err.Log()
	if buffer.Len() < 1 {
		t.Fail()
	}
}

func TestJSONErr(t *testing.T) {
	buffer := bytes.NewBufferString("")
	errx := NewJSONErr("GENERIC ERROR", "Test error")
	errx.Send(buffer)
	if buffer.Len() < 1 {
		t.Fail()
	}
}
