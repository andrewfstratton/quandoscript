package parse

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestParseId(t *testing.T) {
	// test valid id at start of string
	id, remaining, err := GetId("90 ignore")
	assert.Eq(t, id, 90)
	assert.Eq(t, remaining, " ignore")
	assert.Eq(t, err, nil)

	// test empty string
	id, remaining, err = GetId("")
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, "")

	// test missing id in function call
	match := ",key=false)"
	id, remaining, err = GetId(match)
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, match)
}
