package block

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
	"github.com/andrewfstratton/quandoscript/block/widget/percentinput"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
)

type FailDefn struct {
	TypeName string
	Class    string
}

func TestEmpty(t *testing.T) {
	block := CreateFromDefinition(&FailDefn{}) // test with empty blocktype here
	assert.Eq(t, block, nil)                   // n.b. will panic when not testing
}

type TagDefn struct {
	TypeName string `_:"system.log"`
	Class    string `_:"sys"`
}

func TestEmptyTag(t *testing.T) {
	block := CreateFromDefinition(&TagDefn{}) // test with struct initialised blocktype and class here
	assert.Eq(t, block.Class, "quando-sys")
	assert.Eq(t, block.TypeName, "system.log")
}

type TextDefn struct {
	TypeName string    `_:"system.log"`
	Class    string    `_:"system"`
	_        text.Text `txt:"Log "`
}

func TestNewFull(t *testing.T) {
	block := CreateFromDefinition(&TextDefn{})
	assert.Eq(t, block.Replace("{{ .Class }}"), "quando-system")
	assert.Eq(t, block.Replace("{{ .TypeName }}"), "system.log")
	assert.Eq(t, block.Replace("{{ .Widgets }}"), "Log ")
	assert.Eq(t, block.Replace("{{ .Params }}"), "")
}

type InputDefn struct {
	TypeName string    `_:"system.log"`
	Class    string    `_:"system"`
	_        text.Text `txt:"Log "`
	Name     stringinput.String
	Num      numberinput.Number
}

func TestFullInput(t *testing.T) {
	block := CreateFromDefinition(&InputDefn{})
	assert.Eq(t, block.Replace("{{ .Class }}"), "quando-system")
	assert.Eq(t, block.Replace("{{ .TypeName }}"), "system.log")
	assert.Eq(t, block.Replace("{{ .Widgets }}"), `Log &quot;<input data-quando-name='Name' type='text'/>&quot;<input data-quando-name='Num' type='number'/>`)
	assert.Eq(t, block.Replace("{{ .Params }}"), `Name"${Name}",Num#${Num}`)
}

type PercentDefn struct {
	TypeName string    `_:"display.width"`
	Class    string    `_:"display"`
	_        text.Text `txt:"Width "`
	Width    percentinput.Percent
}

func TestPercentInput(t *testing.T) {
	block := CreateFromDefinition(&PercentDefn{})
	assert.Eq(t, block.Replace("{{ .Class }}"), "quando-display")
	assert.Eq(t, block.Replace("{{ .TypeName }}"), "display.width")
	assert.Eq(t, block.Replace("{{ .Widgets }}"), `Width <input data-quando-name='Width' type='number' min='0' max='100'/>%`)
	assert.Eq(t, block.Replace("{{ .Params }}"), `Width#${Width}`)
}
