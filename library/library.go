package library

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/block"
	"github.com/andrewfstratton/quandoscript/blocklist"
)

var blocks map[string]*block.Block             // lookup for all blocks on qid
var blocklists map[string]*blocklist.BlockList // groups of blocks by 'class' for menu
var classes []string

func NewBlock(qid string, class string) (b *block.Block) {
	_, inuse := blocks[qid]
	if inuse {
		fmt.Println(`BLOCK "` + qid + `" ALREADY EXISTS`)
		if testing.Testing() {
			return
		}
		debug.PrintStack()
		os.Exit(99)
	}
	b = block.New(qid, class)
	blocks[qid] = b
	bl, ok := blocklists[class]
	if !ok {
		bl = blocklist.New(class)
		blocklists[class] = bl
		classes = append(classes, class)
	}
	bl.Add(b)
	return
}

func FindBlock(qid string) (block *block.Block, found bool) {
	block, found = blocks[qid]
	return
}

func Classes() []string {
	return classes
}

func init() {
	blocks = make(map[string]*block.Block)
	blocklists = make(map[string]*blocklist.BlockList)
}
