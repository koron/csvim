package colorscheme

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/koron/csvim/internal/highlight"
)

func ReadFile(name string) (*ColorScheme, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Read(f)
}

var (
	rxBg   = regexp.MustCompile(`^set (?:background|bg)=(.*)`)
	rxName = regexp.MustCompile(`^let\s+(?:g:)?colors_name\s*=\s*"([^"]*)"`)
	rxWS   = regexp.MustCompile(`\s+`)
)

func Read(rd io.Reader) (*ColorScheme, error) {
	r := bufio.NewReader(rd)
	cs := new(ColorScheme)

	parseLine := func(s string) error {
		s = strings.TrimSpace(s)
		if m := rxBg.FindStringSubmatch(s); m != nil {
			cs.Background = Background(m[1])
			return nil
		}
		if m := rxName.FindStringSubmatch(s); m != nil {
			cs.Name = m[1]
			return nil
		}
		if !strings.HasPrefix(s, "hi") {
			return nil
		}

		items := rxWS.Split(s, -1)
		if !strings.HasPrefix("highlight", items[0]) {
			return nil
		}
		if len(items) < 2 {
			return nil
		}
		first, items := items[1], items[2:]
		var defaultFlag = false

		if first == "clear" {
			return nil
		}
		if first == "default" {
			defaultFlag = true
			if len(items) < 1 {
				return nil
			}
			first, items = items[0], items[1:]
		}

		if first == "link" {
			if len(items) < 2 {
				return nil
			}
			if err := cs.AddLink(&highlight.Link{
				From:    items[0],
				To:      items[1],
				Default: defaultFlag,
			}); err != nil {
				return err
			}
			return nil
		}

		g := &highlight.Group{Name: first, Default: defaultFlag}
		for _, item := range items {
			if item == "NONE" {
				return nil
			}
			x := strings.IndexRune(item, '=')
			if x == -1 {
				continue
			}
			n, v := item[:x], item[x+1:]
			switch n {
			case "term":
				g.Term = toAttrList(strings.Split(v, ","))
			case "start":
				g.Start = strings.Split(v, ",")
			case "stop":
				g.Stop = strings.Split(v, ",")
			case "cterm":
				g.CTerm = toAttrList(strings.Split(v, ","))
			case "ctermfg":
				g.CTermFg = highlight.ColorNr(v)
			case "ctermbg":
				g.CTermBg = highlight.ColorNr(v)
			case "gui":
				g.GUI = toAttrList(strings.Split(v, ","))
			case "font":
				g.Font = v
			case "guifg":
				g.GUIFg = highlight.ColorName(v)
			case "guibg":
				g.GUIBg = highlight.ColorName(v)
			case "guisp":
				g.GUISp = highlight.ColorName(v)
			}
		}
		if err := cs.AddGroup(g); err != nil {
			return err
		}
		return nil
	}

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, err
			}
			if s == "" {
				return cs, nil
			}
		}
		if err := parseLine(s); err != nil {
			return nil, err
		}
	}
}

func toAttrList(src []string) highlight.AttrList {
	dst := make(highlight.AttrList, len(src))
	for i, v := range src {
		dst[i] = highlight.Attr(v)
	}
	return dst

}
