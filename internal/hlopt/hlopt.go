/*
Package hlopt provides various implementations of highlight.Option.
*/
package hlopt

import "github.com/koron/csvim/internal/highlight"

type nopOption struct{}

func (*nopOption) Apply(g *highlight.Group) {}

var nop *nopOption
