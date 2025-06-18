package main

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block/widget/idinput"
	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
	"github.com/andrewfstratton/quandoscript/library"
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
	library.Block(log).Op(
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

type Defn struct {
	TypeName struct{}                `_:"system.after"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"After " iconify:"true"`
	Secs     numberinput.NumberInput `empty:"seconds" min:"0" max:"999" width:"4" default:"1"` // min:0 max:999 width:4 default:1`
	_        text.Text               `txt:"seconds"`
	Callback idinput.IdInput
}

func init_after() {
	defn := &Defn{}
	library.Block(defn).Op(
		func(early param.Params) func(param.Params) {
			secs := defn.Secs.Param(early)
			callback := defn.Callback.Param(early)
			return func(late param.Params) {
				secs.Update(late)
				callback.Update(late)
				time.AfterFunc(secs.Duration(), func() {
					action.Run(callback.Val)
				})
			}
		})
}

type Every struct {
	TypeName struct{}                `_:"system.every"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"Every " iconify:"true"`
	Secs     numberinput.NumberInput `empty:"seconds" min:"0" max:"999" width:"4" default:"1"`
	_        text.Text               `txt:"seconds"`
	Callback idinput.IdInput
}

func init_every() {
	defn := &Every{}
	library.Block(defn).Op(
		func(early param.Params) func(param.Params) {
			secs := defn.Secs.Param(early)
			callback := defn.Callback.Param(early)
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
	TEST_LINES = `0 system.log(Greeting"Hi",Txt"Bob")
1 system.after(Secs#2,Callback:3)
2 system.every(Secs#0.5,Callback:7)

3 system.log(Greeting"Hello",Txt"Jane")
4 system.log(Greeting"Bye",Txt"Bob")
5 system.after(Secs#0.9,Callback:6)

6 system.log(Greeting"Bye",Txt"Jane")

7 system.log(Txt".")
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
