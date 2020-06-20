package main

import (
	"log"
	"os"

	"github.com/koron/csvim/internal/colorscheme"
	"github.com/koron/csvim/internal/highlight"
)

func main() {
	colorscheme.WarnDefaultGroups = true
	cs := colorscheme.New("blackandwhite").WithBackground(colorscheme.Dark)

	var (
		attrNone = highlight.AttrList{highlight.None}
		attrBold = highlight.AttrList{highlight.Bold}
		attrRev  = highlight.AttrList{highlight.Reverse}

		baseAS = highlight.AttrSet{Term: attrNone, GUI: attrNone}
		boldAS = highlight.AttrSet{Term: attrBold, GUI: attrBold}
		revAS  = highlight.AttrSet{Term: attrRev, GUI: attrRev}

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
		darkCS = highlight.ColorSet{Fg: baseFg, Bg: darkBg}

		foldColArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: darkCS}

		statusCS    = highlight.ColorSet{Fg: lightFg, Bg: baseFg}
		statusNCCS  = highlight.ColorSet{Fg: baseBg, Bg: baseFg}
		vertSplitCS = highlight.ColorSet{Fg: darkFg, Bg: baseFg}

		subCurCS   = highlight.ColorSet{Fg: baseFg, Bg: lightBg}
		subCurArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: subCurCS}

		subBoldCS   = highlight.ColorSet{Fg: darkFg, Bg: baseBg}
		subBoldArgs = highlight.Arguments{AttrSet: baseAS, ColorSet: subBoldCS}
	)

	cs.AddGroup(highlight.NewGroup("Normal").WithArguments(baseArgs))
	cs.AddGroup(highlight.NewGroup("NonText").WithColorSet(lightCS).WithAttrSet(boldAS))
	cs.AddGroup(highlight.NewGroup("FoldColumn").WithArguments(foldColArgs))
	cs.AddGroup(highlight.NewGroup("SignColumn").WithArguments(foldColArgs))

	cs.AddGroup(highlight.NewGroup("StatusLine").WithColorSet(statusCS).WithAttrSet(boldAS))
	cs.AddGroup(highlight.NewGroup("StatusLineNC").WithColorSet(statusNCCS).WithAttrSet(baseAS))
	cs.AddGroup(highlight.NewGroup("VertSplit").WithColorSet(vertSplitCS).WithAttrSet(baseAS))

	cs.AddGroup(highlight.NewGroup("Cursor").WithColorSet(baseCS).WithAttrSet(revAS))
	cs.AddGroup(highlight.NewGroup("CursorColumn").WithArguments(subCurArgs))
	cs.AddGroup(highlight.NewGroup("CursorLine").WithArguments(subCurArgs))
	cs.AddGroup(highlight.NewGroup("ColorColumn").WithArguments(subCurArgs))

	cs.AddGroup(highlight.NewGroup("lCursor").WithColorSet(baseCS).WithAttrSet(revAS))
	cs.AddGroup(highlight.NewGroup("CursorIM").WithColorSet(baseCS).WithAttrSet(revAS))

	cs.AddGroup(highlight.NewGroup("LineNr").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("CursorLineNr").WithColorSet(lightCS).WithAttrSet(boldAS))

	cs.AddGroup(highlight.NewGroup("Conceal").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("SpecialKey").WithArguments(lightArgs))

	// custom groups

	cs.AddGroup(highlight.NewGroup("Comment").WithArguments(lightArgs))
	cs.AddGroup(highlight.NewGroup("Statement").WithColorSet(baseCS).WithAttrSet(boldAS))
	cs.AddGroup(highlight.NewGroup("Type").WithArguments(subBoldArgs))
	cs.AddGroup(highlight.NewGroup("PreProc").WithArguments(subBoldArgs))
	cs.AddGroup(highlight.NewGroup("Identifier").WithArguments(subBoldArgs))

	cs.AddGroup(highlight.NewGroup("Special").WithAttrSet(baseAS).WithColorSet(highlight.ColorSet{Fg: darkFg, Bg: lightBg}))
	cs.AddGroup(highlight.NewGroup("Constant").WithAttrSet(baseAS).WithColorSet(highlight.ColorSet{Fg: baseFg, Bg: lightBg}))

	if err := cs.Marshal(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
