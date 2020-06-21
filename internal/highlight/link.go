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

func NewLink(from, to string) *Link {
	return &Link{From: from, To: to}
}

func (ln *Link) Marshal(w io.Writer) error {
	if ln.From == "" || len(ln.To) == 0 {
		return errors.New("link with empty From or To is not allowed")
	}
	var defSP string
	if ln.Default {
		defSP = "default "
	}
	_, err := fmt.Fprintf(w, "%[1]s %[4]slink %[2]s %[3]s\n", Command, ln.From, ln.To, defSP)
	if err != nil {
		return err
	}
	return nil
}
