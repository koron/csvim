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

		baseFg   = highlight.Color{Nr: "Black", Name: "gray25"}
		baseBg   = highlight.Color{Nr: "DarkGray", Name: "gray70"}
		baseCS   = highlight.ColorSet{Fg: baseFg, Bg: baseBg}
		baseArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: baseCS}

		lightFg   = highlight.Color{Nr: "LightGray", Name: "gray90"}
		lightBg   = highlight.Color{Nr: "LightGray", Name: "gray80"}
		lightCS   = highlight.ColorSet{Fg: lightFg, Bg: baseBg}
		lightArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: lightCS}

		darkFg = highlight.Color{Nr: "Black", Name: "gray10"}
		darkBg = highlight.Color{Nr: "DarkGray", Name: "gray60"}
		darkCS = highlight.ColorSet{Fg: baseFg, Bg: highlight.Color{Nr: "DarkGray", Name: "gray50"}}

		foldColArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: darkCS}

		termBgC = highlight.Color{Nr: "Black", Name: "gray40"}

		statusCS       = highlight.ColorSet{Fg: lightFg, Bg: baseFg}
		statusNCCS     = highlight.ColorSet{Fg: baseBg, Bg: baseFg}
		statusTermCS   = highlight.ColorSet{Fg: lightFg, Bg: termBgC}
		statusTermNCCS = highlight.ColorSet{Fg: baseBg, Bg: termBgC}
		vertSplitCS    = highlight.ColorSet{Fg: darkFg, Bg: baseFg}

		subCurCS   = highlight.ColorSet{Fg: baseFg, Bg: lightBg}
		subCurArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: subCurCS}

		semiBoldCS   = highlight.ColorSet{Fg: darkFg, Bg: baseBg}
		semiBoldArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: semiBoldCS}

		darkAccent1Args = highlight.Arguments{
			AttrSet: baseAS,
			ColorSet: highlight.ColorSet{
				Fg: highlight.Color{Nr: "LightGray", Name: "gray90"},
				Bg: highlight.Color{Nr: "DarkGray", Name: "gray10"},
			},
		}
		darkAccent2Args = highlight.Arguments{
			AttrSet: baseAS,
			ColorSet: highlight.ColorSet{
				Fg: highlight.Color{Nr: "LightGray", Name: "gray80"},
				Bg: highlight.Color{Nr: "DarkGray", Name: "gray50"},
			},
		}
		lightBoldArgs = highlight.Arguments{AttrSet: boldAS, ColorSet: lightCS}

		scrollBarC   = highlight.Color{Nr: "Black", Name: "gray40"}
		scrollThumbC = highlight.Color{Nr: "White", Name: "gray90"}
	)

	cs.AddGroup(highlight.NewGroup("Normal").WithArguments(baseArgs))
	cs.AddGroup(highlight.NewGroup("NonText").WithArguments(highlight.Arguments{
		AttrSet:  boldAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Bg: darkBg},
	}))
	cs.AddGroup(highlight.NewGroup("FoldColumn").WithArguments(foldColArgs))
	cs.AddGroup(highlight.NewGroup("SignColumn").WithArguments(foldColArgs))

	cs.AddGroup(highlight.NewGroup("StatusLine").WithColorSet(statusCS).WithAttrSet(boldAS))
	cs.AddGroup(highlight.NewGroup("StatusLineNC").WithColorSet(statusNCCS).WithAttrSet(baseAS))
	cs.AddGroup(highlight.NewGroup("VertSplit").WithColorSet(vertSplitCS).WithAttrSet(baseAS))
	cs.AddGroup(highlight.NewGroup("StatusLineTerm").WithAttrSet(boldAS).WithColorSet(statusTermCS))
	cs.AddGroup(highlight.NewGroup("StatusLineTermNC").WithAttrSet(boldAS).WithColorSet(statusTermNCCS))

	cs.AddGroup(highlight.NewGroup("Cursor").WithColorSet(baseCS).WithAttrSet(revAS))
	cs.AddGroup(highlight.NewGroup("CursorColumn").WithArguments(subCurArgs))
	cs.AddGroup(highlight.NewGroup("CursorLine").WithArguments(subCurArgs))
	cs.AddGroup(highlight.NewGroup("ColorColumn").WithArguments(subCurArgs))
	cs.AddGroup(highlight.NewGroup("MatchParen").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: lightFg},
	}))

	extraCursorArgs := highlight.Arguments{
		AttrSet: revAS,
		ColorSet: highlight.ColorSet{
			Fg: highlight.Color{Nr: "White", Name: "white"},
			Bg: highlight.Color{Nr: "DarkGray", Name: "gray25"},
		},
	}
	cs.AddGroup(highlight.NewGroup("lCursor").WithArguments(extraCursorArgs))
	cs.AddGroup(highlight.NewGroup("CursorIM").WithArguments(extraCursorArgs))

	cs.AddGroup(highlight.NewGroup("LineNr").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("CursorLineNr").WithColorSet(lightCS).WithAttrSet(boldAS))
	cs.AddGroup(highlight.NewGroup("Search").WithArguments(darkAccent2Args))
	cs.AddLink(highlight.NewLink("IncSearch", "Cursor"))

	cs.AddGroup(highlight.NewGroup("WildMenu").WithArguments(extraCursorArgs))

	cs.AddGroup(highlight.NewGroup("Folded").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("Question").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("Title").WithArguments(lightBoldArgs))
	cs.AddGroup(highlight.NewGroup("Conceal").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("SpecialKey").WithArguments(lightArgs))

	visualBg := highlight.Color{Nr: "LightGray", Name: "gray85"}
	cs.AddGroup(&highlight.Group{Name: "Visual", Term: attrRev, CTerm: attrRev, GUIBg: visualBg.Name})
	cs.AddGroup((&highlight.Group{Name: "VisualNOS", GUIBg: visualBg.Name}).WithAttrSet(boldULAS))

	cs.AddGroup(highlight.NewGroup("Directory").WithArguments(semiBoldArgs))
	cs.AddGroup(highlight.NewGroup("ErrorMsg").WithArguments(darkAccent1Args))
	cs.AddGroup(highlight.NewGroup("ModeMsg").WithArguments(semiBoldArgs))
	cs.AddGroup(highlight.NewGroup("MoreMsg").WithArguments(semiBoldArgs))
	cs.AddGroup(highlight.NewGroup("WarningMsg").WithArguments(lightBoldArgs))

	cs.AddGroup(highlight.NewGroup("DiffAdd").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: lightBg},
	}))
	cs.AddGroup(highlight.NewGroup("DiffDelete").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Bg: lightBg},
	}))
	cs.AddGroup(highlight.NewGroup("DiffChange").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Bg: darkBg},
	}))
	cs.AddGroup(highlight.NewGroup("DiffText").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: darkBg},
	}))

	cs.AddGroup(highlight.NewGroup("TablineFill").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Bg: darkFg},
	}))
	cs.AddGroup(highlight.NewGroup("Tabline").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: darkBg, Bg: darkFg},
	}))
	cs.AddGroup(highlight.NewGroup("TablineSel").WithArguments(highlight.Arguments{
		AttrSet:  boldAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: darkBg},
	}))

	cs.AddGroup(highlight.NewGroup("Pmenu").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: baseFg, Bg: darkBg},
	}))
	cs.AddGroup(highlight.NewGroup("PmenuSel").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Bg: darkBg},
	}))
	cs.AddGroup(highlight.NewGroup("PmenuSbar").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Bg: scrollBarC},
	}))
	cs.AddGroup(highlight.NewGroup("PmenuThumb").WithArguments(highlight.Arguments{
		AttrSet:  baseAS,
		ColorSet: highlight.ColorSet{Bg: scrollThumbC},
	}))

	cs.AddGroup(highlight.NewGroup("SpellBad").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Sp: lightFg},
	}))
	cs.AddGroup(highlight.NewGroup("SpellCap").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Sp: baseFg},
	}))
	cs.AddGroup(highlight.NewGroup("SpellRare").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Sp: lightFg},
	}))
	cs.AddGroup(highlight.NewGroup("SpellLocal").WithArguments(highlight.Arguments{
		AttrSet:  ucAS,
		ColorSet: highlight.ColorSet{Fg: lightFg, Sp: baseFg},
	}))

	// custom groups

	cs.AddGroup(highlight.NewGroup("Comment").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("Statement").WithColorSet(baseCS).WithAttrSet(boldAS))
	cs.AddGroup(highlight.NewGroup("Type").WithArguments(semiBoldArgs))
	cs.AddGroup(highlight.NewGroup("PreProc").WithArguments(semiBoldArgs))
	cs.AddGroup(highlight.NewGroup("Identifier").WithArguments(semiBoldArgs))
	cs.AddGroup(highlight.NewGroup("Special").WithAttrSet(baseAS).WithColorSet(highlight.ColorSet{Fg: darkFg, Bg: lightBg}))
	cs.AddGroup(highlight.NewGroup("Constant").WithAttrSet(baseAS).WithColorSet(highlight.ColorSet{Fg: baseFg, Bg: lightBg}))
	cs.AddGroup(highlight.NewGroup("Error").WithArguments(darkAccent1Args))
	cs.AddGroup(highlight.NewGroup("Underlined").WithArguments(highlight.Arguments{AttrSet: ulAS, ColorSet: baseCS}))

	cs.AddGroup(highlight.NewGroup("Todo").WithArguments(highlight.Arguments{
		AttrSet: baseAS,
		ColorSet: highlight.ColorSet{
			Fg: highlight.Color{Nr: "LightGray", Name: "gray90"},
			Bg: highlight.Color{Nr: "DarkGray", Name: "gray50"},
		},
	}))

	if err := cs.Marshal(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
