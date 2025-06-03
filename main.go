package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
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

func init_after() {
	seconds := numberinput.New("secs").Empty("seconds").Min(0).Max(999).Width(4).Default(1)
	callback := numberinput.New("callback").Empty("callback").Min(0).Max(999).Width(4).Default(999)

	library.NewBlockType("system.after", "misc", timeop).Add(
		text.New("After "),
		seconds,
		text.New("secs"),
		callback,
	)
}

func timeop(outer param.Params) func(param.Params) {
	secs := param.NumberParam{Lookup: "secs", Val: 1.0}
	secs.Update(outer)
	callback := param.IdParam{Lookup: "callback", Val: 0}
	callback.Update(outer)
	return func(inner param.Params) {
		secs.Update(inner)
		callback.Update(inner)
		// fmt.Println("After:", secs.Val, "secs, callback:", callback)
		isecs := int(secs.Val)
		time.AfterFunc(time.Duration(isecs)*time.Second, func() {
			action.Run(callback.Val)
		})
	}
}

func init() {
	init_log()
	init_after()
}

func parseLines(in string) {
	scanner := bufio.NewScanner(strings.NewReader(in))
	new_group := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// fmt.Println("End of main block")
			new_group = true
			continue
		}
		lineid, word, params, err := parse.Line(scanner.Text())
		if err != nil {
			fmt.Println(lineid, word, params, err)
		}
		o := library.NewAction(word, params, nil)
		if new_group {
			action.NewGroup()
			new_group = false
		}
		action.Add(lineid, o)
	}
}

const (
	TEST_LINES = `0 system.log(greeting"Hi",txt"Bob")
1 system.after(secs#2,callback:3)

3 system.log(greeting"Hello",txt"Jane")
4 system.log(greeting"Bye",txt"Bob")
5 system.after(secs#2,callback:6)

6 system.log(greeting"Bye",txt"Jane")
`
)

func main() {
	parseLines(TEST_LINES)
	msg := action.Start()
	if msg != "" {
		fmt.Println(msg)
	}
	time.Sleep(10 * time.Second)
}
