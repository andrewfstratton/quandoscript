package parse

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestParseId(t *testing.T) {
	match := Input{line: ""}
	id := getId(&match)
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")

	match = Input{line: ",key=false)"}
	id = getId(&match)
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, ",key=false)")

	match = Input{line: "90 ignore"}
	id = getId(&match)
	assert.Eq(t, id, 90)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, " ignore")
}

func TestParseSpacer(t *testing.T) {
	match := Input{line: ""}
	stripSpacer(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")

	match = Input{line: "word.word"}
	stripSpacer(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "word.word")

	match = Input{line: " \t  w"}
	stripSpacer(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "w")

}

func TestParseWord(t *testing.T) {
	match := Input{line: ""}
	word := getWord(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, word, "")

	match = Input{line: "w"}
	word = getWord(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, word, "w")

	match = Input{line: "word.word()"}
	word = getWord(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "()")
	assert.Eq(t, word, "word.word")
}

func TestParseEmptyParam(t *testing.T) {
	match := ""
	params, remaining, err := getParams(match)
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, match)
	assert.Eq(t, len(params), 0)

	match = "()"
	params, remaining, err = getParams(match)
	assert.Eq(t, err, nil)
	assert.Eq(t, remaining, "")
	assert.Eq(t, len(params), 0) //i.e. no parameters

	match = "word.word()"
	params, remaining, err = getParams(match)
	assert.Neq(t, err, nil)
	assert.Eq(t, remaining, match)
	assert.Eq(t, len(params), 0)
}

func TestParseParamBool(t *testing.T) {
	match := "(x=true)"
	params, remaining, err := getParams(match)
	assert.Eq(t, err, nil)
	assert.Eq(t, remaining, "")
	assert.Eq(t, len(params), 1)
	assert.Eq(t, params["x"], true)

	match = "(y=false)"
	params, remaining, err = getParams(match)
	assert.Eq(t, err, nil)
	assert.Eq(t, remaining, "")
	assert.Eq(t, len(params), 1)
	assert.Eq(t, params["y"], false)
}
