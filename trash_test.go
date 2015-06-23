package trash

import (
	"bytes"
	"testing"
)

func TestErrLog(t *testing.T) {
	buffer := bytes.NewBufferString("")
	errx := NewErr("GENERIC ERROR", "Test error", "json")
	errx.Send(buffer)
	if buffer.Len() > 1 {
		t.Fail()
	}
}
