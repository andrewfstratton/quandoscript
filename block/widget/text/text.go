package text

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/block/widget"
)

type Text struct {
	widget.TextWidget
	txt string
}

func New(t string) *Text {
	return &Text{txt: t}
}

func (t *Text) String() (txt string) {
	txt = t.txt
	if t.Style.Italic {
		txt = Tag(txt, "i")
	}
	if t.Style.Bold {
		txt = Tag(txt, "b")
	}
	if t.Style.Iconify {
		txt = fmt.Sprintf("<%v>%v</%v>", `span class="iconify"`, txt, "span")
	}
	return
}

func Tag(txt string, tag string) string {
	return fmt.Sprintf("<%v>%v</%v>", tag, txt, tag)
}
