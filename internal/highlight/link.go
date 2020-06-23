package highlight

import (
	"errors"
	"fmt"
	"io"
)

// Link represents "highlight link" entry.
type Link struct {
	From string
	To   string

	Default bool
}

// NewLink creates a link.
func NewLink(from, to string) *Link {
	return &Link{From: from, To: to}
}

// Marshal outputs "highlight link ..." command to io.Writer.
func (ln *Link) Marshal(w io.Writer) error {
	if ln.From == "" || ln.To == "" {
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
