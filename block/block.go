package block

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/andrewfstratton/quandoscript/block/widget"
)

type Block struct {
	Lookup  string
	widgets []widget.Widget
}

func New(lookup string) Block {
	if lookup == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH ""`)
		debug.PrintStack()
		os.Exit(99)
	}
	return Block{Lookup: lookup}
}

func (block *Block) Add(widget widget.Widget) {
	block.widgets = append(block.widgets, widget)
}
