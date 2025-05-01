package parse

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/andrewfstratton/quandoscript/op"
)

func Line(line string) (fn op.Op, err error) {
	var id int
	id, line, err = GetId(line)
	if err != nil {
		return nil, errors.New("Failed to find id at start of line:\n\t" + err.Error())
	}
	fmt.Printf("Found id :%v\n leaving :'%v'\n", id, line)
	return fn, err
}

// returns a [0..9]+ digit value at start of line, or err.  remaining is the rest of the string
func GetId(line string) (id int, remaining string, err error) {
	re := regexp.MustCompile("^([0-9])+")
	arr := re.FindStringIndex(line)
	if len(arr) != 2 {
		remaining = line
		err = errors.New("Failed to find digits at start of '" + line + "'")
	} else {
		count := arr[1]                       // start must be 0 due to regexp starting ^
		id, err = strconv.Atoi(line[0:count]) // err should always be nil
		remaining = line[count:]
	}
	return
}
