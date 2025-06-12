package text

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block/widget"
)

func TestTextSimple(t *testing.T) {
	txt := New("")
	assert.Eq(t, txt.Html(), "")

	txt = New("Hello")
	assert.Eq(t, txt.Html(), "Hello")
	widget.Setup(txt, "", `italic:"true"`)
	assert.Eq(t, txt.Italic, true)
	assert.Eq(t, txt.Html(), "<i>Hello</i>")
	widget.Setup(txt, "", `bold:"true"`)
	widget.Setup(txt, "", `iconify:"true"`)
	assert.Eq(t, txt.Html(), `<span class="iconify"><b><i>Hello</i></b></span>`)
	txt = &Text{}
	widget.Setup(txt, "", `txt:"Hi Bob" italic:"false" iconify:"true" bold:"true"`)
	assert.Eq(t, txt.Html(), `<span class="iconify"><b>Hi Bob</b></span>`) // n.b, order is not preserved?!
}
