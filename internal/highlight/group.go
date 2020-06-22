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

	Default bool
}

func NewGroup(name string) *Group {
	return &Group{Name: name}
}

func (g *Group) WithTerm(src AttrList) *Group {
	g.Term = src
	return g
}

func (g *Group) WithCTerm(src AttrList) *Group {
	g.CTerm = src
	return g
}

func (g *Group) WithGUI(src AttrList) *Group {
	g.GUI = src
	return g
}

func (g *Group) WithStart(src TermList) *Group {
	g.Start = src
	return g
}

func (g *Group) WithStop(src TermList) *Group {
	g.Stop = src
	return g
}

func (g *Group) WithCTermFg(c TermColor) *Group {
	if c == nil {
		g.CTermFg = ""
		return g
	}
	g.CTermFg = c.TermColor()
	return g
}

func (g *Group) WithCTermBg(c TermColor) *Group {
	if c == nil {
		g.CTermBg = ""
		return g
	}
	g.CTermBg = c.TermColor()
	return g
}

func (g *Group) WithGUIFg(c GUIColor) *Group {
	if c == nil {
		g.GUIFg = ""
		return g
	}
	g.GUIFg = c.GUIColor()
	return g
}

func (g *Group) WithGUIBg(c GUIColor) *Group {
	if c == nil {
		g.GUIBg = ""
		return g
	}
	g.GUIBg = c.GUIColor()
	return g
}

func (g *Group) WithGUISp(c GUIColor) *Group {
	if c == nil {
		g.GUISp = ""
		return g
	}
	g.GUISp = c.GUIColor()
	return g
}

func (g *Group) WithFg(c Color) *Group {
	return g.WithCTermFg(c).WithGUIFg(c)
}

func (g *Group) WithBg(c Color) *Group {
	return g.WithCTermBg(c).WithGUIBg(c)
}

// Merge merges/includes "groups" into "g" and returns modified "g".
func (g *Group) Merge(groups ...*Group) *Group {
	for _, src := range groups {
		g.Term.merge(src.Term)
		g.Start.merge(src.Start)
		g.Stop.merge(src.Stop)
		g.CTerm.merge(src.CTerm)
		g.CTermFg.merge(src.CTermFg)
		g.CTermBg.merge(src.CTermBg)
		g.GUI.merge(src.GUI)
		if src.Font != "" {
			g.Font = src.Font
		}
		g.GUIFg.merge(src.GUIFg)
		g.GUIBg.merge(src.GUIBg)
		g.GUISp.merge(src.GUISp)
		if src.Default {
			g.Default = true
		}
	}
	return g
}

func (g *Group) MergeTerm(src AttrList) *Group {
	g.Term.merge(src)
	return g
}

func (g *Group) MergeCTerm(src AttrList) *Group {
	g.CTerm.merge(src)
	return g
}

func (g *Group) MergeGUI(src AttrList) *Group {
	g.GUI.merge(src)
	return g
}

func (g *Group) MergeStart(src TermList) *Group {
	g.Start.merge(src)
	return g
}

func (g *Group) MergeStop(src TermList) *Group {
	g.Stop.merge(src)
	return g
}

type GroupOption interface {
	ApplyGroup(*Group) *Group
}

type GroupOptionFunc func(*Group) *Group

var _ GroupOption = GroupOptionFunc(nil)

func (fn GroupOptionFunc) ApplyGroup(g *Group) *Group {
	return fn(g)
}

func (g *Group) Apply(opts ...GroupOption) *Group {
	for _, o := range opts {
		if o == nil {
			continue
		}
		g = o.ApplyGroup(g)
	}
	return g
}

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
