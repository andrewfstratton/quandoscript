package stringinput

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestTextFieldEmpty(t *testing.T) {
	tf := New("")
	assert.Eq(t, tf.Html(), `&quot;<input data-quando-name='' type='text'/>&quot;`)
}

func TestTextFieldSimple(t *testing.T) {
	tf := New("name")
	tf.Default("default")
	tf.Empty("empty")
	assert.Eq(t, tf.Html(), `&quot;<input data-quando-name='name' type='text' value='default' placeholder='empty'/>&quot;`)
}
