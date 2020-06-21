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

func (g *Group) WithFg(c Color) *Group {
	g.CTermFg = c.Nr
	g.GUIFg = c.Name
	return g
}

func (g *Group) WithBg(c Color) *Group {
	g.CTermBg = c.Nr
	g.GUIBg = c.Name
	return g
}

func (g *Group) WithSp(c Color) *Group {
	g.GUISp = c.Name
	return g
}

func (g *Group) WithColorSet(colorSet ColorSet) *Group {
	return g.WithFg(colorSet.Fg).WithBg(colorSet.Bg).WithSp(colorSet.Sp)
}

func (g *Group) WithArguments(args Arguments) *Group {
	return g.WithAttrSet(args.AttrSet).WithColorSet(args.ColorSet)
}

type ColorSet struct {
	Fg Color
	Bg Color
	Sp Color
}

type Color struct {
	Nr   ColorNr
	Name ColorName
}
