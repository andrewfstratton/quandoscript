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

// removes and returns a [0..9] integer from start of input.line, or err.
func getId(input *Input) (id int) {
	re := regexp.MustCompile("^([0-9])+")
	arr := re.FindStringIndex(input.line)
	if len(arr) != 2 {
		input.err = errors.New("Failed to find Id as digits at start of '" + input.line + "'")
		return
	}
	count := arr[1]                          // start must be 0 due to regexp starting ^
	id, _ = strconv.Atoi(input.line[:count]) // err must be nil
	input.line = input.line[count:]
	return
}

// strips space/tab from start of input.line, or err if missing
func stripSpacer(input *Input) {
	re := regexp.MustCompile("^[( )\t]+")
	arr := re.FindStringIndex(input.line)
	if len(arr) != 2 {
		input.err = errors.New("Failed to find space/tab at start of '" + input.line + "'")
		return
	}
	count := arr[1] // start must be 0 due to regexp starting ^
	input.line = input.line[count:]
}

// removes and returns a word at start of input.line, or err if missing.
// word starts with a letter, then may also include digits . or _
func getWord(input *Input) (word string) {
	re := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_.]*")
	arr := re.FindStringIndex(input.line)
	if len(arr) != 2 {
		input.err = errors.New("Failed to find a word starting with a-z or A-Z at start of '" + input.line + "'")
		return
	}
	count := arr[1] // start must be 0 due to regexp starting ^
	word = input.line[:count]
	input.line = input.line[count:]
	return
}

// returns parameters as nil if just (), or Param parameters, err if not starting with ( or not terminated correctly with ).
// remaining is the rest of the string
func getParams(input *Input) (params Params) {
	re := regexp.MustCompile(`^\(`)
	arr := re.FindStringIndex(input.line)
	if len(arr) != 2 {
		input.err = errors.New("Failed to find ( at start of '" + input.line + "'")
		return
	}
	params = make(Params)
	input.line = input.line[1:] // strip first character
	for {
		key, val := getParam(input)
		if key == "" { // no of keys
			break
		}
		if input.err != nil {
			return
		}
		params[key] = val
		re := regexp.MustCompile(`^,`)
		arr := re.FindStringIndex(input.line)
		if len(arr) != 2 {
			break
		}
		input.line = input.line[1:]
	}
	re = regexp.MustCompile(`^\)`)
	arr = re.FindStringIndex(input.line)
	if len(arr) != 2 {
		input.err = errors.New("Failed to find ) at start of '" + input.line + "'")
		return
	}
	input.line = input.line[1:] // strip first character
	return
}

// key returns "" when none found
func getParam(input *Input) (key string, val bool) {
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*=`)
	arr := re.FindStringIndex(input.line)
	if len(arr) != 2 {
		return // ignore error since it could be )
	}
	count := arr[1]                 // start must be 0 due to regexp starting ^
	key = input.line[:count-1]      // -1 to drop =
	remaining := input.line[count:] // TODO check about parsing when not boolean - fail?!
	re = regexp.MustCompile(`^(true|false)`)
	arr = re.FindStringIndex(remaining)
	if len(arr) != 2 {
		key = "" // have to reset since already stored
		input.err = errors.New("Failed to find 'true' or 'false' at start of '" + remaining + "'")
		return
	}
	count = arr[1] // start must be 0 due to regexp starting ^
	if remaining[:count] == "true" {
		val = true
	}
	input.line = remaining[count:]
	return
}
