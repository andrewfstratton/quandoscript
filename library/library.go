package library

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/block"
	"github.com/andrewfstratton/quandoscript/blocklist"
)

const (
	SERVER  = "server"
	UNKNOWN = ""
)

var blocks map[string]*block.BlockType         // lookup for all block types
var blocklists map[string]*blocklist.BlockList // groups of blocks by 'class' for menu
var classes []string

func NewBlockType(block_type string, class string) (b *block.BlockType) {
	_, inuse := blocks[block_type]
	if inuse {
		fmt.Println(`BLOCK "` + block_type + `" ALREADY EXISTS`)
		if testing.Testing() {
			return
		}
		debug.PrintStack()
		os.Exit(99)
	}
	b = block.New(block_type, class)
	blocks[block_type] = b
	bl, ok := blocklists[class]
	if !ok {
		bl = blocklist.New(class)
		blocklists[class] = bl
		classes = append(classes, class)
	}
	bl.Add(b)
	return
}

func FindBlockType(block_type string) (block *block.BlockType, found bool) {
	block, found = blocks[block_type]
	return
}

func Classes() []string {
	return classes
}

func init() {
	blocks = make(map[string]*block.BlockType)
	blocklists = make(map[string]*blocklist.BlockList)
}
