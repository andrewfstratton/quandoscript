package percentinput

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestNumberEmpty(t *testing.T) { // n.b. should never happen
	tf := New("")
	assert.Eq(t, tf.Html(), `<input data-quando-name='' type='number' min='0' max='100'/>%`)
}

func TestTextFieldSimple(t *testing.T) {
	tf := New("percent")
	assert.Eq(t, tf.Html(), `<input data-quando-name='percent' type='number' min='0' max='100'/>%`)
	tf.Default(50).Empty("empty").Width(4)
	assert.Eq(t, tf.Html(), `<input data-quando-name='percent' type='number' value='50' placeholder='empty' style='width:4em' min='0' max='100'/>%`)
}

func TestScriptSimple(t *testing.T) {
	tf := New("name")
	assert.Eq(t, tf.Generate(), `name#${name}`)
}
