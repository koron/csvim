package main

import (
	"log"
	"os"

	"github.com/koron/csvim/internal/colorscheme"
	"github.com/koron/csvim/internal/highlight"
)

func main() {
	cs := &colorscheme.ColorScheme{
		Name:       "morning",
		Background: colorscheme.Light,
	}
	cs.AddGroup(&highlight.Group{Name: "Normal", CTermFg: "Black", CTermBg: "LightGrey", GUIFg: "Black", GUIBg: "grey90"})
	cs.AddGroup(&highlight.Group{Name: "ErrorMsg", Term: highlight.AttrList{highlight.Standout}, CTermBg: "DarkRed", CTermFg: "White", GUIBg: "Red", GUIFg: "White"})
	err := cs.Marshal(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
