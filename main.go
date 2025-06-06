package main

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block"
	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
	"github.com/andrewfstratton/quandoscript/parse"
)

func init_log() {
	greetingUI := stringinput.New("greeting").Empty("greeting")
	txtUI := stringinput.New("txt").Empty("greeting")

	block.AddNew("system.log", "misc",
		text.New("Log "),
		greetingUI,
		text.New(" "),
		txtUI,
	).Op(
		func(outer param.Params) func(param.Params) {
			greeting := param.NewString(greetingUI.Name, "", outer)
			txt := param.NewString(txtUI.Name, "", outer)
			return func(inner param.Params) {
				txt.Update(inner)
				greeting.Update(inner)
				now := time.Now()
				fmt.Println("Log:", greeting.Val, txt.Val, now.Format(time.TimeOnly))
			}
		})
}

func init_after() {
	secsUI := numberinput.New("secs").Empty("seconds").Min(0).Max(999).Width(4).Default(1)
	callbackUI := numberinput.New("callback").Empty("callback").Min(0).Max(999).Width(4).Default(999)

	block.AddNew("system.after", "misc",
		text.New("After "),
		secsUI,
		text.New("secs"),
		callbackUI,
	).Op(
		func(outer param.Params) func(param.Params) {
			secs := param.NewNumber(secsUI.Name, 1.0, outer)
			callback := param.NewId(callbackUI.Name, 0, outer)
			return func(inner param.Params) {
				secs.Update(inner)
				callback.Update(inner)
				// fmt.Println("After:", secs.Val, "secs, callback:", callback)
				isecs := int(secs.Val)
				time.AfterFunc(time.Duration(isecs)*time.Second, func() {
					action.Run(callback.Val)
				})
			}
		})
}

func init() {
	init_log()
	init_after()
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
	parse.Lines(TEST_LINES)
	warn := action.Start()
	if warn != "" {
		fmt.Println(warn)
	}
	time.Sleep(10 * time.Second)
}
