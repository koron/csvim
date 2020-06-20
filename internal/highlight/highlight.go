package highlight

import (
	"fmt"
	"io"
)

var Command = "highlight"

type Attr string

const (
	None          Attr = "NONE"
	Bold          Attr = "bold"
	Underline     Attr = "underline"
	Undercurl     Attr = "undercurl"
	Strikethrough Attr = "strikethrough"
	Reverse       Attr = "reverse"
	Inverse       Attr = "inverse"
	Italic        Attr = "italic"
	Standout      Attr = "standout"
	Nocombine     Attr = "nocombine"
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
