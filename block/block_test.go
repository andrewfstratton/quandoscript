package block

import (
	"testing"

	"quandoscript/assert"
	"quandoscript/block/widget/numberinput"
	"quandoscript/block/widget/percentinput"
	"quandoscript/block/widget/stringinput"
	"quandoscript/block/widget/text"
)

type FailDefn struct {
	TypeName string
	Class    string
}

func TestEmpty(t *testing.T) {
	block := New(&FailDefn{}) // test with empty blocktype here
	assert.Eq(t, block, nil)  // n.b. will panic when not testing
}

type TagDefn struct {
	TypeName string `_:"system.log"`
	Class    string `_:"sys"`
}

func TestEmptyTag(t *testing.T) {
	block := New(&TagDefn{}) // test with struct initialised blocktype and class here
	assert.Eq(t, block.Class, "quando-sys")
	assert.Eq(t, block.TypeName, "system.log")
}

type TextDefn struct {
	TypeName string    `_:"system.log"`
	Class    string    `_:"system"`
	_        text.Text `txt:"Log "`
}

func TestNewFull(t *testing.T) {
	block := New(&TextDefn{})
	assert.Eq(t, block.Replace("{{ .Class }}"), "quando-system")
	assert.Eq(t, block.Replace("{{ .TypeName }}"), "system.log")
	assert.Eq(t, block.Replace("{{ .Widgets }}"), "Log ")
	assert.Eq(t, block.Replace("{{ .Params }}"), "")
}

type InputDefn struct {
	TypeName string    `_:"system.log"`
	Class    string    `_:"system"`
	_        text.Text `txt:"Log "`
	Name     stringinput.StringInput
	Num      numberinput.NumberInput
}

func TestFullInput(t *testing.T) {
	block := New(&InputDefn{})
	assert.Eq(t, block.Replace("{{ .Class }}"), "quando-system")
	assert.Eq(t, block.Replace("{{ .TypeName }}"), "system.log")
	assert.Eq(t, block.Replace("{{ .Widgets }}"), `Log &quot;<input data-quando-name='Name' type='text'/>&quot;<input data-quando-name='Num' type='number'/>`)
	assert.Eq(t, block.Replace("{{ .Params }}"), `Name"${Name}",Num#${Num}`)
}

type PercentDefn struct {
	TypeName string    `_:"display.width"`
	Class    string    `_:"display"`
	_        text.Text `txt:"Width "`
	Width    percentinput.PercentInput
}

func TestPercentInput(t *testing.T) {
	block := New(&PercentDefn{})
	assert.Eq(t, block.Replace("{{ .Class }}"), "quando-display")
	assert.Eq(t, block.Replace("{{ .TypeName }}"), "display.width")
	assert.Eq(t, block.Replace("{{ .Widgets }}"), `Width <input data-quando-name='Width' type='number' min='0' max='100'/>%`)
	assert.Eq(t, block.Replace("{{ .Params }}"), `Width#${Width}`)
}
