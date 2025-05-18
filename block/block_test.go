package block

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block/widget/character"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
)

func TestEmpty(t *testing.T) {
	block := New("")
	assert.Eq(t, block, nil) // n.b. will panic when not testing
}

func TestSimple(t *testing.T) {
	block := New("system.log")
	assert.Eq(t, block.Class(), "system")
}

func TestNew(t *testing.T) {
	block := New("system.log")
	block.Add(text.New("Log").Bold())
	block.Add(character.New(character.FIXED_SPACE))
	block.Add(stringinput.New("name").Empty("message"))
	assert.Neq(t, block.html(), "")
	assert.Eq(t, block.script(), `name"${name}"`)
	block.Add(numberinput.New("val"))
	assert.Eq(t, block.script(), `name"${name}",val#${val}`)
}
