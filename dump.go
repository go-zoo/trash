package trash

import (
	"io"
)

type Dump struct {
	Errs   map[string]Err
	Format string
	Writer io.Writer
}

func NewDump(w io.Writer, f string) *Dump {
	return &Dump{make(map[string]Err), f, w}
}

func (d *Dump) Insert(t string, err string) {
	d.Errs[err] = NewErr(t, err, d.Format)
}

func (d *Dump) Remove(err Err) {
	delete(d.Errs, err.Text())
}

func (d *Dump) Get(err string) Err {
	return d.Errs[err]
}

func (d *Dump) Catch(err string) {
	if d.Errs[err] == nil {
		d.Errs[err] = NewErr(err, err, d.Format)
	}
	d.Errs[err].Send(d.Writer)
	d.Errs[err].Log()
}
