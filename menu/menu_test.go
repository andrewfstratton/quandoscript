package menu

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block"
)

func TestNew(t *testing.T) {
	menu := New("")
	assert.Eq(t, menu.Class, "")

	menu = New("system")
	assert.Eq(t, menu.Class, "system")
}

type SimpleDefn struct {
	TypeName string `_:"system.log"`
	Class    string `_:"system"`
}

func TestAddBlock(t *testing.T) {
	b := block.CreateFromDefinition(&SimpleDefn{})
	menu := New("system")
	assert.Eq(t, len(menu.Blocks), 0)
	menu.Add(b)
	assert.Eq(t, len(menu.Blocks), 1)
}

func TestClass(t *testing.T) {
	menu := New("")
	assert.Eq(t, menu.CSSClass("quando-"), "quando-unknown")

	menu = New("system")
	assert.Eq(t, menu.CSSClass("quando-library-"), "quando-library-system")
}
