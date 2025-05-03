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
	var id int
	input := Input{line: line}
	id = getId(&input)
	if err != nil {
		return nil, errors.New("Failed to find id at start of line:\n\t" + err.Error())
	}
	fmt.Printf("Found id :%v\n leaving :'%v'\n", id, line)
	return fn, err
}

// returns a [0..9]+ digit value at start of line, or err.  remaining is the rest of the string
func getId(input *Input) (id int) {
	re := regexp.MustCompile("^([0-9])+")
	arr := re.FindStringIndex(input.line)
	if len(arr) != 2 {
		input.err = errors.New("Failed to find digits at start of '" + input.line + "'")
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

// returns word at start of line, or err if missing.  remaining is the rest of the string
// word starts with a letter, then may also include digits . or _
func getWord(line string) (word string, remaining string, err error) {
	re := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_.]*")
	arr := re.FindStringIndex(line)
	if len(arr) != 2 {
		remaining = line
		err = errors.New("Failed to find a word starting with a-z or A-Z at start of '" + line + "'")
	} else {
		count := arr[1] // start must be 0 due to regexp starting ^
		remaining = line[count:]
		word = line[0:count]
	}
	return
}

// returns parameters as nil if just (), or Param parameters, err is true if not starting with ( or not terminated correctly with ).
// remaining is the rest of the string
func getParams(line string) (params Params, remaining string, err error) {
	re := regexp.MustCompile(`^\(`)
	arr := re.FindStringIndex(line)
	if len(arr) != 2 {
		remaining = line
		err = errors.New("Failed to find ( at start of '" + line + "'")
		return
	}
	params = make(Params)
	line = line[1:] // strip first character
	var key string
	var val bool
	key, val, remaining, err = getParam(line)
	if err == nil {
		if key != "" {
			params[key] = val
		}
	}
	line = remaining
	re = regexp.MustCompile(`^\)`)
	arr = re.FindStringIndex(line)
	if len(arr) != 2 {
		remaining = line
		err = errors.New("Failed to find ) at start of '" + line + "'")
		return
	}
	line = line[1:] // strip first character
	remaining = line
	return
}

// key returns "" when none found
func getParam(line string) (key string, val bool, remaining string, err error) {
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*=`)
	arr := re.FindStringIndex(line)
	if len(arr) != 2 {
		remaining = line
		return // ignore error since it could be )
	}
	count := arr[1] // start must be 0 due to regexp starting ^
	key = line[:count-1]
	line = line[count:]
	re = regexp.MustCompile(`^(true|false)`)
	arr = re.FindStringIndex(line)
	if len(arr) != 2 {
		remaining = line
		err = errors.New("Failed to find 'true' or 'false' at start of '" + line + "'")
		return
	}
	count = arr[1] // start must be 0 due to regexp starting ^
	if line[:count] == "true" {
		val = true
	}
	line = line[count:]
	remaining = line
	return
}
