package textfield

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestTextFieldEmpty(t *testing.T) {
	tf := New("")
	assert.Eq(t, tf.Html(), `<input data-quando-name='' type='text' data-quando-encode='normal'/>`)
}

func TestTextFieldSimple(t *testing.T) {
	tf := New("name")
	tf.Default("default")
	tf.Empty("empty")
	assert.Eq(t, tf.Html(), `<input data-quando-name='name' type='text' value='default' placeholder='empty' data-quando-encode='normal'/>`)
}
