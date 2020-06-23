package highlight

type Option interface {
	ApplyGroup(*Group) *Group
}

type OptionFunc func(*Group) *Group

var _ Option = OptionFunc(nil)

func (fn OptionFunc) ApplyGroup(g *Group) *Group {
	return fn(g)
}

func (g *Group) Apply(opts ...Option) *Group {
	for _, o := range opts {
		if o == nil {
			continue
		}
		g = o.ApplyGroup(g)
	}
	return g
}
