package highlight

import (
	"fmt"
	"io"
)

// Attr is a value for "term", "cterm" and "gui" arguments.
type Attr string

const (
	// None is no attributes used (used to reset it).
	None Attr = "NONE"
	// Bold is "bold".
	Bold Attr = "bold"
	// Underline is for underline decolation.
	Underline Attr = "underline"
	// Undercurl is for undercurl decolation, not always available.
	Undercurl Attr = "undercurl"
	// Strikethrough is for strike through decolation, not always available.
	Strikethrough Attr = "strikethrough"
	// Reverse is an effect to swap Fg and Bg.
	Reverse Attr = "reverse"
	// Inverse is same as Reverse.
	Inverse Attr = "inverse"
	// Italic means using italic font.
	Italic Attr = "italic"
	// Standout means ... just stand out.
	Standout Attr = "standout"
	// Nocombine overrides attributes instead of combining them.
	Nocombine Attr = "nocombine"
)

// AttrList is set of Attr for "term", "cterm" and "gui" arguments.
type AttrList []Attr

// Append adds AttrList.
func (attrs *AttrList) Append(src AttrList) {
	if len(src) == 0 {
		return
	}
	*attrs = append(*attrs, src...)
}

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
