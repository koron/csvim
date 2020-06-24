package hlopt

import "github.com/koron/csvim/internal/highlight"

// Start returns Option to apply terms to Start field of Group.
func Start(terms ...string) highlight.Option {
	if len(terms) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		appendTerms(&g.Start, terms)
	})
}

// Stop returns Option to apply terms to Stop field of Group.
func Stop(terms ...string) highlight.Option {
	if len(terms) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		appendTerms(&g.Stop, terms)
	})
}
