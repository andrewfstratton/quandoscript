package action

import (
	"github.com/andrewfstratton/quandoscript/action/param"
)

type Action struct {
	op     func(param.Params)
	Params param.Params
	// context
}

type Op func(param.Params)
type OpOp func(param.Params) func(param.Params)

func New(o Op, late param.Params) *Action {
	return &Action{op: o, Params: late}
}

func (action *Action) Exec() {
	action.op(action.Params)
}

// lineid, word, params, err := parse.Line(`0 system.log(greeting"Hi",txt"Bob")`)
// fmt.Println(lineid, word, params, err)
