package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
	"github.com/andrewfstratton/quandoscript/library"
	"github.com/andrewfstratton/quandoscript/parse"
)

func init_log() {
	greeting := stringinput.New("greeting").Empty("greeting")
	txt := stringinput.New("txt").Empty("greeting")

	library.NewBlockType("system.log", "misc", logop).Add(
		text.New("Log "),
		greeting,
		text.New(" "),
		txt)
}

func logop(outer param.Params) func(param.Params) {
	greeting := param.StringParam{Lookup: "greeting", Val: ""}
	greeting.Update(outer)
	txt := param.StringParam{Lookup: "txt", Val: ""}
	txt.Update(outer)
	return func(inner param.Params) {
		txt.Update(inner)
		greeting.Update(inner)
		now := time.Now()
		fmt.Println("Log:", greeting.Val, txt.Val, now.Format(time.TimeOnly))
	}
}

func init() {
	init_log()
}

func main() {
	lineid, word, params, err := parse.Line(`0 system.log(greeting"Hi",txt"Bob")`)
	fmt.Println(lineid, word, params, err)
	o := library.NewAction(word, params, nil)

	action.Actions[lineid] = o

	lineid, word, params, err = parse.Line(`1 system.log(greeting"Hello",txt"Jane")`)
	fmt.Println(lineid, word, params, err)
	o = library.NewAction(word, params, nil)

	action.Actions[lineid] = o

	for l, act := range action.Actions {
		fmt.Println("<" + strconv.Itoa(l) + ">")
		act.Exec()
	}
	// bt, _ := library.FindBlockType("system.log")
	// fmt.Println(bt)
	// fmt.Println(bt.Replace("{{.Widgets}}"))
}
