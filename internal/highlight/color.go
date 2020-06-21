package highlight

import (
	"fmt"
	"io"
)

type ColorNr string

func (cn ColorNr) String() string {
	return string(cn)
}

func (cn ColorNr) isValid() bool {
	return cn != ""
}

func (cn ColorNr) writeTo(w io.Writer, label string) error {
	if !cn.isValid() {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=%s", label, cn.String())
	return err
}

func (cn *ColorNr) merge(src ColorNr) {
	if src.isValid() {
		*cn = src
	}
}

type ColorName string

func (cn ColorName) String() string {
	return string(cn)
}

func (cn ColorName) isValid() bool {
	return cn != ""
}

func (cn *ColorName) merge(src ColorName) {
	if src.isValid() {
		*cn = src
	}
}

func (cn ColorName) writeTo(w io.Writer, label string) error {
	if !cn.isValid() {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=%s", label, cn.String())
	return err
}
