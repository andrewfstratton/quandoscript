package main

import (
	"fmt"
	"quandoscript/action"
	"quandoscript/action/param"
	"quandoscript/block"
	"quandoscript/block/widget/idinput"
	"quandoscript/block/widget/numberinput"
	"quandoscript/block/widget/stringinput"
	"quandoscript/block/widget/text"
	"quandoscript/library"
	"time"
)

type LogDefn struct {
	TypeName struct{}                `_:"system.log"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"Log "`
	Greeting stringinput.StringInput `empty:"greeting"`
	_        text.Text               `txt:" "`
	Txt      stringinput.StringInput `empty:"name"`
}

func init_log() { // sets up the Block (UI) ONLY  - doesn't setup any action even though the early/late functions are referenced
	log := &LogDefn{}
	block.New(log).Op(
		func(early param.Params) func(param.Params) {
			greeting := log.Greeting.Param(early)
			txt := log.Txt.Param(early)
			return func(late param.Params) {
				txt.Update(late)
				greeting.Update(late)
				now := time.Now()
				fmt.Println("Log:", greeting.Val, txt.Val, now.Format(time.TimeOnly))
			}
		})
}

type AfterDefn struct {
	TypeName struct{}                `_:"system.after"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"After " iconify:"true"`
	Secs     numberinput.NumberInput `empty:"seconds" min:"0" max:"999" width:"4" default:"1"` // min:0 max:999 width:4 default:1`
	_        text.Text               `txt:"seconds"`
	Callback idinput.IdInput
}

func init_after() {
	after := &AfterDefn{}
	block.New(after).Op(
		func(early param.Params) func(param.Params) {
			secs := after.Secs.Param(early)
			callback := after.Callback.Param(early)
			return func(late param.Params) {
				secs.Update(late)
				callback.Update(late)
				time.AfterFunc(secs.Duration(), func() {
					action.Run(callback.Val)
				})
			}
		})
}

type EveryDefn struct {
	TypeName struct{}                `_:"system.every"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"Every " iconify:"true"`
	Secs     numberinput.NumberInput `empty:"seconds" min:"0" max:"999" width:"4" default:"1"`
	_        text.Text               `txt:"seconds"`
	Callback idinput.IdInput
}

func init_every() {
	every := &EveryDefn{}
	block.New(every).Op(
		func(early param.Params) func(param.Params) {
			secs := every.Secs.Param(early)
			callback := every.Callback.Param(early)
			return func(late param.Params) {
				secs.Update(late)
				callback.Update(late)
				go func() {
					for range time.Tick(secs.Duration()) {
						action.Run(callback.Val)
					}
				}()
			}
		})
}
func init() {
	init_log()
	init_after()
	init_every()
}

const (
	TEST_LINES = `0 system.log(greeting"Hi",txt"Bob")
1 system.after(secs#2,callback:3)
2 system.every(secs#0.5,callback:7)

3 system.log(greeting"Hello",txt"Jane")
4 system.log(greeting"Bye",txt"Bob")
5 system.after(secs#0.9,callback:6)

6 system.log(greeting"Bye",txt"Jane")

7 system.log(txt".")
`
)

func main() {
	library.Parse(TEST_LINES)
	warn := action.Start()
	if warn != "" {
		fmt.Println(warn)
	}
	time.Sleep(10 * time.Second)
}
