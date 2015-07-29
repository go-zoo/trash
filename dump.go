package trash

import (
	"io"
	"sync"
)

type Dump struct {
	trash   map[int]*Trash
	errChan chan genErr
}

type writer struct {
	sync.RWMutex
	io.Writer
}

func NewDump(w io.Writer, trash ...*Trash) *Dump {
	d := &Dump{trash: make(map[int]*Trash), errChan: make(chan genErr)}
	for k, v := range trash {
		v.dump = d
		d.trash[k] = v
	}
	go d.listen(writer{Writer: w})
	return d
}

func (d *Dump) listen(w writer) {
	for {
		select {
		case err := <-d.errChan:
			w.Lock()
			w.Write([]byte(err.FormatErr()))
			w.Unlock()
		default:
			continue
		}
	}
}
