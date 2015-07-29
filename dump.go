package trash

import (
	"io"
	"sync"
)

// Dump data structure
type Dump struct {
	trash   map[int]*Trash
	errChan chan genErr
}

type writer struct {
	*sync.RWMutex
	io.Writer
}

// NewDump generate a new dump instance and start listening for incoming err.
func NewDump(w io.Writer, trash ...*Trash) *Dump {
	d := &Dump{trash: make(map[int]*Trash), errChan: make(chan genErr)}
	for k, v := range trash {
		v.dump = d
		d.trash[k] = v
	}
	go d.listen(writer{&sync.RWMutex{}, w})
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
