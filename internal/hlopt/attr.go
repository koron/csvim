package hlopt

import "github.com/koron/csvim/internal/highlight"

// AttrAll returns highlight.Option to apply attributes to all fields of Group:
// Term, CTerm and GUI.
func AttrAll(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group){
		g.Term.Append(attrs)
		g.CTerm.Append(attrs)
		g.GUI.Append(attrs)
	})
}

// Term returns highlight.Option to apply attributes to Term field of Group.
func Term(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group){
		g.Term.Append(attrs)
	})
}

// CTerm returns highlight.Option to apply attributes to CTerm field of Group.
func CTerm(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group){
		g.CTerm.Append(attrs)
	})
}

// GUI returns highlight.Option to apply attributes to GUI field of Group.
func GUI(attrs ...highlight.Attr) highlight.Option {
	if len(attrs) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group){
		g.CTerm.Append(attrs)
	})
}
