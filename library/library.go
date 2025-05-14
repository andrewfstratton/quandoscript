package library

import (
	"github.com/andrewfstratton/quandoscript/block"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
)

var blocks map[string]block.Block

func init() {
	block := block.New("system.log")
	log := text.New("Log").Bold()
	block.Add(log)
	blocks[block.Lookup] = block
}
