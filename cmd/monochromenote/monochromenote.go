package main

import (
	"log"
	"os"

	"github.com/koron/csvim/internal/colorscheme"
	"github.com/koron/csvim/internal/highlight"
	"github.com/koron/csvim/internal/hlopt"
)

var (
	none      = hlopt.AttrAll(highlight.None)
	bold      = hlopt.AttrAll(highlight.Bold)
	reverse   = hlopt.AttrAll(highlight.Reverse)
	underline = hlopt.AttrAll(highlight.Underline)
	undercurl = hlopt.AttrAll(highlight.Undercurl)
)

var palette = []hlopt.Color{
	{Nr: "Black", Name: "grey0"},      // 0:
	{Nr: "Black", Name: "grey10"},     // 1: darkFg
	{Nr: "DarkGrey", Name: "grey25"},  // 2: baseFg
	{Nr: "DarkGrey", Name: "grey40"},  // 3: scrollBarC
	{Nr: "DarkGrey", Name: "grey50"},  // 4:
	{Nr: "LightGrey", Name: "grey60"}, // 5: darkBg
	{Nr: "LightGrey", Name: "grey70"}, // 6: baseBg
	{Nr: "White", Name: "grey80"},     // 7: lightBg
	{Nr: "White", Name: "grey85"},     // 8:
	{Nr: "White", Name: "grey90"},     // 9: lightFg
	{Nr: "White", Name: "grey100"},    // 10: white
}

var (
	_normal = hlopt.Colors{Fg: palette[2], Bg: palette[6]}
	_light  = hlopt.Colors{Fg: palette[9], Bg: palette[6]}
	_dark   = hlopt.Colors{Fg: palette[1], Bg: palette[6]}

	normalColors     = _normal
	nonTextColors    = hlopt.Colors{Fg: palette[9], Bg: palette[5]}
	terminalColors   = hlopt.Colors{Fg: palette[6], Bg: palette[2]}
	foldColumnColors = hlopt.Colors{Fg: palette[2], Bg: palette[4]}

	lineNrColors = hlopt.Colors{Fg: palette[9], Bg: palette[6]}

	statusLineColors       = hlopt.Colors{Fg: palette[9], Bg: palette[2]}
	statusLineNCColors     = hlopt.Colors{Fg: palette[6], Bg: palette[2]}
	vertSplitColors        = hlopt.Colors{Fg: palette[2], Bg: palette[2]}
	statusLineTermColors   = hlopt.Colors{Fg: palette[9], Bg: palette[3]}
	statusLineTermNCColors = hlopt.Colors{Fg: palette[6], Bg: palette[3]}

	subCursorColors   = hlopt.Colors{Fg: palette[2], Bg: palette[7]}
	matchParenColors  = hlopt.Colors{Fg: palette[2], Bg: palette[9]}
	extraCursorColors = hlopt.Colors{Fg: palette[10], Bg: palette[2]}

	wildMenuColors   = extraCursorColors
	searchColors     = hlopt.Colors{Fg: palette[7], Bg: palette[4]}
	visualColor      = palette[8] // for setting to GUIBg only
	errorMsgColors   = hlopt.Colors{Fg: palette[9], Bg: palette[1]}
	warningMsgColors = hlopt.Colors{Fg: palette[9], Bg: palette[2]}

	diffAddColors    = hlopt.Colors{Fg: palette[2], Bg: palette[7]}
	diffDeleteColors = hlopt.Colors{Fg: palette[9], Bg: palette[7]}
	diffChangeColors = hlopt.Colors{Fg: palette[9], Bg: palette[5]}
	diffTextColors   = hlopt.Colors{Fg: palette[2], Bg: palette[5]}

	tabLineFillColors = hlopt.Colors{Bg: palette[1]}
	tabLineColors     = hlopt.Colors{Fg: palette[5], Bg: palette[1]}
	tabLineSelColors  = hlopt.Colors{Fg: palette[2], Bg: palette[5]}

	pMenuColors      = hlopt.Colors{Fg: palette[2], Bg: palette[5]}
	pMenuSelColors   = hlopt.Colors{Fg: palette[9], Bg: palette[5]}
	pMenuSbarColors  = hlopt.Colors{Bg: palette[3]}
	pMenuThumbColors = hlopt.Colors{Bg: palette[9]}

	spellBadColors   = hlopt.Colors{Fg: palette[2], Sp: palette[9]}
	spellCapColors   = hlopt.Colors{Fg: palette[2], Sp: palette[2]}
	spellRareColors  = hlopt.Colors{Fg: palette[9], Sp: palette[9]}
	spellLocalColors = hlopt.Colors{Fg: palette[9], Sp: palette[2]}

	specialColors  = hlopt.Colors{Fg: palette[1], Bg: palette[7]}
	constantColors = hlopt.Colors{Fg: palette[2], Bg: palette[7]}
	errorColors    = hlopt.Colors{Fg: palette[4], Bg: palette[10]}
	todoColors     = hlopt.Colors{Fg: palette[10], Bg: palette[7]}
)

