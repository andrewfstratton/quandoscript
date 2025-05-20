package library

import (
	"fmt"
	"github.com/andrewfstratton/quandoscript/blocklist"
	"os"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/block"
)

const (
	UNKNOWN  = ""
	SYSTEM   = "system"
	ADVANCED = "advanced"
)

var blocks map[string]*block.Block
var blocklists map[string]blocklist.List

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
