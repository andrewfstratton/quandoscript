package library

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/block"
)

var blocks map[string]block.Block

func NewBlock(lookup string) (b *block.Block) {
	_, inuse := blocks[lookup]
	if inuse {
		fmt.Println(`BLOCK "` + lookup + `" ALREADY EXISTS`)
		debug.PrintStack()
		os.Exit(99)
	}
	b = block.New(lookup)
	if testing.Testing() {
		return
	}
	blocks[lookup] = *b
	return
}
