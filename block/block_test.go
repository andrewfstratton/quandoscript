package block

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block/widget/character"
	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
	"github.com/andrewfstratton/quandoscript/block/widget/percentinput"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
)

func TestEmpty(t *testing.T) {
	block := New("", "")     // test with empty class here
	assert.Eq(t, block, nil) // n.b. will panic when not testing
}

func TestNewSimple(t *testing.T) {
	block := New("log", "")
	assert.Eq(t, block.class, "")
	assert.Eq(t, block.qid, "log")

	block = New("system.log", "sys")
	assert.Eq(t, block.class, "sys")
	assert.Eq(t, block.qid, "system.log")
}

func TestNew(t *testing.T) {
	block := New("system.log", "system")
	block.Add(text.New("Log"))
	block.Add(character.New(character.FIXED_SPACE))
	be := block.Expand()
	assert.Eq(t, be.Replace("{{ .Class }}"), "system")
	assert.Eq(t, be.Replace("{{ .QID }}"), "system.log")
	assert.Eq(t, be.Replace("{{ .Widgets }}"), "Log&nbsp;")
	assert.Eq(t, be.Replace("{{ .Params }}"), "")
}

func TestNewStringInput(t *testing.T) {
	block := New("system.log", "system")
	block.Add(text.New("Log ").Bold())
	block.Add(stringinput.New("name").Default("!").Empty("message"))
	be := block.Expand()
	assert.Eq(t, be.Replace("{{ .Params }}"), `name"${name}"`)
}

func TestNewNumberInput(t *testing.T) {
	block := New("system.log", "system")
	block.Add(text.New("Log "))
	block.Add(numberinput.New("num").Empty("message").Default(0.5).Min(0).Max(1).Width(4))
	be := block.Expand()
	assert.Eq(t, be.Replace("{{ .Params }}"), `num#${num}`)
	assert.Eq(t, be.Replace("{{ .Widgets }}"),
		`Log <input data-quando-name='num' type='number' value='0.5' placeholder='message' style='width:4em' min='0' max='1'/>`)
}

func TestNewPercentInput(t *testing.T) {
	block := New("display.width", "display")
	block.Add(text.New("Width "))
	block.Add(percentinput.New("width").Empty("0-100").Default(50))
	be := block.Expand()
	assert.Eq(t, be.Replace("{{ .Class }}"), "display")
	assert.Eq(t, be.Replace("{{ .QID }}"), "display.width")
	assert.Eq(t, be.Replace("{{ .Widgets }}"), `Width <input data-quando-name='width' type='number' value='50' placeholder='0-100' min='0' max='100'/>%`)
	assert.Eq(t, be.Replace("{{ .Params }}"), "width#${width}")
}
