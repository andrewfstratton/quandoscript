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

type Block struct {
	TypeName string
	Class    string
	widgets  []widget.Widget
	Early    action.Early
}

var AddToLibrary func(*Block) // injected by library

func AddNew(typeName string, class string, early action.Early, widgets ...widget.Widget) (block *Block) {
	if typeName == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH "" BLOCK TYPE`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	if early == nil {
		fmt.Printf("Warning: block type '%s' has nil operation\n", typeName)
	}
	if class != "" {
		class = "quando-" + class
	}
	block = &Block{
		TypeName: typeName,
		Class:    class,
		Early:    early,
	}
	block.widgets = append(block.widgets, widgets...)
	if testing.Testing() && AddToLibrary == nil { // handle tests when AddToLibrary has not been injected by library.init()
		return
	}
	AddToLibrary(block)
	return
}

func (block *Block) Replace(original string) string {
	var buf bytes.Buffer
	t, err := template.New("tmp").Parse(original)
	if err != nil {
		fmt.Println(`TEMPLATE PARSING ERROR`)
		if testing.Testing() {
			return ""
		}
		debug.PrintStack()
		os.Exit(99)
	}
	t.Execute(&buf, block)
	return buf.String()
}

func (block *Block) Widgets() (asHtml string) {
	for _, w := range block.widgets {
		asHtml += w.Html()
	}
	return
}

func (block *Block) Params() (asHtml string) {
	for _, w := range block.widgets {
		s, ok := w.(script.Generator)
		if ok {
			if asHtml != "" { // separate parameters with comma
				asHtml += ","
			}
			asHtml += s.Generate()
		}
	}
	return asHtml
}
