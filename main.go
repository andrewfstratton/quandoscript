package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
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

func parseLines(in string) {
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		lineid, word, params, err := parse.Line(scanner.Text())
		fmt.Println(lineid, word, params, err)
		o := library.NewAction(word, params, nil)
		action.IdOrdered = append(action.IdOrdered, lineid)
		action.Actions[lineid] = o
	}
}

const (
	TEST_LINES = `0 system.log(greeting"Hi",txt"Bob")
1 system.log(greeting"Hello",txt"Jane")
2 system.log(greeting"Hello",txt"Jane")
3 system.log(greeting"Hello",txt"Jane")
4 system.log(greeting"Hello",txt"Jane")
5 system.log(greeting"Hello",txt"Jane")
6 system.log(greeting"Hello",txt"Jane")
7 system.log(greeting"Hello",txt"Jane")`
)

func main() {
	parseLines(TEST_LINES)
	for l := range action.IdOrdered {
		fmt.Print(strconv.Itoa(l) + "-")
		action.Actions[l].Exec()
	}
}
