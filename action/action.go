package action

import (
	"github.com/andrewfstratton/quandoscript/action/param"
)

type Action struct {
	late          func(param.Params)
	params        param.Params
	firstChildren []int
	parent        int
	nextId        int // default 0 means none following
	// context
}

var actions map[int]*Action = map[int]*Action{} // map id to action
var last *Action
var startId int = 0

type (
	// outer/setup function that returns late inner function
	Early func(param.Params) func(param.Params) // e.g. used for menu options/constants chosen by the end user

	// inner function that is run every invocation
	Late func(param.Params) // e.g. used for variable substitution
)

func New(late Late, params param.Params, first_child_ids []int) *Action {
	action := Action{late: late, params: params, firstChildren: first_child_ids}
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
	setParents(startId)
	Run(startId)
	return
}

func setParents(parentId int) {
	parent, found := actions[parentId]
	if !found {
		return
	}
	for _, childId := range parent.firstChildren {
		child, found := actions[childId]
		if !found {
			return
		}
		child.parent = parentId
		setParents(childId)
	}
}
