package blocklist

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block"
)

func TestNew(t *testing.T) {
	bl := New("")
	assert.Eq(t, bl.class, "")

	bl = New("system")
	assert.Eq(t, bl.class, "system")
}

func TestAddBlock(t *testing.T) {
	b := block.New("quando.unique.id", "system")
	bl := New("system")
	assert.Eq(t, len(bl.blocks), 0)
	bl.Add(b)
	assert.Eq(t, len(bl.blocks), 1)
}

func TestClass(t *testing.T) {
	bl := New("")
	assert.Eq(t, bl.CSSClass("quando-"), "quando-unknown")

	bl = New("system")
	assert.Eq(t, bl.CSSClass("quando-library-"), "quando-library-system")
}
