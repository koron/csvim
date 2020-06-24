package hlopt

import "github.com/koron/csvim/internal/highlight"

// AttrAll returns Option to apply attributes to all fields of Group:
// Term, CTerm and GUI.
func AttrAll(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		appendAttrs(&g.Term, attrs)
		appendAttrs(&g.CTerm, attrs)
		appendAttrs(&g.GUI, attrs)
	})
}

// Term returns Option to apply attributes to Term field of Group.
func Term(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		appendAttrs(&g.Term, attrs)
	})
}

// CTerm returns Option to apply attributes to CTerm field of Group.
func CTerm(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		appendAttrs(&g.CTerm, attrs)
	})
}

// GUI returns Option to apply attributes to GUI field of Group.
func GUI(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		appendAttrs(&g.GUI, attrs)
	})
}
