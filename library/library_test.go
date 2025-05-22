package library

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestNew(t *testing.T) {
	b := NewBlockType("system.log", "system")
	assert.True(t, b != nil)
	assert.True(t, blocklists != nil)
	assert.True(t, blocks != nil)
	b = NewBlockType("system.log", "system")
	assert.True(t, b == nil)
}

func TestFind(t *testing.T) {
	b, found := FindBlockType("")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	b, found = FindBlockType("display.log")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	_ = NewBlockType("display.log", "display")
	b, found = FindBlockType("display.log")
	assert.True(t, b != nil)
	assert.Eq(t, found, true)

}

func TestClasses(t *testing.T) {
	_ = NewBlockType("system.log", "system")
	_ = NewBlockType("display.show", "display")
	_ = NewBlockType("debug", "")
	assert.Eq(t, len(Classes()), 3)
}
