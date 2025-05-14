package text

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestTextSimple(t *testing.T) {
	txt := New("")
	assert.Eq(t, txt.String(), "")

	txt = New("Hello")
	assert.Eq(t, txt.String(), "Hello")
	txt.Italic()
	assert.Eq(t, txt.String(), "<i>Hello</i>")
	txt.Bold()
	assert.Eq(t, txt.String(), "<b><i>Hello</i></b>")
	txt.Iconify()
	assert.Eq(t, txt.String(), `<span class="iconify"><b><i>Hello</i></b></span>`)

	txt = New("Hi Bob")
	txt.Bold().Italic().Iconify()
	assert.Eq(t, txt.String(), `<span class="iconify"><i><b>Hi Bob</b></i></span>`)
}
