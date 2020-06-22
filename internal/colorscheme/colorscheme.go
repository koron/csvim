package colorscheme

import (
	"errors"
	"fmt"
	"io"

	"github.com/koron/csvim/internal/highlight"
)

// WarnDefaultGroups is a flag to list up unused default highlight groups on
// Marshal().
var WarnDefaultGroups bool

// Background is value for 'background' option of vim.
type Background string

const (
	// Dark is constant "dark" for Background.
	Dark Background = "dark"

	// Light is constant "light" for Background.
	Light Background = "light"
)

type groups []*highlight.Group

func (gs groups) marshal(w io.Writer) error {
	for _, g := range gs {
		if g == nil {
			continue
		}
		if err := g.Marshal(w); err != nil {
			return err
		}
	}
	return nil
}

type links []*highlight.Link

func (ls links) marshal(w io.Writer) error {
	for _, l := range ls {
		if l == nil {
			continue
		}
		if err := l.Marshal(w); err != nil {
			return err
		}
	}
	return nil
}

// ColorScheme is a type for Vim's colorscheme.
type ColorScheme struct {
	Name       string
	Background Background
	Normal     *highlight.Group

	groupIdx map[string]int
	groups   groups

	linkIdx map[string]int
	links   links
}

// New creates a new ColorScheme.
func New(name string) *ColorScheme {
	return &ColorScheme{Name: name}
}

// WithBackground sets Background.
func (cs *ColorScheme) WithBackground(bg Background) *ColorScheme {
	cs.Background = bg
	return cs
}

func (cs *ColorScheme) removeGroup(name string) bool {
	x, ok := cs.groupIdx[name]
	if !ok {
		return false
	}
	cs.groups[x] = nil
	return true
}

func (cs *ColorScheme) removeLink(name string) bool {
	x, ok := cs.linkIdx[name]
	if !ok {
		return false
	}
	cs.links[x] = nil
	return true
}

// AddGroup adds a Group to colorscheme.
func (cs *ColorScheme) AddGroup(g *highlight.Group) error {
	if g.Name == "" {
		return errors.New("group should have name")
	}
	if g.Name == highlight.Normal {
		cs.Normal = g
		return nil
	}
	cs.removeLink(g.Name)
	if !cs.removeGroup(g.Name) && cs.groupIdx == nil {
		cs.groupIdx = make(map[string]int)
	}
	cs.groupIdx[g.Name] = len(cs.groups)
	cs.groups = append(cs.groups, g)
	return nil
}

// Group adds a Group with name and return it. This will panic when invalid
// name is passed.
func (cs *ColorScheme) Group(name string) *highlight.Group {
	g := highlight.NewGroup(name)
	err := cs.AddGroup(g)
	if err != nil {
		panic("failed to add a group: " + err.Error())
	}
	return g
}

// AddLink adds a Link to colorscheme.
func (cs *ColorScheme) AddLink(l *highlight.Link) error {
	if l.From == "" || l.From == highlight.Normal {
		return errors.New("link should have name and it isn't \"Normal\"")
	}
	cs.removeGroup(l.From)
	if !cs.removeLink(l.From) && cs.linkIdx == nil {
		cs.linkIdx = make(map[string]int)
	}
	cs.linkIdx[l.From] = len(cs.links)
	cs.links = append(cs.links, l)
	return nil
}

// Link adds a Link with names ("from" and "to") and return it. This will panic
// when invalid name is passed.
func (cs *ColorScheme) Link(from, to string) *highlight.Link {
	ln := highlight.NewLink(from, to)
	err := cs.AddLink(ln)
	if err != nil {
		panic("failed to add a link: " + err.Error())
	}
	return ln
}

// Has checks color scheme has a group with the name.
func (cs *ColorScheme) Has(name string) bool {
	if name == highlight.Normal {
		return cs.Normal != nil
	}
	if _, ok := cs.groupIdx[name]; ok {
		return true
	}
	_, ok := cs.linkIdx[name]
	return ok
}

