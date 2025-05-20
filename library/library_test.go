package library

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/blocklist"
)

func TestNew(t *testing.T) {
	bl := blocklist.New("")
	assert.Eq(t, bl.Class(), "quando")

	bl = blocklist.New("system")
	assert.Eq(t, bl.Class(), "quando-system")
}
