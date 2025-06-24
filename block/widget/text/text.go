package text

import (
	"github.com/andrewfstratton/quandoscript/block/widget"
)

type Text struct {
	Txt     string
	Italic  bool
	Bold    bool
	Iconify bool
}

func New() *Text {
	return &Text{}
}

func (t *Text) Html() (txt string) {
	txt = t.Txt
	if t.Italic {
		txt = widget.TagText(txt, "i")
	}
	if t.Bold {
		txt = widget.TagText(txt, "b")
	}
	if t.Iconify {
		txt = widget.OpenCloseTag(txt, `span class="iconify"`, "span")
	}
	return
}
