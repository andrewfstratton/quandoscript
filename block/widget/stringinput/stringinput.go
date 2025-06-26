package stringinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block/widget"
)

type String struct {
	Name    string
	Default string
	Empty   string
	Length  string
	Show    string
}

func New(name string) *String {
	return &String{Name: name}
}

func (si *String) Html() (txt string) {
	txt = `&quot;<input data-quando-name='` + si.Name + `' type='text'`
	if si.Default != "" {
		txt += " value='" + si.Default + "'"
	}
	if si.Empty != "" {
		txt += " placeholder='" + si.Empty + "'"
	}
	if si.Length != "" {
		txt += " maxlength='" + si.Length + "'"
	}
	txt += `/>&quot;`
	if si.Show != "" {
		txt = widget.OpenCloseTag(txt, `span data-quando-toggle="`+si.Show+`"`, "span")
	}
	return
}

func (si *String) Generate() string {
	return fmt.Sprintf(`%v"${%v}"`, si.Name, si.Name)
}

func (si *String) Param(outer param.Params) *param.StringParam {
	return param.NewString(si.Name, "", outer)
}
