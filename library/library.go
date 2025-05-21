package library

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/blocklist"

	"github.com/andrewfstratton/quandoscript/block"
)

var blocks map[string]*block.Block
var blocklists map[string]*blocklist.BlockList

func NewBlock(qid string, class string) (b *block.Block) {
	_, inuse := blocks[qid]
	if inuse {
		fmt.Println(`BLOCK "` + qid + `" ALREADY EXISTS`)
		debug.PrintStack()
		if testing.Testing() {
			return
		}
		os.Exit(99)
	}
	b = block.New(qid, class)
	blocks[qid] = b
	bl, ok := blocklists[class]
	if !ok {
		bl = blocklist.New(class)
		blocklists[class] = bl
	}
	bl.Add(b)
	return
}

func init() {
	blocks = make(map[string]*block.Block)
	blocklists = make(map[string]*blocklist.BlockList)

}
