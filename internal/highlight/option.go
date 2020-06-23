package highlight

// Option is an option to modify Group.
type Option interface {
	Apply(*Group)
}

// OptionFunc is short hand to implements Option by a function. 
type OptionFunc func(*Group)

var _ Option = OptionFunc(nil)

// Apply implements Option interface.
func (fn OptionFunc) Apply(g *Group) {
	fn(g)
}

// Apply applies all options to a Group.
func (g *Group) Apply(opts ...Option) *Group {
	for _, o := range opts {
		if o == nil {
			continue
		}
		o.Apply(g)
	}
	return g
}
