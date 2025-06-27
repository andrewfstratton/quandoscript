package param

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/definition"
)

type Param any // horrible but easiest - must be currently
type Params map[string]Param
type StringParam struct {
	Name string
	Val  definition.STRING
}

type NumberParam struct {
	Name string
	Val  definition.NUMBER
}

type IdParam struct {
	Name string
	Val  definition.LINEID
}

func NewString(name string, val string, params Params) (param *StringParam) {
	param = &StringParam{Name: name, Val: val}
	param.Update(params)
	return
}

func (param *StringParam) Update(params Params) {
	p, found := params[param.Name]
	if !found {
		return // nothing to update
	}
	switch s := p.(type) {
	case definition.STRING:
		param.Val = s
	case definition.VARIABLE:
		// lookup variable here...
	default:
		fmt.Println("Error : ", param.Name, " incorrect type")
	}
}

func NewNumber(name string, val float64, params Params) (param *NumberParam) {
	param = &NumberParam{Name: name, Val: val}
	param.Update(params)
	return
}

func (param *NumberParam) Update(params Params) {
	p, found := params[param.Name]
	if !found {
		return
	}
	switch n := p.(type) {
	case definition.NUMBER:
		param.Val = n
	case definition.VARIABLE:
		// lookup variable here...
	default:
		fmt.Println("Error : ", param.Name, " incorrect type")
	}
}

func (param *NumberParam) Int() int {
	return int(param.Val)
}

func (param *NumberParam) Bool() bool {
	return param.Val >= 0.5
}

func (param *NumberParam) Duration() time.Duration {
	// N.B. below is to allow for sub second times...
	return time.Duration(param.Val * float64(time.Second))
}

func NewId(name string, val int, params Params) (param *IdParam) {
	param = &IdParam{Name: name, Val: val}
	param.Update(params)
	return
}

func (param *IdParam) Update(params Params) {
	p, found := params[param.Name]
	if !found {
		return
	}
	switch l := p.(type) {
	case definition.LINEID:
		param.Val = l
	case definition.VARIABLE:
		// lookup variable here...
	default:
		fmt.Println("Error : ", param.Name, " incorrect type")
	}
}
