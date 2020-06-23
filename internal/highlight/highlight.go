package highlight

// Command is a literal of "highlight" command which used for marshaling.
// Overriding it with "hi", makes Marshal() generate smaller script file.
var Command = "highlight"

// Normal is "Normal" default group name.
const Normal = "Normal"

// DefaultGroupNames is set of default group names.
var DefaultGroupNames = []string{
	Normal,
	"ColorColumn",
	"Conceal",
	"Cursor",
	"CursorColumn",
	"CursorIM",
	"CursorLine",
	"CursorLineNr",
	"DiffAdd",
	"DiffChange",
	"DiffDelete",
	"DiffText",
	"Directory",
	"EndOfBuffer",
	"ErrorMsg",
	"FoldColumn",
	"Folded",
	"IncSearch",
	"LineNr",
	"LineNrAbove",
	"LineNrBelow",
	"MatchParen",
	"ModeMsg",
	"MoreMsg",
	"NonText",
	"Pmenu",
	"PmenuSbar",
	"PmenuSel",
	"PmenuThumb",
	"Question",
	"QuickFixLine",
	"Search",
	"SignColumn",
	"SpecialKey",
	"SpellBad",
	"SpellCap",
	"SpellLocal",
	"SpellRare",
	"StatusLine",
	"StatusLineNC",
	"StatusLineTerm",
	"StatusLineTermNC",
	"TabLine",
	"TabLineFill",
	"TabLineSel",
	"Terminal",
	"Title",
	"VertSplit",
	"Visual",
	"VisualNOS",
	"WarningMsg",
	"WildMenu",
	"lCursor",
}

var defaultGroups map[string]struct{}

func init() {
	defaultGroups = make(map[string]struct{}, len(DefaultGroupNames))
	for _, g := range DefaultGroupNames {
		defaultGroups[g] = struct{}{}
	}
}

// IsDefaultGroup checks names is for default group or not.
func IsDefaultGroup(name string) bool {
	_, ok := defaultGroups[name]
	return ok
}
