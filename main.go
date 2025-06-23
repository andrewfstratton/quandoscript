package main

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block/widget/boxinput"
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
	TypeName struct{}                `_:"time.after"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"After " iconify:"true"`
	Secs     numberinput.NumberInput `empty:"seconds" min:"0" max:"999" width:"4" default:"1"` // min:0 max:999 width:4 default:1`
	_        text.Text               `txt:"seconds"`
	Box      boxinput.BoxInput
}

func init_after() {
	defn := &Defn{}
	library.Block(defn).Op(
		func(early param.Params) func(param.Params) {
			secs := defn.Secs.Param(early)
			callback := defn.Box.Param(early)
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
	TypeName struct{}                `_:"time.every"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"Every " iconify:"true"`
	Secs     numberinput.NumberInput `empty:"seconds" min:"0" max:"999" width:"4" default:"1"`
	_        text.Text               `txt:"seconds"`
	Callback boxinput.BoxInput       `txt:"box"`
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
	TEST_LINES = `1 system.log(Greeting"1",Txt"1")
2 system.log(Greeting"2",Txt"2")
3 system.log(Greeting"3",Txt"3")
4 time.after(Secs#4,Box:5)
6 time.after(Secs#6,Box:0)
7 time.after(Secs#7,Box:8)

5 system.log(Greeting"5",Txt"5")

8 system.log(Greeting"8",Txt"8")
9 time.after(Secs#9,Box:10)

10 system.log(Greeting"10",Txt"10")
`
)

func main() {
	library.Parse(TEST_LINES)
	warn := action.Start()
	if warn != "" {
		fmt.Println(warn)
	}
	time.Sleep(20 * time.Second)
}