func main() {
	colorscheme.WarnDefaultGroups = true
	cs := colorscheme.New("monochromenote").WithBackground(colorscheme.Light)

	cs.Group("Normal").Apply(normalColors)

	cs.Group("NonText").Apply(nonTextColors, bold)
	cs.Group("Terminal").Apply(terminalColors)
	cs.Group("FoldColumn").Apply(foldColumnColors)
	cs.Group("SignColumn").Apply(foldColumnColors)

	cs.Group("LineNr").Apply(lineNrColors)
	cs.Group("CursorLineNr").Apply(lineNrColors, bold)

	cs.Group("StatusLine").Apply(statusLineColors, bold)
	cs.Group("StatusLineNC").Apply(statusLineNCColors, none)
	cs.Group("VertSplit").Apply(vertSplitColors, none)
	cs.Group("StatusLineTerm").Apply(statusLineTermColors, bold)
	cs.Group("StatusLineTermNC").Apply(statusLineTermNCColors, none)

	cs.Group("Cursor").Apply(normalColors, reverse)
	cs.Group("CursorColumn").Apply(subCursorColors, none)
	cs.Group("CursorLine").Apply(subCursorColors, none)
	cs.Group("ColorColumn").Apply(subCursorColors, none)
	cs.Group("MatchParen").Apply(matchParenColors)
	cs.Group("lCursor").Apply(extraCursorColors, reverse)
	cs.Group("CursorIM").Apply(extraCursorColors, reverse)

	cs.Group("WildMenu").Apply(wildMenuColors, reverse)

	cs.Group("Search").Apply(searchColors)
	cs.Link("IncSearch", "Cursor")

	cs.Group("Folded").Apply(_light)
	cs.Group("Question").Apply(_light)
	cs.Group("Title").Apply(_light, bold)
	cs.Group("Conceal").Apply(_light)
	cs.Group("SpecialKey").Apply(_light)

	cs.Group("Visual").Apply(
		hlopt.GUIBg(visualColor),
		hlopt.Term(highlight.Reverse),
		hlopt.CTerm(highlight.Reverse),
	)
	cs.Group("VisualNOS").Apply(hlopt.GUIBg(visualColor), bold, underline)

	cs.Group("Directory").Apply(_dark)
	cs.Group("ErrorMsg").Apply(errorMsgColors)
	cs.Group("ModeMsg").Apply(_dark)
	cs.Group("MoreMsg").Apply(_dark)
	cs.Group("WarningMsg").Apply(warningMsgColors, bold)

	cs.Group("DiffAdd").Apply(diffAddColors)
	cs.Group("DiffDelete").Apply(diffDeleteColors)
	cs.Group("DiffChange").Apply(diffChangeColors)
	cs.Group("DiffText").Apply(diffTextColors)

	cs.Group("TabLineFill").Apply(tabLineFillColors)
	cs.Group("TabLine").Apply(tabLineColors)
	cs.Group("TabLineSel").Apply(tabLineSelColors, bold)

	cs.Group("Pmenu").Apply(pMenuColors)
	cs.Group("PmenuSel").Apply(pMenuSelColors)
	cs.Group("PmenuSbar").Apply(pMenuSbarColors)
	cs.Group("PmenuThumb").Apply(pMenuThumbColors)

	cs.Group("SpellBad").Apply(spellBadColors, undercurl)
	cs.Group("SpellCap").Apply(spellCapColors, undercurl)
	cs.Group("SpellRare").Apply(spellRareColors, undercurl)
	cs.Group("SpellLocal").Apply(spellLocalColors, undercurl)

	// custom groups

	cs.Group("Comment").Apply(_light)
	cs.Group("Statement").Apply(_normal, bold)
	cs.Group("Type").Apply(_dark)
	cs.Group("PreProc").Apply(_dark)
	cs.Group("Identifier").Apply(_dark)
	cs.Group("Special").Apply(specialColors)
	cs.Group("Constant").Apply(constantColors)
	cs.Group("Error").Apply(errorColors)
	cs.Group("Underlined").Apply(_normal, underline)
	cs.Group("Todo").Apply(todoColors)

	if err := cs.Marshal(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
