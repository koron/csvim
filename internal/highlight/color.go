package highlight

import (
	"fmt"
	"io"
)

// ColorNr is color number for terminal.
type ColorNr string

func (cn ColorNr) isValid() bool {
	return cn != ""
}

func (cn *ColorNr) merge(src ColorNr) {
	if src.isValid() {
		*cn = src
	}
}

func (cn ColorNr) writeTo(w io.Writer, label string) error {
	if !cn.isValid() {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=%s", label, cn)
	return err
}

// ColorName is color name for GUI.
type ColorName string

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
	_, err := fmt.Fprintf(w, " %s=%s", label, cn)
	return err
}
