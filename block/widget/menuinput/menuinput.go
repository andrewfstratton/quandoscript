package menuinput

import (
	"fmt"
	"maps"
	"slices"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type (
	IntMapOfString map[int]string
)

type MenuInt struct {
	Name    string
	Choices IntMapOfString
}

func NewMenuInt(name string) *MenuInt {
	return &MenuInt{Name: name, Choices: make(IntMapOfString)}
}

func (widg *MenuInt) Html() (txt string) {
	txt = fmt.Sprintf(`<select data-quando-name="%v">`, widg.Name)
	keys := slices.Sorted(maps.Keys(widg.Choices))
	for key := range keys {
		val := widg.Choices[key]
		if val != "" {
			txt += fmt.Sprintf("\n<option value='%d'>%s</option>", key, val)
		}
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
