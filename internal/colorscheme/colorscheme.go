package colorscheme

import (
	"errors"
	"fmt"
	"io"

	"github.com/koron/csvim/internal/highlight"
)

type Background string

const (
	Dark  Background = "dark"
	Light            = "light"
)

type groups []*highlight.Group

func (gs groups) marshal(w io.Writer) error {
	for _, g := range gs {
		if err := g.Marshal(w); err != nil {
			return err
		}
	}
	return nil
}

type links []*highlight.Link

func (ls links) marshal(w io.Writer) error {
	for _, l := range ls {
		if err := l.Marshal(w); err != nil {
			return err
		}
	}
	return nil
}

type ColorScheme struct {
	Name       string
	Background Background
	Normal     *highlight.Group

	groupIdx map[string]int
	groups   groups

	linkIdx map[string]int
	links   []*highlight.Link
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
`, cs.Name, cs.Background); err != nil {
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
		if _, err := fmt.Fprintln(w, "\n\" custom groups\n"); err != nil {
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
	if _, err := fmt.Fprintln(w, "\n\" links\n"); err != nil {
		return err
	}
	// FIXME: sort links by dependencies.
	for _, ln := range cs.links {
		if ln == nil {
			continue
		}
		if err := ln.Marshal(w); err != nil {
			return err
		}
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
