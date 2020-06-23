package highlight

import (
	"fmt"
	"io"
)

type Attr string

var (
	None          = AttrList{"NONE"}
	Bold          = AttrList{"bold"}
	Underline     = AttrList{"underline"}
	Undercurl     = AttrList{"undercurl"}
	Strikethrough = AttrList{"strikethrough"}
	Reverse       = AttrList{"reverse"}
	Inverse       = AttrList{"inverse"}
	Italic        = AttrList{"italic"}
	Standout      = AttrList{"standout"}
	Nocombine     = AttrList{"nocombine"}
)

type AttrList []Attr

func (attrs AttrList) ApplyGroup(g *Group) *Group {
	return g.MergeTerm(attrs).MergeCTerm(attrs).MergeGUI(attrs)
}

func (attrs *AttrList) merge(src AttrList) {
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
