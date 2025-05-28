package library

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/run/param"
)

func TestNew(t *testing.T) {
	b := NewBlockType("system.log", "system", nil)
	assert.True(t, b != nil)
	assert.True(t, blocklists != nil)
	assert.True(t, blocks != nil)
	b = NewBlockType("system.log", "system", nil)
	assert.True(t, b == nil)
}

func TestFind(t *testing.T) {
	b, found := FindBlockType("")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	b, found = FindBlockType("display.log")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	_ = NewBlockType("display.log", "display", nil)
	b, found = FindBlockType("display.log")
	assert.True(t, b != nil)
	assert.Eq(t, found, true)
}

func TestString(t *testing.T) {
	params := param.Params{}
	var none *string
	params.String("txt", none)
	assert.True(t, none == nil)
}

func TestClasses(t *testing.T) {
	_ = NewBlockType("system.log", "system", nil)
	_ = NewBlockType("display.show", "display", nil)
	_ = NewBlockType("debug", "", nil)
	assert.Eq(t, len(Classes()), 3)
}
