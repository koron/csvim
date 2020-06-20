package highlight

import (
	"errors"
	"fmt"
	"io"
)

type Link struct {
	From string
	To   string

	Default bool
}

func (ln *Link) Marshal(w io.Writer) error {
	if ln.From == "" || len(ln.To) == 0 {
		return errors.New("link with empty From or To is not allowed")
	}
	_, err := fmt.Fprintf(w, "%s link %s %s\n", Command, ln.From, ln.To)
	if err != nil {
		return err
	}
	return nil
}
