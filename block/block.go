package block

import (
	"bytes"
	"fmt"
	"os"
	"runtime/debug"
	"testing"
	"text/template"

	"github.com/andrewfstratton/quandoscript/block/script"
	"github.com/andrewfstratton/quandoscript/block/widget"
)

type Block struct {
	qid     string
	class   string
	widgets []widget.Widget
}

type BlockExpanded struct {
	QID     string
	Class   string
	Widgets string
	Params  string
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

func (block *Block) Expand() BlockExpanded {
	return BlockExpanded{
		QID:     block.qid,
		Class:   block.class,
		Widgets: block.WidgetsHtml(),
		Params:  block.Params(),
	}
}

func (blockExpanded *BlockExpanded) Replace(original string) string {
	var by bytes.Buffer
	t, err := template.New("tmp").Parse(original)
	if err != nil {
		fmt.Println(`TEMPLATE PARSING ERROR`)
		if testing.Testing() {
			return ""
		}
		debug.PrintStack()
		os.Exit(99)
	}
	t.Execute(&by, blockExpanded)
	return by.String()
}

func (block *Block) WidgetsHtml() string {
	wh := ""
	for _, w := range block.widgets {
		wh += w.Html()
	}
	return wh
}

func (block *Block) Params() string {
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
