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

func TestSimple(t *testing.T) {
	block := New("system.log", "")
	assert.Eq(t, block.class, "system")
}

func TestNew(t *testing.T) {
	block := New("system.log", "")
	block.Add(text.New("Log"))
	block.Add(character.New(character.FIXED_SPACE))
	out := block.Output()
	assert.Eq(t, out.qid, "system.log")
	assert.Eq(t, out.class, "")
	assert.Eq(t, out.params, "")
	assert.Eq(t, out.widgetHtml, "Log&nbsp;")
}

func TestNewStringInput(t *testing.T) {
	block := New("system.log", "system")
	block.Add(text.New("Log ").Bold())
	block.Add(stringinput.New("name").Default("!").Empty("message"))
	out := block.Output()
	assert.Eq(t, out.qid, "system.log")
	assert.Eq(t, out.class, "system")
	assert.Eq(t, out.params, `name"${name}"`)
	assert.Eq(t, out.widgetHtml, `<b>Log </b>&quot;<input data-quando-name='name' type='text' value='!' placeholder='message'/>&quot;`)
}

func TestNewNumberInput(t *testing.T) {
	block := New("system.log", "system")
	block.Add(text.New("Log "))
	block.Add(numberinput.New("name").Empty("message").Default(0.5).Min(0).Max(1).Width(4))
	out := block.Output()
	assert.Eq(t, out.qid, "system.log")
	assert.Eq(t, out.class, "system")
	assert.Eq(t, out.params, `name#${name}`)
	assert.Eq(t, out.widgetHtml, `Log <input data-quando-name='name' type='number' value='0.5' placeholder='message' style='width:4em' min='0' max='1'/>`)
}

func TestNewPercentInput(t *testing.T) {
	block := New("display.width", "display")
	block.Add(text.New("Width "))
	block.Add(percentinput.New("width").Empty("0-100").Default(50))
	out := block.Output()
	assert.Eq(t, out.qid, "display.width")
	assert.Eq(t, out.class, "display")
	assert.Eq(t, out.params, `width#${width}`)
	assert.Eq(t, out.widgetHtml, `Width <input data-quando-name='width' type='number' value='50' placeholder='0-100' min='0' max='100'/>%`)
}