// Marshal writes down Vim's colorscheme script to io.Writer.
func (cs *ColorScheme) Marshal(w io.Writer) error {
	bg := cs.Background
	if bg == "" {
		bg = Dark
	}
	if _, err := fmt.Fprintf(w, `set background=%[2]s
highlight clear
if exists("syntax_on")
  syntax reset
endif
let g:colors_name = %[1]q
`, cs.Name, bg); err != nil {
		return nil
	}
	if err := cs.marshalNormal(w); err != nil {
		return err
	}
	if err := cs.marshalGroups(w); err != nil {
		return err
	}
	if err := cs.marshalLinks(w); err != nil {
		return err
	}
	if WarnDefaultGroups {
		names := cs.undefinedDefaultGroups()
		if len(names) > 0 {
			if _, err := fmt.Fprint(w, "\n\" WARNING: undefined default groups:\n"); err != nil {
				return err
			}
			for _, n := range names {
				if _, err := fmt.Fprintf(w, "\"  - %s\n", n); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (cs *ColorScheme) marshalNormal(w io.Writer) error {
	if cs.Normal == nil {
		if _, err := fmt.Fprint(w, "\n\" WARNING: Normal group is missing!\n"); err != nil {
			return err
		}
		return nil
	}
	if _, err := fmt.Fprint(w, "\n\" Normal group at first\n"); err != nil {
		return err
	}
	if err := cs.Normal.Marshal(w); err != nil {
		return err
	}
	return nil
}

func (cs *ColorScheme) marshalGroups(w io.Writer) error {
	gs0 := cs.defaultGroups()
	if len(gs0) > 0 {
		if _, err := fmt.Fprint(w, "\n\" default groups\n"); err != nil {
			return err
		}
		if err := gs0.marshal(w); err != nil {
			return err
		}
	}
	gs1 := cs.customGroups()
	if len(gs1) > 0 {
		if _, err := fmt.Fprint(w, "\n\" custom groups\n"); err != nil {
			return err
		}
		if err := gs1.marshal(w); err != nil {
			return err
		}
	}
	return nil
}

func (cs *ColorScheme) marshalLinks(w io.Writer) error {
	if len(cs.links) == 0 {
		return nil
	}
	if _, err := fmt.Fprint(w, "\n\" links\n"); err != nil {
		return err
	}
	// FIXME: sort links by dependencies.
	if err := cs.links.marshal(w); err != nil {
		return err
	}
	return nil
}

func (cs *ColorScheme) getGroup(name string) *highlight.Group {
	x, ok := cs.groupIdx[name]
	if !ok {
		return nil
	}
	return cs.groups[x]
}

func (cs *ColorScheme) getLink(name string) *highlight.Link {
	x, ok := cs.linkIdx[name]
	if !ok {
		return nil
	}
	return cs.links[x]
}

func (cs *ColorScheme) defaultGroups() groups {
	if len(cs.groupIdx) == 0 {
		return nil
	}
	gs := make(groups, 0, len(cs.groupIdx))
	for _, n := range highlight.DefaultGroupNames {
		if g := cs.getGroup(n); g != nil {
			gs = append(gs, g)
		}
	}
	return gs
}

func (cs *ColorScheme) customGroups() groups {
	if len(cs.groupIdx) == 0 {
		return nil
	}
	gs := make(groups, 0, len(cs.groupIdx))
	for _, g := range cs.groups {
		if g != nil && !highlight.IsDefaultGroup(g.Name) {
			gs = append(gs, g)
		}
	}
	return gs
}

func (cs *ColorScheme) undefinedDefaultGroups() []string {
	names := make([]string, 0, len(cs.groupIdx))
	for _, n := range highlight.DefaultGroupNames {
		if g := cs.getGroup(n); g != nil {
			continue
		}
		if l := cs.getLink(n); l != nil {
			continue
		}
		if n == highlight.Normal && cs.Normal != nil {
			continue
		}
		names = append(names, n)
	}
	return names
}
