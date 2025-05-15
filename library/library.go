package library

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/andrewfstratton/quandoscript/block"
)

var blocks map[string]block.Block

func NewBlock(lookup string) block.Block {
	_, inuse := blocks[lookup]
	if inuse {
		fmt.Println(`BLOCK "` + lookup + `" ALREADY EXISTS`)
		debug.PrintStack()
		os.Exit(99)
	}
	b := block.New(lookup)
	blocks[lookup] = b
	return b
}
