package main

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block"
	"github.com/andrewfstratton/quandoscript/block/widget/idinput"
	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
	"github.com/andrewfstratton/quandoscript/parse"
)

func init_log() {
	_greeting := stringinput.New("greeting").Empty("greeting")
	_txt := stringinput.New("txt").Empty("greeting")

	block.AddNew("system.log", "misc",
		text.New("Log "),
		_greeting,
		text.New(" "),
		_txt,
	).Op(
		func(outer param.Params) func(param.Params) {
			greeting := _greeting.Param(outer)
			txt := _txt.Param(outer)
			return func(inner param.Params) {
				txt.Update(inner)
				greeting.Update(inner)
				now := time.Now()
				fmt.Println("Log:", greeting.Val, txt.Val, now.Format(time.TimeOnly))
			}
		})
}

func init_after() {
	_secs := numberinput.New("secs").Empty("seconds").Min(0).Max(999).Width(4).Default(1)
	_callback := idinput.New("callback")
	block.AddNew("system.after", "misc",
		text.New("After "),
		_secs,
		text.New("secs"),
		_callback,
	).Op(
		func(outer param.Params) func(param.Params) {
			secs := _secs.Param(outer)
			callback := _callback.Param(outer)
			return func(inner param.Params) {
				//  inner.Update(&secs, &callback)
				secs.Update(inner)
				callback.Update(inner)
				// fmt.Println("After:", secs.Val, "secs, callback:", callback)
				time.AfterFunc(time.Duration(secs.Int())*time.Second, func() {
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
