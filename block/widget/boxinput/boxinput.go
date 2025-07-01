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
	B_ox struct {
		Box Box
	}
)

func New(name string, class string) *Box {
	return &Box{Name: name, Class: class}
}

func (bi *Box) Html() (txt string) {
	txt = fmt.Sprintf("</div><div data-quando-name='%s' class='quando-box'></div>\n<div class='quando-row %s'>\n", bi.Name, bi.Class)
	return
}

func (bi *Box) Generate() string {
	return fmt.Sprintf(`%v:${%v}`, bi.Name, bi.Name)
}

func (bi *Box) Param(outer param.Params) *param.IdParam {
	return param.NewId(bi.Name, 0, outer)
}
