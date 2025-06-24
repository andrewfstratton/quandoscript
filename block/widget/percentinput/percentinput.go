package percentinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
)

type Percent struct {
	Name    string
	Default numberinput.Pfloat
	Empty   string
	Width   numberinput.Pint
}

func New(name string) *Percent {
	return &Percent{Name: name}
}

func (pi *Percent) Html() (txt string) {
	txt = fmt.Sprintf("<input data-quando-name='%v' type='number'", pi.Name)
	if pi.Default != nil {
		txt += fmt.Sprintf(" value='%v'", *pi.Default)
	}
	if pi.Empty != "" {
		txt += " placeholder='" + pi.Empty + "'"
	}
	if pi.Width != nil {
		txt += fmt.Sprintf(" style='width:%dem'", *pi.Width)
	}
	txt += ` min='0' max='100'/>%`
	return
}

func (pi *Percent) Generate() string {
	return fmt.Sprintf(`%v#${%v}`, pi.Name, pi.Name)
}
