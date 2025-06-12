package stringinput

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/block/widget"
)

func TestTextFieldEmpty(t *testing.T) { // n.b. should never happen
	tf := New("")
	assert.Eq(t, tf.Html(), `&quot;<input data-quando-name='' type='text'/>&quot;`)
}

func TestTextFieldSimple(t *testing.T) {
	tf := New("name")
	widget.Setup(tf, "", `default:"default"`)
	widget.Setup(tf, "_", `empty:"empty"`)
	assert.Eq(t, tf.Html(), `&quot;<input data-quando-name='name' type='text' value='default' placeholder='empty'/>&quot;`)
}

func TestScriptSimple(t *testing.T) {
	tf := New("name")
	assert.Eq(t, tf.Generate(), `name"${name}"`)
}
