package library

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestNew(t *testing.T) {
	lib := NewBlock("system.log", "system")
	assert.True(t, lib != nil)
}
