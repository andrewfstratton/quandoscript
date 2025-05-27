package numberinput

import (
	"testing"

	"quando/quandoscript/assert"
)

func TestNumberEmpty(t *testing.T) { // n.b. should never happen
	tf := New("")
	assert.Eq(t, tf.Html(), `<input data-quando-name='' type='number'/>`)
}

func TestTextFieldSimple(t *testing.T) {
	tf := New("name")
	tf.Default(10)
	tf.Empty("empty")
	assert.Eq(t, tf.Html(), `<input data-quando-name='name' type='number' value='10' placeholder='empty'/>`)
	tf.Width(4).Min(0).Max(100)
	assert.Eq(t, tf.Html(), `<input data-quando-name='name' type='number' value='10' placeholder='empty' style='width:4em' min='0' max='100'/>`)
}

func TestScriptSimple(t *testing.T) {
	tf := New("name")
	assert.Eq(t, tf.Generate(), `name#${name}`)
}
