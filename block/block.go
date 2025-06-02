package block

import (
	"bytes"
	"fmt"
	"os"
	"runtime/debug"
	"testing"
	"text/template"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/block/script"
	"github.com/andrewfstratton/quandoscript/block/widget"
)

type BlockType struct {
	TypeName string
	Class    string
	widgets  []widget.Widget
	OpOp     action.OpOp
}

func New(typeName string, class string, opop action.OpOp) *BlockType {
	if typeName == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH "" BLOCK TYPE`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	if opop == nil {
		fmt.Printf("Warning: block type '%s' has nil operation\n", typeName)
	}
	if class != "" {
		class = "quando-" + class
	}
	return &BlockType{
		TypeName: typeName,
		Class:    class,
		OpOp:     opop,
	}
}

func (block *BlockType) Add(widgets ...widget.Widget) {
	// TODO: handle duplicate name
	block.widgets = append(block.widgets, widgets...)
}

func (block *BlockType) Replace(original string) string {
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
	t.Execute(&by, block)
	return by.String()
}

func (block *BlockType) Widgets() string {
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
