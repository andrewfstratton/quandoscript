package numberinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type (
	Pfloat *float64
	Pint   *int64
)

type (
	Number struct {
		Name    string
		Default Pfloat
		Empty   string
		Width   Pint
		Min     Pint
		Max     Pint
	}
	S_ecs struct {
		Secs Number `empty:"seconds" min:"0" max:"999" width:"4" default:"1"`
	}
)

func New(name string) *Number {
	return &Number{Name: name}
}

func (ni *Number) Html() (txt string) {
	txt = fmt.Sprintf("<input data-quando-name='%v' type='number'", ni.Name)
	if ni.Default != nil {
		txt += fmt.Sprintf(" value='%v'", *ni.Default)
	}
	if ni.Empty != "" {
		txt += " placeholder='" + ni.Empty + "'"
	}
	if ni.Width != nil {
		txt += fmt.Sprintf(" style='width:%dem'", *ni.Width)
	}
	// needs '' around number to avoid issues?!
	if ni.Min != nil {
		txt += fmt.Sprintf(" min='%d'", *ni.Min)
	}
	if ni.Max != nil {
		txt += fmt.Sprintf(" max='%d'", *ni.Max)
	}
	txt += `/>`
	return
}

func (ni *Number) Generate() string {
	return fmt.Sprintf(`%v#${%v}`, ni.Name, ni.Name)
}

func (ni *Number) Param(outer param.Params) *param.NumberParam {
	return param.NewNumber(ni.Name, 0, outer)
}
