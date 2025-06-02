package parse

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type Input struct {
	line string
	err  error
}

func (input *Input) matchStart(rxp string, lookfor string) (found string) {
	arr := regexp.MustCompile("^" + rxp).FindStringIndex(input.line)
	if len(arr) != 2 {
		input.err = errors.New("Failed to match " + lookfor + " with '" + rxp + "' at start of '" + input.line + "'")
		return
	}
	count := arr[1] // start must be 0 due to regexp starting ^
	found = input.line[:count]
	input.line = input.line[count:]
	return
}

func Line(line string) (lineid int, word string, params param.Params, err error) {
	if line == "" { // fn and err are nil for a blank line
		return
	}
	input := Input{line: line}
	lineid = input.getId()
	if input.err != nil {
		err = input.err
		return
	}
	input.stripSpacer()
	if input.err != nil {
		err = input.err
		return
	}
	word = input.getWord()
	if input.err != nil {
		err = input.err
		return
	}
	params = input.getParams()
	if input.err != nil {
		err = input.err
		return
	}
	return
}

// removes and returns a [0..9] integer from start of input.line, or input.err.
func (input *Input) getId() (id int) {
	found := input.matchStart("([0-9])+", "Id")
	if found == "" {
		return
	}
	var err error
	id, err = strconv.Atoi(found) // error must be nil
	if err != nil {
		fmt.Println("CODING ERROR IN parse:getId()")
		os.Exit(99)
	}
	return
}

// strips space/tab from start of input.line, or input.err if missing
func (input *Input) stripSpacer() {
	_ = input.matchStart("[( )\t]+", "space/tab")
}

// removes and returns a word at start of input.line, or err if missing.
// word starts with a letter, then may also include digits . or _
func (input *Input) getWord() string {
	return input.matchStart("[a-zA-Z][a-zA-Z0-9_.]*", "word starting a-z or A-Z")
}

// removes and returns a string" at start of input.line, or err if missing.
// the string may contain \\, \", \t and \n, which will be substituted
// N.B. string does NOT start with '"' - this will already have parsed
func (input *Input) getString() (str string) {
	for {
		if len(input.line) == 0 {
			input.err = errors.New(`string does not terminate with '"' before end of line`)
			break
		}
		ch := input.line[:1]
		input.line = input.line[1:] // consume first character
		var skip bool
		if ch == `"` {
			break
		}
		if ch == `\` && len(input.line) > 0 {
			ch2 := input.line[:1]
			switch ch2 {
			case `\`:
				ch = `\`
				skip = true
			case "t":
				ch = "\t"
				skip = true
			case "n":
				ch = "\n"
				skip = true
			case `"`:
				ch = `"`
				skip = true
			}
		}
		str += ch
		if skip {
			input.line = input.line[1:]
		}
	}
	return str
}

// removes and returns a decimal floating point number at start of input.line, or err if missing.
func (input *Input) getFloat() (f float64) {
	found := input.matchStart("[+-]?[0-9]+[.]?[0-9]*([eE][+-]?[0-9]+)?", "floating point number")
	if found == "" {
		return
	}
	var err error
	f, err = strconv.ParseFloat(found, 64) // error must be nil
	if err != nil {
		fmt.Println("CODING ERROR IN parse:getFloat()")
		os.Exit(99)
	}
	return
}

// returns parameters as nil if just (), or Param parameters, err if not starting with ( or not terminated correctly with ).
// remaining is the rest of the string
func (input *Input) getParams() (params param.Params) {
	found := input.matchStart(`\(`, "(")
	if found == "" {
		return
	}
	params = make(param.Params)
	for {
		key, param := getParam(input)
		if key == "" { // no key
			break
		}
		if input.err != nil {
			return
		}
		params[key] = param
		found = input.matchStart(`,`, "")
		if found != "," {
			input.err = nil // clear out err and drop out of loop
			break
		}
	}
	_ = input.matchStart(`\)`, ")") // we can return whether we found or not
	return
}

// key returns "" when none found
func getParam(input *Input) (key string, p param.Param) {
	restore := input.line
	// Check for ) and return without error or change to input if found
	found := input.matchStart(`\)`, "")
	input.err = nil   // supress error
	if found == `)` { // found ) so reset
		input.line = restore
		input.err = nil
		return
	}
	key = input.getWord()
	if key == "" {
		return
	}
	// check for valid prefix
	prefix := input.matchStart(`[!:="#]`, "type prefix/assignment missing ")
	switch prefix {
	case "":
		p.Qtype = param.UNKNOWN
	case "!": // check for boolean
		found = input.matchStart("(true|false)", "")
		if found != "" { // i.e. if found
			p.Qtype = param.BOOLEAN
			p.Val = (found == "true")
			return
		}
	case ":": // check for lineid
		lineid := input.getId()
		if input.err == nil {
			p.Val = lineid
			p.Qtype = param.LINEID
			return
		}
	case "=": // check for variable
		name := input.getWord()
		if input.err == nil {
			p.Val = name
			p.Qtype = param.VARIABLE
			return
		}
	case `"`: // check for string
		str := input.getString()
		if input.err == nil {
			p.Val = str
			p.Qtype = param.STRING
			return
		}
	case "#": // check for float
		num := input.getFloat()
		if input.err == nil {
			p.Val = num
			p.Qtype = param.NUMBER
			return
		}
	}
	// error (or not handled, e.g. due to mismatch with generator)
	key = "" // have to reset since already stored
	input.line = restore
	return
}
