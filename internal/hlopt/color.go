package hlopt

import "github.com/koron/csvim/internal/highlight"

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

type Colors struct {
	Fg highlight.Color
	Bg highlight.Color
	Sp highlight.Color
}

func (cs Colors) ApplyGroup(g *highlight.Group) *highlight.Group {
	return g.WithFg(cs.Fg).WithBg(cs.Bg).WithGUISp(cs.Sp)
}
