package highlight

import (
	"errors"
	"fmt"
	"io"
)

// Group represents "highlight [default] {group-name} ..." command.
type Group struct {
	Name string

	Term  AttrList
	Start TermList
	Stop  TermList

	CTerm   AttrList
	CTermFg ColorNr
	CTermBg ColorNr

	GUI   AttrList
	Font  string
	GUIFg ColorName
	GUIBg ColorName
	GUISp ColorName

	Default bool
}

// NewGroup creates a Group with name.
func NewGroup(name string) *Group {
	return &Group{Name: name}
}

// Marshal outputs "highlight [default] {group-name} ..." command.
func (g *Group) Marshal(w io.Writer) error {
	if g.Name == "" {
		return errors.New("highlight with empty Name is not allowed")
	}
	var defSP string
	if g.Default {
		defSP = "default "
	}
	fmt.Fprintf(w, "%[1]s %[3]s%[2]s", Command, g.Name, defSP)

	if err := g.Term.writeTo(w, "term"); err != nil {
		return err
	}
	if err := g.Start.writeTo(w, "start"); err != nil {
		return err
	}
	if err := g.Stop.writeTo(w, "stop"); err != nil {
		return err
	}

	if err := g.CTerm.writeTo(w, "cterm"); err != nil {
		return err
	}
	if err := g.CTermFg.writeTo(w, "ctermfg"); err != nil {
		return err
	}
	if err := g.CTermBg.writeTo(w, "ctermbg"); err != nil {
		return err
	}

	if err := g.GUI.writeTo(w, "gui"); err != nil {
		return err
	}
	if g.Font != "" {
		if _, err := fmt.Fprintf(w, " font=%s", g.Font); err != nil {
			return err
		}
	}
	if err := g.GUIFg.writeTo(w, "guifg"); err != nil {
		return err
	}
	if err := g.GUIBg.writeTo(w, "guibg"); err != nil {
		return err
	}
	if err := g.GUISp.writeTo(w, "guisp"); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	return nil
}
