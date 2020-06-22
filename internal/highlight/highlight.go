package highlight

var Command = "highlight"

const (
	Normal           = "Normal"
	ColorColumn      = "ColorColumn"
	Conceal          = "Conceal"
	Cursor           = "Cursor"
	CursorColumn     = "CursorColumn"
	CursorIM         = "CursorIM"
	CursorLine       = "CursorLine"
	CursorLineNr     = "CursorLineNr"
	DiffAdd          = "DiffAdd"
	DiffChange       = "DiffChange"
	DiffDelete       = "DiffDelete"
	DiffText         = "DiffText"
	Directory        = "Directory"
	EndOfBuffer      = "EndOfBuffer"
	ErrorMsg         = "ErrorMsg"
	FoldColumn       = "FoldColumn"
	Folded           = "Folded"
	IncSearch        = "IncSearch"
	LineNr           = "LineNr"
	LineNrAbove      = "LineNrAbove"
	LineNrBelow      = "LineNrBelow"
	MatchParen       = "MatchParen"
	ModeMsg          = "ModeMsg"
	MoreMsg          = "MoreMsg"
	NonText          = "NonText"
	Pmenu            = "Pmenu"
	PmenuSbar        = "PmenuSbar"
	PmenuSel         = "PmenuSel"
	PmenuThumb       = "PmenuThumb"
	Question         = "Question"
	QuickFixLine     = "QuickFixLine"
	Search           = "Search"
	SignColumn       = "SignColumn"
	SpecialKey       = "SpecialKey"
	SpellBad         = "SpellBad"
	SpellCap         = "SpellCap"
	SpellLocal       = "SpellLocal"
	SpellRare        = "SpellRare"
	StatusLine       = "StatusLine"
	StatusLineNC     = "StatusLineNC"
	StatusLineTerm   = "StatusLineTerm"
	StatusLineTermNC = "StatusLineTermNC"
	TabLine          = "TabLine"
	TabLineFill      = "TabLineFill"
	TabLineSel       = "TabLineSel"
	Terminal         = "Terminal"
	Title            = "Title"
	VertSplit        = "VertSplit"
	Visual           = "Visual"
	VisualNOS        = "VisualNOS"
	WarningMsg       = "WarningMsg"
	WildMenu         = "WildMenu"
	LCursor          = "lCursor"
)

var DefaultGroupNames = []string{
	Normal,
	ColorColumn,
	Conceal,
	Cursor,
	CursorColumn,
	CursorIM,
	CursorLine,
	CursorLineNr,
	DiffAdd,
	DiffChange,
	DiffDelete,
	DiffText,
	Directory,
	EndOfBuffer,
	ErrorMsg,
	FoldColumn,
	Folded,
	IncSearch,
	LineNr,
	LineNrAbove,
	LineNrBelow,
	MatchParen,
	ModeMsg,
	MoreMsg,
	NonText,
	Pmenu,
	PmenuSbar,
	PmenuSel,
	PmenuThumb,
	Question,
	QuickFixLine,
	Search,
	SignColumn,
	SpecialKey,
	SpellBad,
	SpellCap,
	SpellLocal,
	SpellRare,
	StatusLine,
	StatusLineNC,
	StatusLineTerm,
	StatusLineTermNC,
	TabLine,
	TabLineFill,
	TabLineSel,
	Terminal,
	Title,
	VertSplit,
	Visual,
	VisualNOS,
	WarningMsg,
	WildMenu,
	LCursor,
}

var defaultGroups map[string]struct{}

func init() {
	defaultGroups = make(map[string]struct{}, len(DefaultGroupNames))
	for _, g := range DefaultGroupNames {
		defaultGroups[g] = struct{}{}
	}
}

func IsDefaultGroup(name string) bool {
	_, ok := defaultGroups[name]
	return ok
}
