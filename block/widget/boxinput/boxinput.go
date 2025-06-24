package boxinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type Box struct {
	Name string
}

func New(name string) *Box {
	return &Box{Name: name}
}

func (bi *Box) Html() (txt string) {
	txt = fmt.Sprintf("</div><div data-quando-name='%s' class='quando-box'></div>\n<div class='quando-row quando-time'>\n", bi.Name)
	return
}

func (bi *Box) Generate() string {
	return fmt.Sprintf(`%v:${%v}`, bi.Name, bi.Name)
}

func (bi *Box) Param(outer param.Params) *param.IdParam {
	return param.NewId(bi.Name, 0, outer)
}
