package blocklist

import (
	"github.com/andrewfstratton/quandoscript/block"
)

type List struct {
	class  string
	blocks []block.Block
}

func New(class string) (list List) {
	return List{class: class}
}

func (list *List) Add(block *block.Block) {
	list.blocks = append(list.blocks, *block)
}

func (list *List) Class() (css_class string) {
	css_class = "quando"
	if list.class != "" {
		css_class += "-" + list.class
	}
	return
}
