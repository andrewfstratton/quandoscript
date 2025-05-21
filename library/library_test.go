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
	b, found := Block("")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	b, found = Block("system.log")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	_ = NewBlock("system.log", "system")
	b, found = Block("system.log")
	assert.True(t, b != nil)
	assert.Eq(t, found, true)

}
