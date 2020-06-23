package main

import (
	"log"
	"os"

	"github.com/koron/csvim/internal/colorscheme"
	"github.com/koron/csvim/internal/highlight"
)

var (
	none      = attrsAll{list: highlight.None}
	bold      = attrsAll{list: highlight.Bold}
	reverse   = attrsAll{list: highlight.Reverse}
	underline = attrsAll{list: highlight.Underline}
	undercurl = attrsAll{list: highlight.Undercurl}
)

var Palette = []Color{
	Color{Nr: "Black", Name: "grey0"},      // 0:
	Color{Nr: "Black", Name: "grey10"},     // 1: darkFg
	Color{Nr: "DarkGrey", Name: "grey25"},  // 2: baseFg
	Color{Nr: "DarkGrey", Name: "grey40"},  // 3: scrollBarC
	Color{Nr: "DarkGrey", Name: "grey50"},  // 4:
	Color{Nr: "LightGrey", Name: "grey60"}, // 5: darkBg
	Color{Nr: "LightGrey", Name: "grey70"}, // 6: baseBg
	Color{Nr: "White", Name: "grey80"},     // 7: lightBg
	Color{Nr: "White", Name: "grey85"},     // 8:
	Color{Nr: "White", Name: "grey90"},     // 9: lightFg
	Color{Nr: "White", Name: "grey100"},    // 10: white
}

var (
	_normal = colors{fg: Palette[2], bg: Palette[6]}
	_light  = colors{fg: Palette[9], bg: Palette[6]}
	_dark   = colors{fg: Palette[1], bg: Palette[6]}

	normalColors     = _normal
	nonTextColors    = colors{fg: Palette[9], bg: Palette[5]}
	terminalColors   = colors{fg: Palette[6], bg: Palette[2]}
	foldColumnColors = colors{fg: Palette[2], bg: Palette[4]}

	lineNrColors = colors{fg: Palette[9], bg: Palette[6]}

	statusLineColors       = colors{fg: Palette[9], bg: Palette[2]}
	statusLineNCColors     = colors{fg: Palette[6], bg: Palette[2]}
	vertSplitColors        = colors{fg: Palette[2], bg: Palette[2]}
	statusLineTermColors   = colors{fg: Palette[9], bg: Palette[3]}
	statusLineTermNCColors = colors{fg: Palette[6], bg: Palette[3]}

	subCursorColors   = colors{fg: Palette[2], bg: Palette[7]}
	matchParenColors  = colors{fg: Palette[2], bg: Palette[9]}
	extraCursorColors = colors{fg: Palette[10], bg: Palette[2]}

	wildMenuColors   = extraCursorColors
	searchColors     = colors{fg: Palette[7], bg: Palette[4]}
	visualColor      = Palette[8] // for setting to GUIBg only
	errorMsgColors   = colors{fg: Palette[9], bg: Palette[1]}
	warningMsgColors = colors{fg: Palette[9], bg: Palette[2]}

	diffAddColors    = colors{fg: Palette[2], bg: Palette[7]}
	diffDeleteColors = colors{fg: Palette[9], bg: Palette[7]}
	diffChangeColors = colors{fg: Palette[9], bg: Palette[5]}
	diffTextColors   = colors{fg: Palette[2], bg: Palette[5]}

	tabLineFillColors = colors{bg: Palette[1]}
	tabLineColors     = colors{fg: Palette[5], bg: Palette[1]}
	tabLineSelColors  = colors{fg: Palette[2], bg: Palette[5]}

	pMenuColors      = colors{fg: Palette[2], bg: Palette[5]}
	pMenuSelColors   = colors{fg: Palette[9], bg: Palette[5]}
	pMenuSbarColors  = colors{bg: Palette[3]}
	pMenuThumbColors = colors{bg: Palette[9]}

	spellBadColors   = colors{fg: Palette[2], sp: Palette[9]}
	spellCapColors   = colors{fg: Palette[2], sp: Palette[2]}
	spellRareColors  = colors{fg: Palette[9], sp: Palette[9]}
	spellLocalColors = colors{fg: Palette[9], sp: Palette[2]}

	specialColors  = colors{fg: Palette[1], sp: Palette[7]}
	constantColors = colors{fg: Palette[2], sp: Palette[7]}
	todoColors     = colors{fg: Palette[9], bg: Palette[4]}
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
