package main

import "github.com/koron/csvim/internal/highlight"

type attrsAll struct {
	list highlight.AttrList
}

func (a attrsAll) ApplyGroup(g *highlight.Group) *highlight.Group {
	return g.WithTerm(a.list).WithCTerm(a.list).WithGUI(a.list)
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
}

func (cs colors) ApplyGroup(g *highlight.Group) *highlight.Group {
	return g.WithFg(cs.fg).WithBg(cs.bg)
}

var Palette = []Color{
	Color{Nr: "Black", Name: "grey0"},      // 0:
	Color{Nr: "Black", Name: "grey10"},     // 1:
	Color{Nr: "DarkGrey", Name: "grey25"},  // 2: baseFg
	Color{Nr: "DarkGrey", Name: "grey40"},  // 3:
	Color{Nr: "DarkGrey", Name: "grey50"},  // 4:
	Color{Nr: "LightGrey", Name: "grey60"}, // 5:
	Color{Nr: "LightGrey", Name: "grey70"}, // 6: baseBg
	Color{Nr: "White", Name: "grey80"},     // 7: lightBg
	Color{Nr: "White", Name: "grey90"},     // 8: lightFg
	Color{Nr: "White", Name: "grey100"},    // 9:
}

var (
	none    = attrsAll{list: highlight.None}
	bold    = attrsAll{list: highlight.Bold}
	reverse = attrsAll{list: highlight.Reverse}
)

var (
	normalColors     = colors{fg: Palette[2], bg: Palette[6]}
	nonTextColors    = colors{fg: Palette[8], bg: Palette[5]}
	terminalColors   = colors{fg: Palette[6], bg: Palette[2]}
	foldColumnColors = colors{fg: Palette[2], bg: Palette[4]}

	statusLineColors       = colors{fg: Palette[8], bg: Palette[2]}
	statusLineNCColors     = colors{fg: Palette[6], bg: Palette[2]}
	vertSplitColors        = colors{fg: Palette[2], bg: Palette[2]}
	statusLineTermColors   = colors{fg: Palette[8], bg: Palette[3]}
	statusLineTermNCColors = colors{fg: Palette[6], bg: Palette[3]}

	subCursorColors  = colors{fg: Palette[2], bg: Palette[7]}
	matchParenColors = colors{fg: Palette[2], bg: Palette[8]}
)
