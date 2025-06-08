package param

import (
	"fmt"
)

type (
	BOOLEAN  = bool
	STRING   = string
	NUMBER   = float64
	LINEID   = int
	UNKNOWN  = struct{}
	VARIABLE string // n.b. is not an alias which may cause extra code
)

type Param any // horrible but easiest - must be currently
type Params map[string]Param
type StringParam struct {
	Lookup string
	Val    STRING
}

type NumberParam struct {
	Lookup string
	Val    NUMBER
}

type IdParam struct {
	Lookup string
	Val    LINEID
}

func NewString(lookup string, val string, params Params) (param *StringParam) {
	param = &StringParam{Lookup: lookup, Val: val}
	param.Update(params)
	return
}

func (param *StringParam) Update(params Params) {
	p, found := params[param.Lookup]
	if !found {
		return // nothing to update
	}
	switch s := p.(type) {
	case STRING:
		param.Val = s
	case VARIABLE:
		// lookup variable here...
	default:
		fmt.Println("Error : ", param.Lookup, " incorrect type")
	}
}

func NewNumber(lookup string, val float64, params Params) (param *NumberParam) {
	param = &NumberParam{Lookup: lookup, Val: val}
	param.Update(params)
	return
}

func (param *NumberParam) Update(params Params) {
	p, found := params[param.Lookup]
	if !found {
		return
	}
	switch n := p.(type) {
	case NUMBER:
		param.Val = n
	case VARIABLE:
		// lookup variable here...
	default:
		fmt.Println("Error : ", param.Lookup, " incorrect type")
	}
}

func (param *NumberParam) Int() int {
	return int(param.Val)
}

func NewId(lookup string, val int, params Params) (param *IdParam) {
	param = &IdParam{Lookup: lookup, Val: val}
	param.Update(params)
	return
}

func (param *IdParam) Update(params Params) {
	p, found := params[param.Lookup]
	if !found {
		return
	}
	switch l := p.(type) {
	case LINEID:
		param.Val = l
	case VARIABLE:
		// lookup variable here...
	default:
		fmt.Println("Error : ", param.Lookup, " incorrect type")
	}
}
