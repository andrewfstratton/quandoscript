package boxinput

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/action/param"
)

type BoxInput struct {
	Name string
}

func New(name string) *BoxInput {
	return &BoxInput{Name: name}
}

func (bi *BoxInput) Html() (txt string) {
	txt = fmt.Sprintf("</div><div data-quando-name='%s' class='quando-box'></div>\n<div class='quando-row quando-time'>\n", bi.Name)
	return
}

func (bi *BoxInput) Generate() string {
	return fmt.Sprintf(`%v:${%v}`, bi.Name, bi.Name)
}

func (bi *BoxInput) Param(outer param.Params) *param.IdParam {
	return param.NewId(bi.Name, 0, outer)
}
