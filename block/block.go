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

type BlockType struct {
	typeName string
	class    string
	widgets  []widget.Widget
}

type BlockExpanded struct {
	TypeName string
	Class    string
	Widgets  string
	Params   string
}

func New(typeName string, class string) *BlockType {
	if typeName == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH "" BLOCK TYPE`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	return &BlockType{
		typeName: typeName,
		class:    class,
	}
}

func (block *BlockType) Add(widget widget.Widget) {
	block.widgets = append(block.widgets, widget)
}

func (block *BlockType) Expand() BlockExpanded {
	return BlockExpanded{
		TypeName: block.typeName,
		Class:    block.class,
		Widgets:  block.WidgetsHtml(),
		Params:   block.Params(),
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

func (block *BlockType) WidgetsHtml() string {
	wh := ""
	for _, w := range block.widgets {
		wh += w.Html()
	}
	return wh
}

func (block *BlockType) Params() string {
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
