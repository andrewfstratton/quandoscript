package block

import "github.com/andrewfstratton/quandoscript/block/widget"

type Block struct {
	Lookup  string
	widgets []widget.Widget
}

func New(lookup string) Block {
	return Block{Lookup: lookup}
}

func (block *Block) Add(widget widget.Widget) {
	block.widgets = append(block.widgets, widget)
}
