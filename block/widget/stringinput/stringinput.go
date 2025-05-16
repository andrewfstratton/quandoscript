package stringinput

import "fmt"

type StringInput struct {
	name     string
	default_ string
	empty    string
}

func New(name string) *StringInput {
	return &StringInput{name: name}
}

func (si *StringInput) Html() (txt string) {
	txt = `&quot;<input data-quando-name='` + si.name + `' type='text'`
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
	return fmt.Sprintf(`%v"${%v}"`, si.name, si.name)
}

func (si *StringInput) Default(s string) *StringInput {
	si.default_ = s
	return si
}

func (si *StringInput) Empty(s string) *StringInput {
	si.empty = s
	return si
}
