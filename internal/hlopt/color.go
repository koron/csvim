package hlopt

import "github.com/koron/csvim/internal/highlight"

// Coloring defines an interface to unify color number and name.
type Coloring interface {
	ColorNr() highlight.ColorNr
	ColorName() highlight.ColorName
}

// CTermFg returns highlight.Option to apply a color to CTermFg field of Group.
func CTermFg(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermFg = c.ColorNr()
	})
}

// CTermBg returns highlight.Option to apply a color to CTermBg field of Group.
func CTermBg(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermBg = c.ColorNr()
	})
}

// GUIFg returns highlight.Option to apply a color to GUIFg field of Group.
func GUIFg(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUIFg = c.ColorName()
	})
}

// GUIBg returns highlight.Option to apply a color to GUIBg field of Group.
func GUIBg(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUIBg = c.ColorName()
	})
}

// GUISp returns highlight.Option to apply a color to GUISp field of Group.
func GUISp(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUISp = c.ColorName()
	})
}

// Fg returns highlight.Option to apply a color to fields of Group: CTermFg and
// GUIFg.
func Fg(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermFg = c.ColorNr()
		g.GUIFg = c.ColorName()
	})
}

// Bg returns highlight.Option to apply a color to fields of Group: CTermBg and
// GUIBg.
func Bg(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermBg = c.ColorNr()
		g.GUIBg = c.ColorName()
	})
}

// Sp returns highlight.Option to apply a color to field of Group: GUISp.
func Sp(c Coloring) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUISp = c.ColorName()
	})
}

// Color is scheme of a color, implements Coloring
type Color struct {
	Nr   highlight.ColorNr
	Name highlight.ColorName
}

// ColorNr implements Coloring interface.
func (c Color) ColorNr() highlight.ColorNr {
	return c.Nr
}

// ColorName implements Coloring interface.
func (c Color) ColorName() highlight.ColorName {
	return c.Name
}

// Colors is a set of Colors for Fg, Bg, Sp.
type Colors struct {
	Fg Coloring
	Bg Coloring
	Sp Coloring
}

// Apply implements highlight.Option for Colors
func (cs Colors) Apply(g *highlight.Group) {
	g.Apply(Fg(cs.Fg), Bg(cs.Bg), Sp(cs.Sp))
}
