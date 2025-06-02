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

func (param *NumberParam) Update(params Params) {
	p, found := params[param.Lookup]
	if found {
		if p.Qtype == NUMBER {
			param.Val = p.Val.(float64)
			return
		}
		// lookup variable here...
		fmt.Println("Error : ", param.Lookup, " incorrect type")
	}
}

func (param *IdParam) Update(params Params) {
	p, found := params[param.Lookup]
	if found {
		if p.Qtype == LINEID {
			param.Val = p.Val.(int)
			return
		}
		// lookup variable here...
		fmt.Println("Error : ", param.Lookup, " incorrect type")
	}
}
