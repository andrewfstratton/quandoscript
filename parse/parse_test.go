package parse

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestParseId(t *testing.T) {
	// test valid id at start of string
	id, remaining, err := getId("90 ignore")
	assert.Eq(t, id, 90)
	assert.Eq(t, remaining, " ignore")
	assert.Eq(t, err, nil)

	// test empty string
	id, remaining, err = getId("")
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, "")

	// test missing id in function call
	match := ",key=false)"
	id, remaining, err = getId(match)
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, match)
}

func TestParseSpacer(t *testing.T) {
	match := ""
	remaining, err := stripSpacer(match)
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, match)

	match = "word.word"
	remaining, err = stripSpacer(match)
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, match)

	match = " \t  w"
	remaining, err = stripSpacer(match)
	assert.Eq(t, err, nil)
	assert.Eq(t, remaining, "w")

}

func TestParseWord(t *testing.T) {
	match := ""
	word, remaining, err := getWord(match)
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, match)
	assert.Eq(t, word, "")

	match = "w"
	word, remaining, err = getWord(match)
	assert.Eq(t, err, nil)
	assert.Eq(t, remaining, "")
	assert.Eq(t, word, match)

	match = "word.word()"
	word, remaining, err = getWord(match)
	assert.Eq(t, err, nil)
	assert.Eq(t, remaining, "()")
	assert.Eq(t, word, "word.word")
}
