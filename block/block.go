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
	"github.com/andrewfstratton/quandoscript/run/op"
)

type BlockType struct {
	typeName string
	class    string
	widgets  []widget.Widget
	Op       op.OpOp
}

type blockInstance struct {
	TypeName string
	Class    string
	Widgets  string
	Params   string
}

func New(typeName string, class string, op op.OpOp) *BlockType {
	if typeName == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH "" BLOCK TYPE`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	if op == nil {
		fmt.Printf("Warning: block type '%s' has nil operation\n", typeName)
	}
	return &BlockType{
		typeName: typeName,
		class:    class,
		Op:       op,
	}
}

func (block *BlockType) Add(widgets ...widget.Widget) {
	// TODO: handle duplicate name
	block.widgets = append(block.widgets, widgets...)
}

func (block *BlockType) instance() blockInstance {
	return blockInstance{
		TypeName: block.typeName,
		Class:    "quando-" + block.class,
		Widgets:  block.widgetsHtml(),
		Params:   block.params(),
	}
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
	bi := block.instance()
	t.Execute(&by, bi)
	return by.String()
}

func (block *BlockType) widgetsHtml() string {
	wh := ""
	for _, w := range block.widgets {
		wh += w.Html()
	}
	return wh
}

func (block *BlockType) params() string {
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
