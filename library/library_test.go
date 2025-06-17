package library

import (
	"testing"

	"quandoscript/assert"
)

type SimpleDefn struct {
	TypeName string `_:"system.log"`
	Class    string `_:"system"`
}

func TestNew(t *testing.T) {
	assert.True(t, menus != nil)
	assert.True(t, blocks != nil)
	b := Block(&SimpleDefn{})
	assert.True(t, b != nil)
}

type SimpleDefn2 struct {
	TypeName string `_:"display.show"`
	Class    string `_:"display"`
}

func TestFind(t *testing.T) {
	b, found := FindBlock("")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	b, found = FindBlock("display.show")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	Block(&SimpleDefn2{}) // don't keep reference here...
	b, found = FindBlock("display.show")
	assert.True(t, b != nil)
	assert.Eq(t, found, true)
}

func TestClasses(t *testing.T) {
	Block(&SimpleDefn{})
	Block(&SimpleDefn2{})
	assert.Eq(t, len(Classes()), 2)
}
