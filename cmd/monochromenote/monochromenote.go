package main

import (
	"log"
	"os"

	"github.com/koron/csvim/internal/colorscheme"
	"github.com/koron/csvim/internal/highlight"
	"github.com/koron/csvim/internal/hlopt"
)

var (
	none      = highlight.None
	bold      = highlight.Bold
	reverse   = highlight.Reverse
	underline = highlight.Underline
	undercurl = highlight.Undercurl
)

var Palette = []hlopt.Color{
	hlopt.Color{Nr: "Black", Name: "grey0"},      // 0:
	hlopt.Color{Nr: "Black", Name: "grey10"},     // 1: darkFg
	hlopt.Color{Nr: "DarkGrey", Name: "grey25"},  // 2: baseFg
	hlopt.Color{Nr: "DarkGrey", Name: "grey40"},  // 3: scrollBarC
	hlopt.Color{Nr: "DarkGrey", Name: "grey50"},  // 4:
	hlopt.Color{Nr: "LightGrey", Name: "grey60"}, // 5: darkBg
	hlopt.Color{Nr: "LightGrey", Name: "grey70"}, // 6: baseBg
	hlopt.Color{Nr: "White", Name: "grey80"},     // 7: lightBg
	hlopt.Color{Nr: "White", Name: "grey85"},     // 8:
	hlopt.Color{Nr: "White", Name: "grey90"},     // 9: lightFg
	hlopt.Color{Nr: "White", Name: "grey100"},    // 10: white
}

var (
	_normal = hlopt.Colors{Fg: Palette[2], Bg: Palette[6]}
	_light  = hlopt.Colors{Fg: Palette[9], Bg: Palette[6]}
	_dark   = hlopt.Colors{Fg: Palette[1], Bg: Palette[6]}

	normalColors     = _normal
	nonTextColors    = hlopt.Colors{Fg: Palette[9], Bg: Palette[5]}
	terminalColors   = hlopt.Colors{Fg: Palette[6], Bg: Palette[2]}
	foldColumnColors = hlopt.Colors{Fg: Palette[2], Bg: Palette[4]}

	lineNrColors = hlopt.Colors{Fg: Palette[9], Bg: Palette[6]}

	statusLineColors       = hlopt.Colors{Fg: Palette[9], Bg: Palette[2]}
	statusLineNCColors     = hlopt.Colors{Fg: Palette[6], Bg: Palette[2]}
	vertSplitColors        = hlopt.Colors{Fg: Palette[2], Bg: Palette[2]}
	statusLineTermColors   = hlopt.Colors{Fg: Palette[9], Bg: Palette[3]}
	statusLineTermNCColors = hlopt.Colors{Fg: Palette[6], Bg: Palette[3]}

	subCursorColors   = hlopt.Colors{Fg: Palette[2], Bg: Palette[7]}
	matchParenColors  = hlopt.Colors{Fg: Palette[2], Bg: Palette[9]}
	extraCursorColors = hlopt.Colors{Fg: Palette[10], Bg: Palette[2]}

	wildMenuColors   = extraCursorColors
	searchColors     = hlopt.Colors{Fg: Palette[7], Bg: Palette[4]}
	visualColor      = Palette[8] // for setting to GUIBg only
	errorMsgColors   = hlopt.Colors{Fg: Palette[9], Bg: Palette[1]}
	warningMsgColors = hlopt.Colors{Fg: Palette[9], Bg: Palette[2]}

	diffAddColors    = hlopt.Colors{Fg: Palette[2], Bg: Palette[7]}
	diffDeleteColors = hlopt.Colors{Fg: Palette[9], Bg: Palette[7]}
	diffChangeColors = hlopt.Colors{Fg: Palette[9], Bg: Palette[5]}
	diffTextColors   = hlopt.Colors{Fg: Palette[2], Bg: Palette[5]}

	tabLineFillColors = hlopt.Colors{Bg: Palette[1]}
	tabLineColors     = hlopt.Colors{Fg: Palette[5], Bg: Palette[1]}
	tabLineSelColors  = hlopt.Colors{Fg: Palette[2], Bg: Palette[5]}

	pMenuColors      = hlopt.Colors{Fg: Palette[2], Bg: Palette[5]}
	pMenuSelColors   = hlopt.Colors{Fg: Palette[9], Bg: Palette[5]}
	pMenuSbarColors  = hlopt.Colors{Bg: Palette[3]}
	pMenuThumbColors = hlopt.Colors{Bg: Palette[9]}

	spellBadColors   = hlopt.Colors{Fg: Palette[2], Sp: Palette[9]}
	spellCapColors   = hlopt.Colors{Fg: Palette[2], Sp: Palette[2]}
	spellRareColors  = hlopt.Colors{Fg: Palette[9], Sp: Palette[9]}
	spellLocalColors = hlopt.Colors{Fg: Palette[9], Sp: Palette[2]}

	specialColors  = hlopt.Colors{Fg: Palette[1], Sp: Palette[7]}
	constantColors = hlopt.Colors{Fg: Palette[2], Sp: Palette[7]}
	todoColors     = hlopt.Colors{Fg: Palette[9], Bg: Palette[4]}
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

	cs.Group("Visual").WithGUIBg(visualColor).
		WithTerm(highlight.Reverse).WithCTerm(highlight.Reverse)
	cs.Group("VisualNOS").WithGUIBg(visualColor).Apply(bold, underline)

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
	cs.Group("Error").Apply(errorMsgColors)
	cs.Group("Underlined").Apply(_normal, underline)
	cs.Group("Todo").Apply(todoColors)

	if err := cs.Marshal(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
