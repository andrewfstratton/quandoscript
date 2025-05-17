package numberinput

import "fmt"

type NumberInput struct {
	name     string
	default_ *float32
	empty    string
	width    *int
	min      *int
	max      *int
}

func New(name string) *NumberInput {
	return &NumberInput{name: name}
}

func (ni *NumberInput) Html() (txt string) {
	txt = fmt.Sprintf("<input data-quando-name='%v' type='number'", ni.name)
	if ni.default_ != nil {
		txt += fmt.Sprintf(" value='%v'", *ni.default_)
	}
	if ni.empty != "" {
		txt += " placeholder='" + ni.empty + "'"
	}
	if ni.width != nil {
		txt += fmt.Sprintf(" style='width:%dem'", *ni.width)
	}
	// needs '' around number to avoid issues?!
	if ni.min != nil {
		txt += fmt.Sprintf(" min='%d'", *ni.min)
	}
	if ni.max != nil {
		txt += fmt.Sprintf(" max='%d'", *ni.max)
	}
	txt += `/>`
	return
}

func (ni *NumberInput) Generate() string {
	return fmt.Sprintf(`%v#${%v}`, ni.name, ni.name)
}

func (ni *NumberInput) Default(f float32) *NumberInput {
	ni.default_ = &f
	return ni
}

func (ni *NumberInput) Min(i int) *NumberInput {
	ni.min = &i
	return ni
}

func (ni *NumberInput) Max(i int) *NumberInput {
	ni.max = &i
	return ni
}

func (ni *NumberInput) Width(i int) *NumberInput {
	ni.width = &i
	return ni
}

func (ni *NumberInput) Empty(s string) *NumberInput {
	ni.empty = s
	return ni
}
