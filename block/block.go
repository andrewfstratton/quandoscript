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

func (block *Block) ReplaceHtml() string { // incomplete for now
	result := `<div data-quando-block-type="` +
		block.qid +
		`" class="quando-block"` +
		//` data-quando-id="true"` + // Removed since always used now
		// Note that quandoscript needs unique block id (data-quando-id) at start of line + space when generating quandoscript call
		` data-quando-quandoscript='` + // single quote to allow quandoscript double quote string embedding
		block.qid + `(` +
		block.params() +
		`)'>` +
		`<div class="quando-left quando-` + block.class + `"></div>` +
		`<div class="quando-right">` +
		`<div class="quando-row quando-` + block.class + `">`
	for _, w := range block.widgets {
		result += w.Html()
	}
	result += `</div></div></div>`
	return result
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
