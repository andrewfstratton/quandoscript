package parse

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/andrewfstratton/quandoscript/op"
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

type Params map[string]any

func Line(line string) (fn op.Op, err error) {
	input := Input{line: line}
	id := getId(&input)
	if input.err != nil {
		err = input.err
		return
	}
	fmt.Printf("Found id :%v\n leaving :'%v'\n", id, line)
	fn = nil
	return
}

// removes and returns a [0..9] integer from start of input.line, or input.err.
func getId(input *Input) (id int) {
	found := input.matchStart("([0-9])+", "Id")
	id, _ = strconv.Atoi(found) // error must be nil
	return
}

// strips space/tab from start of input.line, or input.err if missing
func stripSpacer(input *Input) {
	_ = input.matchStart("[( )\t]+", "space/tab")
}

// removes and returns a word at start of input.line, or err if missing.
// word starts with a letter, then may also include digits . or _
func getWord(input *Input) (word string) {
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
		key, val := getParam(input)
		if key == "" { // no key
			break
		}
		if input.err != nil {
			return
		}
		params[key] = val
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
func getParam(input *Input) (key string, val bool) {
	restore := input.line
	// Check for ) and return without error or change to input if found
	found := input.matchStart(`\)`, "")
	input.err = nil   // supress error
	if found == `)` { // found ) so return straight away, key and val == ""
		input.line = restore
		return
	}
	key = input.matchStart(`[a-zA-Z][a-zA-Z0-9_.]*=`, "word starting a-z or A-Z")
	if key == "" {
		return
	}
	key = key[:len(key)-1] // removes = at end
	found = input.matchStart("(true|false)", "")
	if found == "" {
		key = "" // have to reset since already stored
		input.line = restore
		return
	}
	if found == "true" {
		val = true
	}
	return
}
