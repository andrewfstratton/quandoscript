package param

import (
	"fmt"
)

type Param struct {
	Val   any
	Qtype int
}

const (
	UNKNOWN int = iota
	VARIABLE
	BOOLEAN
	STRING
	NUMBER // may need range and integer
	LINEID
)

type Params map[string]Param
type Op func(Params) func(Params)

func (params Params) String(lookup string, stringp *string) {
	param, found := params[lookup]
	if found {
		if param.Qtype == STRING {
			val := param.Val.(string)
			stringp = &val
			return
		}
		// lookup variable here...
		fmt.Println("Error : ", lookup, " incorrect type")
	}
}
