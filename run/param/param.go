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
type StringParam struct {
	Lookup string
	Val    string
}

func (param *StringParam) Update(params Params) {
	p, found := params[param.Lookup]
	if found {
		if p.Qtype == STRING {
			param.Val = p.Val.(string)
			return
		}
		// lookup variable here...
		fmt.Println("Error : ", param.Lookup, " incorrect type")
	}
}
