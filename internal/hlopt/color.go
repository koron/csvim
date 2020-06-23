package hlopt

import "github.com/koron/csvim/internal/highlight"

// CTermFg returns highlight.Option to apply a color to CTermFg field of Group.
func CTermFg(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermFg.Set(c)
	})
}

// CTermBg returns highlight.Option to apply a color to CTermBg field of Group.
func CTermBg(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermBg.Set(c)
	})
}

// GUIFg returns highlight.Option to apply a color to GUIFg field of Group.
func GUIFg(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUIFg.Set(c)
	})
}

// GUIBg returns highlight.Option to apply a color to GUIBg field of Group.
func GUIBg(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUIBg.Set(c)
})
}

// GUISp returns highlight.Option to apply a color to GUISp field of Group.
func GUISp(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUISp.Set(c)
	})
}

// Fg returns highlight.Option to apply a color to fields of Group: CTermFg and
// GUIFg.
func Fg(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermFg.Set(c)
		g.GUIFg.Set(c)
	})
}

// Bg returns highlight.Option to apply a color to fields of Group: CTermBg and
// GUIBg.
func Bg(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.CTermBg.Set(c)
		g.GUIBg.Set(c)
	})
}

// Sp returns highlight.Option to apply a color to field of Group: GUISp.
func Sp(c highlight.Color) highlight.Option {
	if c == nil {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.GUISp.Set(c)
	})
}

// Color is scheme of a color, implements highlight.Color
type Color struct {
	Nr   highlight.ColorNr
	Name highlight.ColorName
}

// TermColor implements highlight.TermColor
func (c Color) TermColor() highlight.ColorNr {
	return c.Nr
}

// GUIColor implements highlight.GUIColor
func (c Color) GUIColor() highlight.ColorName {
	return c.Name
}

// Colors is a set of Colors for Fg, Bg, Sp.
type Colors struct {
	Fg highlight.Color
	Bg highlight.Color
	Sp highlight.Color
}

// Apply implements highlight.Option for Colors
func (cs Colors) Apply(g *highlight.Group) {
	g.Apply(Fg(cs.Fg), Bg(cs.Bg), Sp(cs.Sp))
}
