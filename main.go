package main

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block/widget/boxinput"
	"github.com/andrewfstratton/quandoscript/block/widget/menuinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
	"github.com/andrewfstratton/quandoscript/library"
)

type Defn struct {
	TypeName     struct{}          `_:"gamepad.button"`
	Class        struct{}          `_:"server-devices"`
	_            text.Text         `txt:"🕹️️️️️ When "`
	Index        menuinput.MenuInt `0:"Ⓐ/✕" 1:"Ⓑ/◯" 2:"Ⓧ/☐" 3:"Ⓨ/🛆" 14:"🠈" 15:"🠊" 12:"🠉" 13:"🠋" 4:"👈 Bumper" 5:"👉 Bumper" 10:"📍👈" 11:"👉📍" 8:"Back 👈" 9:"👉 Start"`
	_            text.Text         `txt:" button " iconify:"true"`
	PressRelease menuinput.MenuInt `2:"⇕" 1:"Press" 0:"Release"`
	Box          boxinput.Box
}

func _init() {
	defn := &Defn{}
	library.Block(defn).Op(
		func(early param.Params) func(param.Params) {
			index := defn.Index.Param(early)
			return func(late param.Params) {
				index.Update(late)
				fmt.Println("Button :", index.Val)
			}
		})
}

func init() {
	_init()
}

const (
	TEST_LINES = `11 gamepad.button(Index#10)`
)

func main() {
	library.Parse(TEST_LINES)
	warn := action.Start()
	if warn != "" {
		fmt.Println(warn)
	}
	time.Sleep(20 * time.Second)
}
