package stringinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type StringInput struct {
	Name     string
	default_ string
	empty    string
}

func New(name string) *StringInput {
	return &StringInput{Name: name}
}

func (si *StringInput) Html() (txt string) {
	txt = `&quot;<input data-quando-name='` + si.Name + `' type='text'`
	if si.default_ != "" {
		txt += " value='" + si.default_ + "'"
	}
	if si.empty != "" {
		txt += " placeholder='" + si.empty + "'"
	}
	txt += `/>&quot;`
	return
}

func (si *StringInput) Generate() string {
	return fmt.Sprintf(`%v"${%v}"`, si.Name, si.Name)
}

func (si *StringInput) Default(s string) *StringInput {
	si.default_ = s
	return si
}

func (si *StringInput) Empty(s string) *StringInput {
	si.empty = s
	return si
}

func (si *StringInput) Param(outer param.Params) *param.StringParam {
	return param.NewString(si.Name, "", outer)
}
