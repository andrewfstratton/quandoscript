package block

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/block/script"
	"github.com/andrewfstratton/quandoscript/block/widget"
)

type Block struct {
	qid     string
	class   string
	widgets []widget.Widget
}

type BlockOutput struct {
	qid        string
	class      string
	params     string
	widgetHtml string
}

func New(qid string, class string) *Block {
	if qid == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH "" QUANDO ID`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	return &Block{
		qid:   qid,
		class: class,
	}
}

func (block *Block) Add(widget widget.Widget) {
	block.widgets = append(block.widgets, widget)
}

func (block *Block) Output() BlockOutput {
	wHtml := ""
	for _, w := range block.widgets {
		wHtml += w.Html()
	}
	return BlockOutput{
		qid:        block.qid,
		class:      block.class,
		params:     block.params(),
		widgetHtml: wHtml,
	}
}

func (block *Block) params() string {
	result := ""
	for _, w := range block.widgets {
		s, ok := w.(script.Generator)
		if ok {
			if result != "" { // separate parameters with comma
				result += ","
			}
			result += s.Generate()
		}
	}
	return result
}
