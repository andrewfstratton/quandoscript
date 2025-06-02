package op

import (
	"github.com/andrewfstratton/quandoscript/run/param"
)

type Op struct {
	Op     func(param.Params)
	Params param.Params
	// context
}

type OpOp func(param.Params) func(param.Params)

func (op *Op) Exec() {
	op.Op(op.Params)
}

// lineid, word, params, err := parse.Line(`0 system.log(greeting"Hi",txt"Bob")`)
// fmt.Println(lineid, word, params, err)
