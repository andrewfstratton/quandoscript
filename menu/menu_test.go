package menu

import (
	"testing"

	"quandoscript/assert"
	"quandoscript/block"
)

func TestNew(t *testing.T) {
	menu := New("")
	assert.Eq(t, menu.class, "")

	menu = New("system")
	assert.Eq(t, menu.class, "system")
}

type SimpleDefn struct {
	TypeName string `_:"system.log"`
	Class    string `_:"system"`
}

func TestAddBlock(t *testing.T) {
	b := block.New(&SimpleDefn{})
	menu := New("system")
	assert.Eq(t, len(menu.blocks), 0)
	menu.Add(b)
	assert.Eq(t, len(menu.blocks), 1)
}

func TestClass(t *testing.T) {
	menu := New("")
	assert.Eq(t, menu.CSSClass("quando-"), "quando-unknown")

	menu = New("system")
	assert.Eq(t, menu.CSSClass("quando-library-"), "quando-library-system")
}
