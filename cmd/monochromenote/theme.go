package main

import "github.com/koron/csvim/internal/highlight"

type attrsAll struct {
	list highlight.AttrList
}

func (a attrsAll) ApplyGroup(g *highlight.Group) *highlight.Group {
	return g.MergeTerm(a.list).MergeCTerm(a.list).MergeGUI(a.list)
}

type Color struct {
	Nr   highlight.ColorNr
	Name highlight.ColorName
}

func (c Color) TermColor() highlight.ColorNr {
	return c.Nr
}

func (c Color) GUIColor() highlight.ColorName {
	return c.Name
}

type colors struct {
	fg Color
	bg Color
	sp Color
}

func (cs colors) ApplyGroup(g *highlight.Group) *highlight.Group {
	return g.WithFg(cs.fg).WithBg(cs.bg).WithGUISp(cs.sp)
}
