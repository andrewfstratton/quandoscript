package text

import (
	"testing"

	"quando/quandoscript/assert"
)

func TestTextSimple(t *testing.T) {
	txt := New("")
	assert.Eq(t, txt.Html(), "")

	txt = New("Hello")
	assert.Eq(t, txt.Html(), "Hello")
	txt.Italic()
	assert.Eq(t, txt.Html(), "<i>Hello</i>")
	txt.Bold()
	assert.Eq(t, txt.Html(), "<b><i>Hello</i></b>")
	txt.Iconify()
	assert.Eq(t, txt.Html(), `<span class="iconify"><b><i>Hello</i></b></span>`)

	txt = New("Hi Bob")
	txt.Bold().Italic().Iconify() // n.b, order is not preserved?!
	assert.Eq(t, txt.Html(), `<span class="iconify"><b><i>Hi Bob</i></b></span>`)
}
