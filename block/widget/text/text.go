package text

import (
	"github.com/andrewfstratton/quandoscript/block/widget"
)

type Text struct {
	Txt     string
	Italic  bool
	Bold    bool
	Iconify bool
	Show    string
	Hover   bool
}

func New() *Text {
	return &Text{}
}

func (t *Text) Html() (txt string) {
	txt = t.Txt
	if t.Hover {
		txt = widget.OpenCloseTag(txt, `span class="hover-display"`, "span")
	}
	if t.Italic {
		txt = widget.TagText(txt, "i")
	}
	if t.Bold {
		txt = widget.TagText(txt, "b")
	}
	if t.Iconify {
		txt = widget.OpenCloseTag(txt, `span class="iconify"`, "span")
	}
	if t.Show != "" {
		txt = widget.OpenCloseTag(txt, `span data-quando-toggle="`+t.Show+`"`, "span")
	}
	return
}
