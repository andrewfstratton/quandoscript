package action

import (
	"github.com/andrewfstratton/quandoscript/action/param"
)

type Action struct {
	op     func(param.Params)
	Params param.Params
	// context
}

var Actions map[int]*Action

type Op func(param.Params)
type OpOp func(param.Params) func(param.Params)

func New(o Op, late param.Params) *Action {
	return &Action{op: o, Params: late}
}

func (action *Action) Exec() {
	action.op(action.Params)
}

func init() {
	Actions = map[int]*Action{}
}
