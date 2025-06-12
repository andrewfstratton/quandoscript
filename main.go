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
	_ "github.com/andrewfstratton/quandoscript/library"
	"github.com/andrewfstratton/quandoscript/parse"
)

type LogDefn struct {
	TypeName struct{}                `_:"system.log"`
	Class    struct{}                `_:"misc"`
	_        text.Text               `txt:"Log "`
	Greeting stringinput.StringInput `empty:"greeting"`
	_        text.Text               `txt:" "`
	Txt      stringinput.StringInput `empty:"name"`
}

func init_log() {
	log := &LogDefn{}
	b := block.New(log)
	fmt.Println("block = ", b)
	b.Op(
		func(outer param.Params) func(param.Params) {
			greeting := log.Greeting.Param(outer)
			txt := log.Txt.Param(outer)
			return func(inner param.Params) {
				txt.Update(inner)
				greeting.Update(inner)
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
	// Secs     numberinput.NumberInput `empty:"seconds" _:"min:0,max:999,width:4,default:1` // min:0 max:999 width:4 default:1`
	_        text.Text `txt:" "`
	Callback idinput.IdInput
}

func init_after() {
	after := &AfterDefn{}
	b := block.New(after)
	fmt.Println("block = ", b)
	b.Op(
		func(outer param.Params) func(param.Params) {
			secs := after.Secs.Param(outer)
			callback := after.Callback.Param(outer)
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
