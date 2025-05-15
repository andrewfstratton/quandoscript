package block

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block/widget/character"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
)

func TestNew(t *testing.T) {
	block := New("")
	assert.Eq(t, block, nil) // n.b. will panic when not testing

	block = New("system.log")
	block.Add(text.New("Log").Bold())
	block.Add(character.New(character.FIXED_SPACE))
	block.Add(stringinput.New("name").Empty("message"))
	assert.Neq(t, block.html(), "")
}
