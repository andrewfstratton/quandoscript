package text

import (
	"fmt"
)

type Text struct {
	Italic  bool
	Bold    bool
	Iconify bool
	Txt     string
}

func New(t string) *Text {
	return &Text{Txt: t}
}

func (t *Text) Html() (txt string) {
	txt = t.Txt
	if t.Italic {
		txt = TagText(txt, "i")
	}
	if t.Bold {
		txt = TagText(txt, "b")
	}
	if t.Iconify {
		txt = OpenCloseTag(txt, `span class="iconify"`, "span")
	}
	return
}

func TagText(txt string, tag string) string {
	return OpenCloseTag(txt, tag, tag)
}

func OpenCloseTag(txt string, open string, close string) string {
	return fmt.Sprintf("<%v>%v</%v>", open, txt, close)
}
