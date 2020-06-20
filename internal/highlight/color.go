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

type ColorName string

func (cn ColorName) String() string {
	return string(cn)
}

func (cn ColorName) isValid() bool {
	return cn != ""
}

func (cn ColorName) writeTo(w io.Writer, label string) error {
	if !cn.isValid() {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=%s", label, cn.String())
	return err
}

type Color struct {
	Nr   ColorNr
	Name ColorName
}

type Colors struct {
	Fg Color
	Bg Color
}
