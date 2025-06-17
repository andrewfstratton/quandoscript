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

func TestFind(t *testing.T) {
	b, found := FindBlock("")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	b, found = FindBlock("system.log")
	assert.True(t, b == nil)
	assert.Eq(t, found, false)

	Block(&SimpleDefn{}) // don't keep reference here...
	b, found = FindBlock("system.log")
	assert.True(t, b != nil)
	assert.Eq(t, found, true)
}

type SimpleDefn2 struct {
	TypeName string `_:"display.show"`
	Class    string `_:"display"`
}

func TestClasses(t *testing.T) {
	Block(&SimpleDefn{})
	Block(&SimpleDefn2{})
	assert.Eq(t, len(Classes()), 2)
}
