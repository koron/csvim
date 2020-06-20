package highlight

import (
	"fmt"
	"io"
)

var Command = "highlight"

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

type Arguments struct {
	AttrSet  AttrSet
	ColorSet ColorSet
}
