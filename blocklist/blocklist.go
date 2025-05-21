package blocklist

import (
	"github.com/andrewfstratton/quandoscript/block"
)

const (
	UNKNOWN_CLASS = "unknown"
)

type BlockList struct {
	class  string
	blocks []block.Block
}

func New(class string) *BlockList {
	return &BlockList{class: class}
}

func (blocklist *BlockList) Add(block *block.Block) {
	blocklist.blocks = append(blocklist.blocks, *block)
}

func (blocklist *BlockList) CSSClass(prefix string) string {
	if blocklist.class == "" {
		return prefix + UNKNOWN_CLASS
	}
	return prefix + blocklist.class
}
