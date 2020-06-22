package main

import (
	"log"
	"os"

	"github.com/koron/csvim/internal/colorscheme"
	"github.com/koron/csvim/internal/highlight"
)

func main() {
	colorscheme.WarnDefaultGroups = true
	cs := colorscheme.New("monochromenote").WithBackground(colorscheme.Light)

	var (
		attrNone   = highlight.AttrList{"NONE"}
		attrBold   = highlight.AttrList{"bold"}
		attrRev    = highlight.AttrList{"reverse"}
		attrUC     = highlight.AttrList{"undercurl"}
		attrUL     = highlight.AttrList{"underline"}
		attrBoldUL = highlight.AttrList{"bold", "underline"}

		baseAS   = highlight.AttrSet{Term: attrNone, GUI: attrNone}
		boldAS   = highlight.AttrSet{Term: attrBold, GUI: attrBold}
		revAS    = highlight.AttrSet{Term: attrRev, GUI: attrRev}
		ucAS     = highlight.AttrSet{Term: attrUL, GUI: attrUC}
		ulAS     = highlight.AttrSet{Term: attrUL, GUI: attrUL}
		boldULAS = highlight.AttrSet{Term: attrBoldUL, CTerm: attrBoldUL, GUI: attrBoldUL}

		baseFg = Color{Nr: "Black", Name: "gray25"}
		baseBg = Color{Nr: "DarkGray", Name: "gray70"}
		baseCS = highlight.ColorSet{Fg: baseFg, Bg: baseBg}

		lightFg   = Color{Nr: "LightGray", Name: "gray90"}
		lightBg   = Color{Nr: "LightGray", Name: "gray80"}
		lightCS   = highlight.ColorSet{Fg: lightFg, Bg: baseBg}
		lightArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: lightCS}

		darkFg = Color{Nr: "Black", Name: "gray10"}
		darkBg = Color{Nr: "DarkGray", Name: "gray60"}

		semiBoldCS   = highlight.ColorSet{Fg: darkFg, Bg: baseBg}
		semiBoldArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: semiBoldCS}

		darkAccent1Args = highlight.Arguments{
			AttrSet: baseAS,
			ColorSet: highlight.ColorSet{
				Fg: Color{Nr: "LightGray", Name: "gray90"},
				Bg: Color{Nr: "DarkGray", Name: "gray10"},
			},
		}
		darkAccent2Args = highlight.Arguments{
			AttrSet: baseAS,
			ColorSet: highlight.ColorSet{
				Fg: Color{Nr: "LightGray", Name: "gray80"},
				Bg: Color{Nr: "DarkGray", Name: "gray50"},
			},
		}
		lightBoldArgs = highlight.Arguments{AttrSet: boldAS, ColorSet: lightCS}

		scrollBarC   = Color{Nr: "Black", Name: "gray40"}
		scrollThumbC = Color{Nr: "White", Name: "gray90"}
	)

	cs.Group("Normal").Apply(normalColors)

	cs.Group("NonText").Apply(nonTextColors, bold)
	cs.Group("Terminal").Apply(terminalColors)
	cs.Group("FoldColumn").Apply(foldColumnColors)
	cs.Group("SignColumn").Apply(foldColumnColors)

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

	// TODO:
	extraCursorArgs := highlight.Arguments{
		AttrSet: revAS,
		ColorSet: highlight.ColorSet{
			Fg: Color{Nr: "White", Name: "white"},
			Bg: Color{Nr: "DarkGray", Name: "gray25"},
		},
	}
	cs.Group("lCursor").WithArguments(extraCursorArgs)
	cs.Group("CursorIM").WithArguments(extraCursorArgs)

	cs.Group("LineNr").WithArguments(lightArgs)
	cs.Group("CursorLineNr").WithColorSet(lightCS).WithAttrSet(boldAS)
	cs.Group("Search").WithArguments(darkAccent2Args)
	cs.Link("IncSearch", "Cursor")

	cs.Group("WildMenu").WithArguments(extraCursorArgs)

	cs.Group("Folded").WithArguments(lightArgs)
	cs.Group("Question").WithArguments(lightArgs)
	cs.Group("Title").WithArguments(lightBoldArgs)
	cs.Group("Conceal").WithArguments(lightArgs)
	cs.Group("SpecialKey").WithArguments(lightArgs)

	visualBg := Color{Nr: "LightGray", Name: "gray85"}
	cs.Group("Visual").WithTerm(attrRev).WithCTerm(attrRev).WithGUIBg(visualBg)
	cs.Group("VisualNOS").WithGUIBg(visualBg).WithAttrSet(boldULAS)

	cs.Group("Directory").WithArguments(semiBoldArgs)
	cs.Group("ErrorMsg").WithArguments(darkAccent1Args)
	cs.Group("ModeMsg").WithArguments(semiBoldArgs)
	cs.Group("MoreMsg").WithArguments(semiBoldArgs)
	cs.Group("WarningMsg").WithArguments(lightBoldArgs)

	cs.Group("DiffAdd").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: lightBg},
	})
	cs.Group("DiffDelete").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Bg: lightBg},
	})
	cs.Group("DiffChange").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Bg: darkBg},
	})
	cs.Group("DiffText").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: darkBg},
	})

	cs.Group("TabLineFill").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Bg: darkFg},
	})
	cs.Group("TabLine").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: darkBg, Bg: darkFg},
	})
	cs.Group("TabLineSel").WithArguments(highlight.Arguments{
		AttrSet:  boldAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: darkBg},
	})

	cs.Group("Pmenu").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: darkBg},
	})
	cs.Group("PmenuSel").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Bg: darkBg},
	})
	cs.Group("PmenuSbar").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Bg: scrollBarC},
	})
	cs.Group("PmenuThumb").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Bg: scrollThumbC},
	})

	cs.Group("SpellBad").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Sp: lightFg},
	})
	cs.Group("SpellCap").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Sp: baseFg},
	})
	cs.Group("SpellRare").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Sp: lightFg},
	})
	cs.Group("SpellLocal").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Sp: baseFg},
	})

	// custom groups

	cs.Group("Comment").WithArguments(lightArgs)
	cs.Group("Statement").WithColorSet(baseCS).WithAttrSet(boldAS)
	cs.Group("Type").WithArguments(semiBoldArgs)
	cs.Group("PreProc").WithArguments(semiBoldArgs)
	cs.Group("Identifier").WithArguments(semiBoldArgs)
	cs.Group("Special").WithAttrSet(baseAS).WithColorSet(highlight.ColorSet{Fg: darkFg, Bg: lightBg})
	cs.Group("Constant").WithAttrSet(baseAS).WithColorSet(highlight.ColorSet{Fg: baseFg, Bg: lightBg})
	cs.Group("Error").WithArguments(darkAccent1Args)
	cs.Group("Underlined").WithArguments(highlight.Arguments{AttrSet: ulAS, ColorSet: baseCS})

	cs.Group("Todo").WithArguments(highlight.Arguments{
		AttrSet: baseAS,
		ColorSet: highlight.ColorSet{
			Fg: Color{Nr: "LightGray", Name: "gray90"},
			Bg: Color{Nr: "DarkGray", Name: "gray50"},
		},
	})

	if err := cs.Marshal(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
