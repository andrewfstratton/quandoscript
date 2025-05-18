package block

import (
	"fmt"
	"os"
	"regexp"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/block/script"
	"github.com/andrewfstratton/quandoscript/block/widget"
)

type Block struct {
	lookup  string
	class   string
	widgets []widget.Widget
}

func New(lookup string) *Block {
	if lookup == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH "" LOOKUP`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	return &Block{
		lookup: lookup,
		class:  regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*").FindString(lookup),
	}
}

func (block *Block) Add(widget widget.Widget) {
	block.widgets = append(block.widgets, widget)
}

func (block *Block) html() string { // incomplete for now so not available externally
	result := ""
	for _, widget := range block.widgets {
		result += widget.Html()
	}
	return result
}

func (block *Block) script() string { // incomplete for now so not available externally
	result := ""
	for _, widget := range block.widgets {
		s, ok := widget.(script.Generator)
		if ok {
			if result != "" {
				result += ","
			}
			result += s.Generate()
		}
	}
	return result
}

func (block *Block) Class() string {
	return block.class
}
