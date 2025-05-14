package text

import (
	"fmt"
)

type Text struct {
	italic  bool
	bold    bool
	iconify bool
	txt     string
}

func New(t string) *Text {
	return &Text{txt: t}
}

func (t *Text) Html() (txt string) {
	txt = t.txt
	if t.italic {
		txt = Tag(txt, "i")
	}
	if t.bold {
		txt = Tag(txt, "b")
	}
	if t.iconify {
		txt = fmt.Sprintf("<%v>%v</%v>", `span class="iconify"`, txt, "span")
	}
	return
}

func (t *Text) Italic() *Text {
	t.italic = true
	return t
}

func (t *Text) Bold() *Text {
	t.bold = true
	return t
}

func (t *Text) Iconify() *Text {
	t.iconify = true
	return t
}

func Tag(txt string, tag string) string {
	return fmt.Sprintf("<%v>%v</%v>", tag, txt, tag)
}
