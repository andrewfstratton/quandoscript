package menuinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type (
	IntStr struct {
		Key int
		Val string
	}
	IntStrings []IntStr
)

type MenuInt struct {
	Name    string
	Hover   bool
	Choices IntStrings
}

func NewMenuInt(name string) *MenuInt {
	return &MenuInt{Name: name, Choices: make(IntStrings, 0)}
}

func (widg *MenuInt) Html() (txt string) {
	txt = fmt.Sprintf(`<select data-quando-name="%v"`, widg.Name)
	if widg.Hover {
		txt += ` class="hover-display"`
	}
	txt += `>`
	for _, choice := range widg.Choices {
		txt += fmt.Sprintf("\n<option value='%d'>%s</option>", choice.Key, choice.Val)
	}
	txt += `\n</select>`
	return
}

func (widg *MenuInt) Generate() string {
	return fmt.Sprintf(`%v#${%v}`, widg.Name, widg.Name)
}

func (widg *MenuInt) Param(outer param.Params) *param.NumberParam {
	return param.NewNumber(widg.Name, 0, outer)
}
