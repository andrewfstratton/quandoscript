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
type StringParam struct {
	Lookup string
	Val    string
}

type NumberParam struct {
	Lookup string
	Val    float64
}

type IdParam struct {
	Lookup string
	Val    int
}

func NewString(lookup string, val string, params Params) (param *StringParam) {
	param = &StringParam{Lookup: lookup, Val: val}
	param.Update(params)
	return
}

func (param *StringParam) Update(params Params) {
	p, found := params[param.Lookup]
	if !found {
		return
	}
	if p.Qtype == STRING {
		param.Val = p.Val.(string)
		return
	}
	// lookup variable here...
	fmt.Println("Error : ", param.Lookup, " incorrect type")
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
	if p.Qtype == NUMBER {
		param.Val = p.Val.(float64)
		return
	}
	// lookup variable here...
	fmt.Println("Error : ", param.Lookup, " incorrect type")
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
	if p.Qtype == LINEID {
		param.Val = p.Val.(int)
		return
	}
	// lookup variable here...
	fmt.Println("Error : ", param.Lookup, " incorrect type")
}
