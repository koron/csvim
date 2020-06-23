package highlight

import (
	"fmt"
	"io"
)

type TermColor interface {
	TermColor() ColorNr
}

type GUIColor interface {
	GUIColor() ColorName
}

type Color interface {
	TermColor
	GUIColor
}

type ColorNr string

func (cn *ColorNr) Set(c TermColor) {
	if c == nil {
		return
	}
	*cn = c.TermColor()
}

func (cn ColorNr) TermColor() ColorNr {
	return cn
}

func (cn ColorNr) isValid() bool {
	return cn != ""
}

func (cn ColorNr) writeTo(w io.Writer, label string) error {
	if !cn.isValid() {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=%s", label, cn)
	return err
}

func (cn *ColorNr) merge(src ColorNr) {
	if src.isValid() {
		*cn = src
	}
}

type ColorName string

func (cn *ColorName) Set(c GUIColor) {
	if c == nil {
		return
	}
	*cn = c.GUIColor()
}

func (cn ColorName) GUIColor() ColorName {
	return cn
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
	_, err := fmt.Fprintf(w, " %s=%s", label, cn)
	return err
}
