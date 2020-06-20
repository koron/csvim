package highlight

import "io"

type writer struct {
	w   io.Writer
	ws  func(s string) (int, error)
	err error

	wrote bool
}

func newWriter(base io.Writer) *writer {
	w := &writer{w: base}
	if sw, ok := base.(io.StringWriter); ok {
		w.ws = sw.WriteString
	} else {
		w.ws = func(s string) (int, error) {
			return w.Write([]byte(s))
		}
	}
	return w
}

func (w *writer) Write(b []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	n, err := w.w.Write(b)
	if n > 0 && !w.wrote {
		w.wrote = true
	}
	if err != nil {
		w.err = err
		return n, w.err
	}
	return n, nil
}

func (w *writer) WriteString(s string) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	n, err := w.ws(s)
	if n > 0 && !w.wrote {
		w.wrote = true
	}
	if err != nil {
		w.err = err
		return n, w.err
	}
	return n, nil
}

func (w writer) writeLabel(s string) {
	if w.wrote {
		w.WriteString(" " + s + "=")
		return
	}
	w.WriteString(s + "=")
}
