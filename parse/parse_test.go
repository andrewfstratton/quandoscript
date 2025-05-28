package parse

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/run/param"
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
	assert.Eq(t, params["x"].Val, true)

	match = Input{line: `(a"hello!",b:12345,x=val,y!true,z#-12.34e56)`}
	params = match.getParams()
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 5)
	assert.Eq(t, params["a"].Qtype, param.STRING)
	assert.Eq(t, params["a"].Val, "hello!")
	assert.Eq(t, params["b"].Qtype, param.LINEID)
	assert.Eq(t, params["b"].Val, 12345)
	assert.Eq(t, params["x"].Qtype, param.VARIABLE)
	assert.Eq(t, params["x"].Val, "val")
	assert.Eq(t, params["y"].Qtype, param.BOOLEAN)
	assert.Eq(t, params["y"].Val, true)
	assert.Eq(t, params["z"].Qtype, param.NUMBER)
	assert.Eq(t, params["z"].Val, -12.34e56)
}

func TestParseParamBool(t *testing.T) {
	match := Input{line: ""}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: `)`} // closing ) ends parameters, no error
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ")")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Eq(t, match.err, nil)

	match = Input{line: "!a"}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "!a")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "a!"}
	key, p = getParam(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "a!")
	assert.Eq(t, key, "")
	assert.Eq(t, p.Qtype, param.UNKNOWN)

	match = Input{line: "x!true"}
	key, p = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, key, "x")
	assert.Eq(t, p.Qtype, param.BOOLEAN)
	assert.Eq(t, p.Val, true)

	match = Input{line: "y!false,z!true"}
	key, p = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, ",z!true")
	assert.Eq(t, key, "y")
	assert.Eq(t, p.Qtype, param.BOOLEAN)
	assert.Eq(t, p.Val, false)
}

func TestParseParamId(t *testing.T) {
	match := Input{line: "a:"}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a:")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: ":a"}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ":a")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x:1"}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, p.Qtype, param.LINEID)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, 1)

	match = Input{line: "y:99,x:12"}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x:12")
	assert.Eq(t, p.Qtype, param.LINEID)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, 99)
}

func TestParseParamVariable(t *testing.T) {
	match := Input{line: "a="}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a=")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "=a"}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "=a")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x=y"}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, p.Qtype, param.VARIABLE)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, "y")

	match = Input{line: "y=V_a9,x=txt"}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x=txt")
	assert.Eq(t, p.Qtype, param.VARIABLE)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, "V_a9")
}

func TestParseParamString(t *testing.T) {
	match := Input{line: `a"`}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, `a"`)
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: `"a`}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, `"a`)
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: `x"y"`}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, p.Qtype, param.STRING)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, "y")

	match = Input{line: `y"\\S\tt\nr\"",x"txt"`}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, `,x"txt"`)
	assert.Eq(t, p.Qtype, param.STRING)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, "\\S\tt\nr"+`"`)
}

func TestParseParamNumber(t *testing.T) {
	match := Input{line: "a#"}
	key, p := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a#")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "#a"}
	key, p = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "#a")
	assert.Eq(t, p.Qtype, param.UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x#1"}
	key, p = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, p.Qtype, param.NUMBER)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, 1.0)

	match = Input{line: "y#-0.99,x#12"}
	key, p = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x#12")
	assert.Eq(t, p.Qtype, param.NUMBER)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, p.Val, -0.99)
}
