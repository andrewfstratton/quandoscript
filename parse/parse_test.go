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
	match.stripSpacer()
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")

	match = Input{line: "word.word"}
	match.stripSpacer()
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "word.word")

	match = Input{line: " \t  w"}
	match.stripSpacer()
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
	params := match.getParams()
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.True(t, params == nil)

	match = Input{line: "()"}
	params = match.getParams()
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 0) //i.e. no parameters

	match = Input{line: "(x!true)"}
	params = match.getParams()
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 1)
	assert.Eq(t, params["x"], true)

	match = Input{line: `(a"hello!",b:12345,x=val,y!true,z#-12.34e56)`}
	params = match.getParams()
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 5)
	assert.Eq(t, params["a"], "hello!")
	assert.Eq(t, params["b"], 12345)
	assert.Eq(t, params["x"], "val")
	assert.Eq(t, params["y"], true)
	assert.Eq(t, params["z"], -12.34e56)
}

func TestParseParamBool(t *testing.T) {
	match := Input{line: ""}
	key, _ := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "")
	assert.Neq(t, match.err, nil)

	match = Input{line: `)`} // closing ) ends parameters, no error
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ")")
	assert.True(t, p == nil)
	assert.Eq(t, match.err, nil)

	match = Input{line: "!a"}
	key, _ = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "!a")
	assert.Neq(t, match.err, nil)

	match = Input{line: "a!"}
	key, p = getParam(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "a!")
	assert.Eq(t, key, "")
	assert.True(t, p == nil)

	match = Input{line: "x!true"}
	key, p = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, key, "x")
	assert.Eq(t, p, true)

	match = Input{line: "y!false,z!true"}
	key, p = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, ",z!true")
	assert.Eq(t, key, "y")
	assert.Eq(t, p, false)
}

func TestParseParamId(t *testing.T) {
	match := Input{line: "a:"}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a:")
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: ":a"}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ":a")
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x:1"}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, 1)

	match = Input{line: "y:99,x:12"}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x:12")
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, 99)
}

func TestParseParamVariable(t *testing.T) {
	match := Input{line: "a="}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a=")
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: "=a"}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "=a")
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x=y"}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, "y")

	match = Input{line: "y=V_a9,x=txt"}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x=txt")
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, "V_a9")
}

func TestParseParamString(t *testing.T) {
	match := Input{line: `a"`}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, `a"`)
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: `"a`}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, `"a`)
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: `x"y"`}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, "y")

	match = Input{line: `y"\\S\tt\nr\"",x"txt"`}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, `,x"txt"`)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, "\\S\tt\nr"+`"`)
}

func TestParseParamNumber(t *testing.T) {
	match := Input{line: "a#"}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a#")
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: "#a"}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "#a")
	assert.True(t, p == nil)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x#1"}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, 1.0)

	match = Input{line: "y#-0.99,x#12"}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x#12")
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p, -0.99)
}

func TestParseLine(t *testing.T) {
	t.Fail()
}

func TestParseLines(t *testing.T) {
	t.Fail()
}
