package action

import (
	"github.com/andrewfstratton/quandoscript/action/param"
)

type Action struct {
	late   func(param.Params)
	params param.Params
	nextId int
	// context
}

var actions map[int]*Action // map id to action
var last *Action
var startId int = 0

type (
	// outer/setup function that returns late inner function
	Early func(param.Params) func(param.Params) // e.g. used for menu options/constants chosen by the end user

	// inner function that is run every invocation
	Late func(param.Params) // e.g. used for variable substitution
)

func New(late Late, params param.Params) *Action {
	action := Action{late: late, params: params, nextId: 0} // N.B. 0 is to show no following action
	return &action
}

func NewGroup() {
	last = nil // so we don't append to the same group
}

func Add(id int, action *Action) {
	if startId == 0 {
		startId = id
	}
	actions[id] = action
	if last != nil {
		last.nextId = id
	}
	last = action
}

func Run(id int) {
	for id != 0 {
		act := actions[id]
		act.late(act.params)
		id = act.nextId
	}
}

func Start() (warn string) {
	if startId == 0 {
		return "No actions found"
	}
	Run(startId)
	return
}

func init() {
	actions = map[int]*Action{} // i.e. the table to find action on id
}
