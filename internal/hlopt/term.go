package hlopt

import "github.com/koron/csvim/internal/highlight"

// Start returns highlight.Option to apply terms to Start field of Group.
func Start(terms ...string) highlight.Option {
	if len(terms) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.Start.Append(terms)
	})
}

// Stop returns highlight.Option to apply terms to Stop field of Group.
func Stop(terms ...string) highlight.Option {
	if len(terms) == 0 {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.Stop.Append(terms)
	})
}
