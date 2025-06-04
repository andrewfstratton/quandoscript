package library

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block"
)

func TestNew(t *testing.T) {
	b := block.AddNew("", "system", nil)
	assert.True(t, b == nil)
	assert.True(t, menus != nil)
	assert.True(t, blocks != nil)
	b = block.AddNew("system.log", "system", nil)
	assert.True(t, b != nil)
}

func TestFind(t *testing.T) {
	b, found := FindBlock("")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	b, found = FindBlock("display.log")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	block.AddNew("display.log", "display", nil)
	b, found = FindBlock("display.log")
	assert.True(t, b != nil)
	assert.Eq(t, found, true)
}

func TestString(t *testing.T) {
	params := param.Params{}
	none := param.StringParam{Lookup: "", Val: ""}
	none.Update(params)
	assert.Eq(t, none.Val, "")
}

func TestClasses(t *testing.T) {
	block.AddNew("system.log", "system", nil)
	block.AddNew("display.show", "display", nil)
	block.AddNew("debug", "", nil)
	assert.Eq(t, len(Classes()), 3)
}
