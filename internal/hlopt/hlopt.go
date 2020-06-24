/*
Package hlopt provides various implementations of highlight.Option.
*/
package hlopt

import "github.com/koron/csvim/internal/highlight"

type nopOption struct{}

func (*nopOption) Apply(g *highlight.Group) {}

var nop *nopOption

// Font returns Option to apply a font to Font field of Group.
func Font(font string) highlight.Option {
	if font == "" {
		return nop
	}
	return highlight.OptionFunc(func(g *highlight.Group) {
		g.Font = font
	})
}
