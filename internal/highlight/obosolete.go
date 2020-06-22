package highlight

type AttrSet struct {
	Term  AttrList
	CTerm AttrList
	GUI   AttrList
}

type Arguments struct {
	AttrSet  AttrSet
	ColorSet ColorSet
}

func (g *Group) WithAttrSet(attrSet AttrSet) *Group {
	g.Term = attrSet.Term
	g.CTerm = attrSet.CTerm
	g.GUI = attrSet.GUI
	return g
}

func (g *Group) WithColorSet(colorSet ColorSet) *Group {
	return g.WithFg(colorSet.Fg).WithBg(colorSet.Bg).WithGUISp(colorSet.Sp)
}

func (g *Group) WithArguments(args Arguments) *Group {
	return g.WithAttrSet(args.AttrSet).WithColorSet(args.ColorSet)
}

type ColorSet struct {
	Fg Color
	Bg Color
	Sp Color
}
