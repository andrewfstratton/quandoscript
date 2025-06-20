package menu

import (
	"github.com/andrewfstratton/quandoscript/block"
)

const (
	UNKNOWN_CLASS = "unknown"
)

type Menu struct {
	Class  string
	Blocks []block.Block
}

func New(class string) *Menu {
	return &Menu{Class: class}
}

func (menu *Menu) Add(block *block.Block) {
	menu.Blocks = append(menu.Blocks, *block)
}

func (menu *Menu) CSSClass(prefix string) string {
	if menu.Class == "" {
		return prefix + UNKNOWN_CLASS
	}
	return prefix + menu.Class
}
