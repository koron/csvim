package hlopt

import "github.com/koron/csvim/internal/highlight"

// Merge returns an Option to merge fields in Group except Name to "dst".
func Merge(g *highlight.Group) highlight.Option {
	if g == nil {
		return nop
	}
	return highlight.OptionFunc(func(dst *highlight.Group) {
		appendAttrs(&dst.Term, g.Term)
		appendTerms(&dst.Start, g.Start)
		appendTerms(&dst.Stop, g.Stop)
		appendAttrs(&dst.CTerm, g.CTerm)
		appendColorNr(&dst.CTermFg, g.CTermFg)
		appendColorNr(&dst.CTermBg, g.CTermBg)
		appendAttrs(&dst.GUI, g.GUI)
		if g.Font != "" {
			dst.Font = g.Font
		}
		appendColorName(&dst.GUIFg, g.GUIFg)
		appendColorName(&dst.GUIBg, g.GUIBg)
		appendColorName(&dst.GUISp, g.GUISp)
		if g.Default {
			dst.Default = true
		}
	})
}

func appendAttrs(dst *highlight.AttrList, src highlight.AttrList) {
	if len(src) == 0 {
		return
	}
	*dst = append(*dst, src...)
}

func appendTerms(dst *highlight.TermList, src highlight.TermList) {
	if len(src) == 0 {
		return
	}
	*dst = append(*dst, src...)
}

func appendColorNr(dst *highlight.ColorNr, src highlight.ColorNr) {
	if src == "" {
		return
	}
	*dst = src
}

func appendColorName(dst *highlight.ColorName, src highlight.ColorName) {
	if src == "" {
		return
	}
	*dst = src
}
