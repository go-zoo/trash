package trash

import (
	"io"
)

type Dump struct {
	trash   map[int]*Trash
	errChan chan genErr
}

func NewDump(w io.Writer, trash ...*Trash) *Dump {
	d := &Dump{trash: make(map[int]*Trash), errChan: make(chan genErr)}
	for k, v := range trash {
		v.dump = d
		d.trash[k] = v
	}
	go d.listen(w)
	return d
}

func (d *Dump) listen(w io.Writer) {
	for {
		select {
		case err := <-d.errChan:
			w.Write([]byte(err.FormatErr()))
		default:
			continue
		}
	}
}
