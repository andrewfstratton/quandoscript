package menuinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type (
	IntString struct {
		Key int
		Val string
	}
	StringString struct {
		Key string
		Val string
	}
)

type (
	MenuInt struct {
		Name    string
		Hover   bool
		Toggle  bool
		Choices []IntString
	}
	MenuStr struct {
		Name    string
		Hover   bool
		Toggle  bool
		Choices []StringString
	}
)

func NewMenuInt(name string) *MenuInt {
	return &MenuInt{Name: name, Choices: make([]IntString, 0)}
}

func NewMenuStr(name string) *MenuStr {
	return &MenuStr{Name: name, Choices: make([]StringString, 0)}
}

func (widg *MenuInt) Html() (txt string) {
	txt = htmlStart(widg.Name, widg.Hover, widg.Toggle)
	for _, choice := range widg.Choices {
		txt += fmt.Sprintf("\n<option value='%d'>%s</option>", choice.Key, choice.Val)
	}
	txt += `\n</select>`
	return
}

func (widg *MenuStr) Html() (txt string) {
	txt = htmlStart(widg.Name, widg.Hover, widg.Toggle)
	for _, choice := range widg.Choices {
		txt += fmt.Sprintf("\n<option value='%s'>%s</option>", choice.Key, choice.Val)
	}
	txt += `\n</select>`
	return
}

func ifStr(b bool, s string) string {
	if b {
		return s
	}
	return ""
}

func htmlStart(name string, hover, toggle bool) string {
	cls := ifStr(hover, `hover-display `) + ifStr(toggle, `quando-toggle `)
	if cls != "" {
		cls = `class="` + cls + `" `
	}
	return fmt.Sprintf(`<select data-quando-name="%v" %s>`, name, cls)
}

func (widg *MenuInt) Generate() string {
	return fmt.Sprintf(`%v#${%v}`, widg.Name, widg.Name)
}

func (widg *MenuStr) Generate() string {
	return fmt.Sprintf(`%v"${%v}"`, widg.Name, widg.Name)
}

func (widg *MenuInt) Param(outer param.Params) *param.NumberParam {
	return param.NewNumber(widg.Name, 0, outer)
}

func (widg *MenuStr) Param(outer param.Params) *param.StringParam {
	return param.NewString(widg.Name, "", outer)
}
