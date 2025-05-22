package library

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestNew(t *testing.T) {
	b := NewBlock("system.log", "system")
	assert.True(t, b != nil)
	assert.True(t, blocklists != nil)
	assert.True(t, blocks != nil)
	b = NewBlock("system.log", "system")
	assert.True(t, b == nil)
}

func TestFind(t *testing.T) {
	b, found := FindBlock("")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	b, found = FindBlock("display.log")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	_ = NewBlock("display.log", "display")
	b, found = FindBlock("display.log")
	assert.True(t, b != nil)
	assert.Eq(t, found, true)

}

func TestClasses(t *testing.T) {
	_ = NewBlock("system.log", "system")
	_ = NewBlock("display.show", "display")
	_ = NewBlock("debug", "")
	assert.Eq(t, len(Classes()), 3)
}
