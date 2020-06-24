package highlight

import (
	"fmt"
	"io"
)

// ColorNr is color number for terminal.
type ColorNr string

func (cn ColorNr) writeTo(w io.Writer, label string) error {
	if cn == "" {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=%s", label, cn)
	return err
}

// ColorName is color name for GUI.
type ColorName string

func (cn ColorName) writeTo(w io.Writer, label string) error {
	if cn == "" {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=%s", label, cn)
	return err
}
