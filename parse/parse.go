package parse

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/andrewfstratton/quandoscript/op"
)

type Input struct {
	line string
	err  error
}

const (
	UNKNOWN int = iota
	VARIABLE
	BOOLEAN
	STRING
	NUMBER // may need range and integer
	ID
)

type Param struct {
	val   any
	qtype int
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

type Params map[string]Param

func Line(line string) (fn op.Op, err error) {
	input := Input{line: line}
	id := input.getId()
	if input.err != nil {
		err = input.err
		return
	}
	fmt.Printf("Found id :%v\n leaving :'%v'\n", id, line)
	fn = nil
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
func stripSpacer(input *Input) {
	_ = input.matchStart("[( )\t]+", "space/tab")
}

// removes and returns a word at start of input.line, or err if missing.
// word starts with a letter, then may also include digits . or _
func (input *Input) getWord() (word string) {
	return input.matchStart("[a-zA-Z][a-zA-Z0-9_.]*", "word starting a-z or A-Z")
}

// returns parameters as nil if just (), or Param parameters, err if not starting with ( or not terminated correctly with ).
// remaining is the rest of the string
func getParams(input *Input) (params Params) {
	found := input.matchStart(`\(`, "(")
	if found == "" {
		return
	}
	params = make(Params)
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
func getParam(input *Input) (key string, param Param) {
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
	prefix := input.matchStart("(!|:)", "type prefix/assignment missing ")
	switch prefix {
	case "":
		param.qtype = UNKNOWN
		break
	case "!": // check for boolean
		found = input.matchStart("(true|false)", "")
		if found != "" { // i.e. if found
			param.qtype = BOOLEAN
			param.val = (found == "true")
			return
		}
	case ":": // check for id
		id := input.getId()
		if input.err == nil {
			param.val = id
			param.qtype = ID
			return
		}
	case "=": // check for variable
		variable := input.getWord()
		if input.err == nil {
			param.val = variable
			param.qtype = VARIABLE
			return
		}
	}
	// error (or not handled, e.g. due ot mismatch with generator)
	key = "" // have to reset since already stored
	input.line = restore
	return
}
