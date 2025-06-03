package action

import (
	"github.com/andrewfstratton/quandoscript/action/param"
)

type Action struct {
	op     func(param.Params)
	Params param.Params
	NextId int
	// context
}

var Actions map[int]*Action // map id to action
var last *Action
var startId int = -1

type Op func(param.Params)
type OpOp func(param.Params) func(param.Params)

func New(o Op, late param.Params) *Action {
	action := Action{op: o, Params: late, NextId: -1} // N.B. -1 is to show no following action
	return &action
}

func NewGroup() {
	last = nil // so we don't append to the same group
}

func Add(id int, action *Action) {
	if startId == -1 {
		startId = id
	}
	Actions[id] = action
	if last != nil {
		last.NextId = id
	}
	last = action
}

func Run(id int) {
	for id != -1 {
		// fmt.Print(strconv.Itoa(id) + "-")
		act := Actions[id]
		act.op(act.Params)
		id = act.NextId
	}
}

func Start() (msg string) {
	if startId == -1 {
		return "No actions found"
	}
	Run(startId)
	return
}

func init() {
	Actions = map[int]*Action{} // i.e. the lookup table to find any action
}
