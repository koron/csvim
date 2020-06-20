package highlight

import (
	"fmt"
	"io"
)

var Command = "highlight"

type Attr string

const (
	None          Attr = "NONE"
	Bold               = "bold"
	Underline          = "underline"
	Undercurl          = "undercurl"
	Strikethrough      = "strikethrough"
	Reverse            = "reverse"
	Inverse            = "inverse"
	Italic             = "italic"
	Standout           = "standout"
	Nocombine          = "nocombine"
)

type AttrList []Attr

func (attrs AttrList) writeTo(w io.Writer, label string) error {
	if len(attrs) == 0 {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=", label)
	if err != nil {
		return err
	}
	for i, attr := range attrs {
		if i != 0 {
			_, err := io.WriteString(w, ",")
			if err != nil {
				return err
			}
		}
		_, err := io.WriteString(w, string(attr))
		if err != nil {
			return err
		}
	}
	return nil
}

type TermList []string

func (terms TermList) writeTo(w io.Writer, label string) error {
	if len(terms) == 0 {
		return nil
	}
	_, err := fmt.Fprintf(w, " %s=", label)
	if err != nil {
		return err
	}
	for i, term := range terms {
		if i != 0 {
			_, err := io.WriteString(w, ",")
			if err != nil {
				return err
			}
		}
		_, err := io.WriteString(w, string(term))
		if err != nil {
			return err
		}
	}
	return nil
}
