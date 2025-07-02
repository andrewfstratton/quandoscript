package boxinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type (
	Box struct {
		Name  string
		Class string
	}
)

func New(name string, class string) *Box {
	return &Box{Name: name, Class: class}
}

func (bi *Box) Html() string {
	return fmt.Sprintf("</div><div data-quando-name='%s' class='quando-box'></div>\n<div class='quando-row %s'>\n", bi.Name, bi.Class)
}

func (bi *Box) Generate() string {
	return fmt.Sprintf(`%v:${%v}`, bi.Name, bi.Name)
}

func (bi *Box) Param(outer param.Params) *param.IdParam {
	return param.NewId(bi.Name, 0, outer)
}
