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
	assert.Eq(t, params["x"].Val, true)

	match = Input{line: `(a"hello!",b:12345,x=val,y!true,z#-12.34e56)`}
	params = getParams(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, len(params), 5)
	assert.Eq(t, params["a"].Qtype, STRING)
	assert.Eq(t, params["a"].Val, "hello!")
	assert.Eq(t, params["b"].Qtype, ID)
	assert.Eq(t, params["b"].Val, 12345)
	assert.Eq(t, params["x"].Qtype, VARIABLE)
	assert.Eq(t, params["x"].Val, "val")
	assert.Eq(t, params["y"].Qtype, BOOLEAN)
	assert.Eq(t, params["y"].Val, true)
	assert.Eq(t, params["z"].Qtype, NUMBER)
	assert.Eq(t, params["z"].Val, -12.34e56)
}

func TestParseParamBool(t *testing.T) {
	match := Input{line: ""}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: `)`} // closing ) ends parameters, no error
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ")")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Eq(t, match.err, nil)

	match = Input{line: "!a"}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "!a")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "a!"}
	key, param = getParam(&match)
	assert.Neq(t, match.err, nil)
	assert.Eq(t, match.line, "a!")
	assert.Eq(t, key, "")
	assert.Eq(t, param.Qtype, UNKNOWN)

	match = Input{line: "x!true"}
	key, param = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, "")
	assert.Eq(t, key, "x")
	assert.Eq(t, param.Qtype, BOOLEAN)
	assert.Eq(t, param.Val, true)

	match = Input{line: "y!false,z!true"}
	key, param = getParam(&match)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, match.line, ",z!true")
	assert.Eq(t, key, "y")
	assert.Eq(t, param.Qtype, BOOLEAN)
	assert.Eq(t, param.Val, false)
}

func TestParseParamId(t *testing.T) {
	match := Input{line: "a:"}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a:")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: ":a"}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, ":a")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x:1"}
	key, param = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.Qtype, ID)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, 1)

	match = Input{line: "y:99,x:12"}
	key, param = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x:12")
	assert.Eq(t, param.Qtype, ID)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, 99)
}

func TestParseParamVariable(t *testing.T) {
	match := Input{line: "a="}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a=")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "=a"}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "=a")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x=y"}
	key, param = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.Qtype, VARIABLE)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, "y")

	match = Input{line: "y=V_a9,x=txt"}
	key, param = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x=txt")
	assert.Eq(t, param.Qtype, VARIABLE)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, "V_a9")
}

func TestParseParamString(t *testing.T) {
	match := Input{line: `a"`}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, `a"`)
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: `"a`}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, `"a`)
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: `x"y"`}
	key, param = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.Qtype, STRING)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, "y")

	match = Input{line: `y"\\S\tt\nr\"",x"txt"`}
	key, param = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, `,x"txt"`)
	assert.Eq(t, param.Qtype, STRING)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, "\\S\tt\nr"+`"`)
}

func TestParseParamNumber(t *testing.T) {
	match := Input{line: "a#"}
	key, param := getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "a#")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "#a"}
	key, param = getParam(&match)
	assert.Eq(t, key, "")
	assert.Eq(t, match.line, "#a")
	assert.Eq(t, param.Qtype, UNKNOWN)
	assert.Neq(t, match.err, nil)

	match = Input{line: "x#1"}
	key, param = getParam(&match)
	assert.Eq(t, key, "x")
	assert.Eq(t, match.line, "")
	assert.Eq(t, param.Qtype, NUMBER)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, 1.0)

	match = Input{line: "y#-0.99,x#12"}
	key, param = getParam(&match)
	assert.Eq(t, key, "y")
	assert.Eq(t, match.line, ",x#12")
	assert.Eq(t, param.Qtype, NUMBER)
	assert.Eq(t, match.err, nil)
	assert.Eq(t, param.Val, -0.99)
}
