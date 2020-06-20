package highlight

import (
	"errors"
	"fmt"
	"io"
)

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
}

func (g *Group) WithAttrSet(attrSet AttrSet) *Group {
	g.Term = attrSet.Term
	g.CTerm = attrSet.CTerm
	g.GUI = attrSet.GUI
	return g
}

func (g *Group) WithFg(c Color) *Group {
	g.CTermFg = c.Nr
	g.GUIFg = c.Name
	return g
}

func (g *Group) WithBg(c Color) *Group {
	g.CTermBg = c.Nr
	g.GUIBg = c.Name
	return g
}

func (g *Group) WithSp(c Color) *Group {
	g.GUISp = c.Name
	return g
}

func (g *Group) WithColorSet(colorSet ColorSet) *Group {
	return g.WithFg(colorSet.Fg).WithBg(colorSet.Bg).WithSp(colorSet.Sp)
}

func (g *Group) WithArguments(args Arguments) *Group {
	return g.WithAttrSet(args.AttrSet).WithColorSet(args.ColorSet)
}

func (g *Group) Marshal(w io.Writer) error {
	if g.Name == "" {
		return errors.New("highlight with empty Name is not allowed")
	}
	fmt.Fprintf(w, "%s %s", Command, g.Name)

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
	if err := g.GUIBg.writeTo(w, "guisp"); err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w); err != nil {
		return err
	}
	return nil
}
