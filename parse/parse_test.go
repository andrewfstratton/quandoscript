package parse

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestParseId(t *testing.T) {
	match := Input{line: ""}
	id := match.getId()
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")

	match = Input{line: ",key=false)"}
	id = match.getId()
	assert.Eq(t, id, 0) // id must be 1+
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, ",key=false)")

	match = Input{line: "90 ignore"}
	id = match.getId()
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
	word := match.getWord()
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, word, "")

	match = Input{line: "w"}
	word = match.getWord()
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, word, "w")

	match = Input{line: "word.word()"}
	word = match.getWord()
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "()")
	assert.Eq(t, word, "word.word")
}

func TestParseParams(t *testing.T) {
	match := Input{line: ""}
	params := getParams(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.True(t, params == nil)

	match = Input{line: "()"}
	params = getParams(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 0) //i.e. no parameters

	match = Input{line: "(x!true)"}
	params = getParams(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 1)
	assert.Eq(t, params["x"].val, true)

	match = Input{line: "(y!false,z!true)"}
	params = getParams(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 2)
	assert.Eq(t, params["y"].val, false)
	assert.Eq(t, params["z"].val, true)
}

func TestParseParamBool(t *testing.T) {
	match := Input{line: ""}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: `)`} // closing ) ends parameters, no error
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ")")
	assert.Eq(t, param.qtype, UNKNOWN)
	assert.Eq(t, match.err, nil)

	match = Input{line: "!a"}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "!a")
	assert.Eq(t, param.qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "a!"}
	key, param = getParam(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "a!")
	assert.Eq(t, key, "")
	assert.Eq(t, param.qtype, UNKNOWN)

	match = Input{line: "x!true"}
	key, param = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, key, "x")
	assert.Eq(t, param.qtype, BOOLEAN)
	assert.Eq(t, param.val, true)

	match = Input{line: "y!false,z!true"}
	key, param = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, ",z!true")
	assert.Eq(t, key, "y")
	assert.Eq(t, param.qtype, BOOLEAN)
	assert.Eq(t, param.val, false)
}

func TestParseParamId(t *testing.T) {
	match := Input{line: "a:"}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a:")
	assert.Eq(t, param.qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: ":a"}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ":a")
	assert.Eq(t, param.qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x:1"}
	key, param = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.qtype, ID)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.val, 1)

	match = Input{line: "y:99,x:12"}
	key, param = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x:12")
	assert.Eq(t, param.qtype, ID)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.val, 99)
}
func TestParseParamVariable(t *testing.T) {
	match := Input{line: "a="}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a=")
	assert.Eq(t, param.qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "=a"}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "=a")
	assert.Eq(t, param.qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x=y"}
	key, param = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.qtype, VARIABLE)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.val, "y")

	match = Input{line: "y=V_a9,x=txt"}
	key, param = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x=txt")
	assert.Eq(t, param.qtype, VARIABLE)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.val, "V_a9")
}
