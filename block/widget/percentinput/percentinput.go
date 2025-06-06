package percentinput

import "fmt"

type PercentInput struct {
	Name     string
	default_ *float64
	empty    string
	width    *int
}

func New(name string) *PercentInput {
	return &PercentInput{Name: name}
}

func (pi *PercentInput) Html() (txt string) {
	txt = fmt.Sprintf("<input data-quando-name='%v' type='number'", pi.Name)
	if pi.default_ != nil {
		txt += fmt.Sprintf(" value='%v'", *pi.default_)
	}
	if pi.empty != "" {
		txt += " placeholder='" + pi.empty + "'"
	}
	if pi.width != nil {
		txt += fmt.Sprintf(" style='width:%dem'", *pi.width)
	}
	txt += ` min='0' max='100'/>%`
	return
}

func (pi *PercentInput) Generate() string {
	return fmt.Sprintf(`%v#${%v}`, pi.Name, pi.Name)
}

func (pi *PercentInput) Default(f float64) *PercentInput {
	pi.default_ = &f
	return pi
}

func (pi *PercentInput) Width(i int) *PercentInput {
	pi.width = &i
	return pi
}

func (pi *PercentInput) Empty(s string) *PercentInput {
	pi.empty = s
	return pi
}
