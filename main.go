package main

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block/widget"
	"github.com/andrewfstratton/quandoscript/block/widget/menuinput"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
	"github.com/andrewfstratton/quandoscript/definition"
	"github.com/andrewfstratton/quandoscript/library"
	"github.com/andrewfstratton/quandoscript/property"
)

type Defn struct {
	TypeName     widget.None        `_:"keyboard.control"`
	Class        widget.None        `_:"server-devices"`
	PressRelease menuinput.MenuInt  `2:"⇕" 1:"press" 0:"release"`
	Vari         stringinput.String `empty:"⇕" show:"PressRelease=2"`
	_            text.Text          `txt:"⌨️ Key "`
	_            text.Text          `txt:"ctrl+" show:"Ctrl=1"`
	_            text.Text          `txt:"alt+" show:"Alt=1"`
	_            text.Text          `txt:"shift+" show:"Shift=1"`
	Key          stringinput.String `empty:"🗚" length:"1"`
	_            text.Text          `txt:" "`
	_            text.Text          `txt:"<br>" hover:"true"`
	Ctrl         menuinput.MenuInt  `0:"no ctrl" 1:"ctrl" hover:"true" toggle:"true"`
	Alt          menuinput.MenuInt  `0:"no alt" 1:"alt" hover:"true" toggle:"true"`
	Shift        menuinput.MenuInt  `0:"no shift" 1:"shift" hover:"true" toggle:"true"`
}

func _init() {
	defn := Defn{}
	definition.Setup(&defn)
	library.NewBlock(defn).Op(
		func(early param.Params) func(param.Params) {
			key := defn.Key.Param(early)
			press_release := defn.PressRelease.Param(early)
			ctrl := defn.Ctrl.Param(early)
			alt := defn.Alt.Param(early)
			shift := defn.Shift.Param(early)
			vari := defn.Vari.Param(early)
			return func(late param.Params) {
				key.Update(late)
				vari.Update(late)
				press := press_release.Int() == widget.PRESS
				if press_release.Int() == widget.PRESS_RELEASE {
					press = property.GetBool(0, vari.Val)
				}
				fmt.Printf("control_keyboard.PressRelease('%s', %t, %t, %t, %t)\n", key.Val, press, ctrl.Bool(), alt.Bool(), shift.Bool())
			}
		})
}

func init() {
	_init()
}

const (
	TEST_LINES = `8 keyboard.control(Vari"a",Key"a",PressRelease#1,Ctrl#1,Alt#1,Shift#1)`
)

func main() {
	library.Parse(TEST_LINES)
	warn := action.Start()
	if warn != "" {
		fmt.Println(warn)
	}
	time.Sleep(20 * time.Second)
}
