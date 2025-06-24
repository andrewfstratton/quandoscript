package block

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime/debug"
	"testing"
	"text/template"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/block/script"
	"github.com/andrewfstratton/quandoscript/block/widget"
	"github.com/andrewfstratton/quandoscript/block/widget/boxinput"
	"github.com/andrewfstratton/quandoscript/block/widget/menuinput"
	"github.com/andrewfstratton/quandoscript/block/widget/numberinput"
	"github.com/andrewfstratton/quandoscript/block/widget/percentinput"
	"github.com/andrewfstratton/quandoscript/block/widget/stringinput"
	"github.com/andrewfstratton/quandoscript/block/widget/text"
	"github.com/andrewfstratton/quandoscript/definition"
)

type Block struct {
	TypeName string
	Class    string
	widgets  []widget.Widget
	Early    action.Early
}

func CreateFromDefinition(defn any) (block *Block) {
	t := reflect.TypeOf(defn).Elem() // i.e. pointer to struct
	// 	N.B. TypeName and Class exist in Defn - not in widgets
	block = &Block{}
	for i := range t.NumField() {
		f := t.Field(i)
		tag := f.Tag
		underscoreTag := tag.Get("_")
		if f.Name == "TypeName" {
			block.TypeName = underscoreTag
			continue
		}
		if f.Name == "Class" {
			block.Class = underscoreTag
			continue
		}
		// Otherwise check the type
		var w widget.Widget
		fullTypeName := f.Type.PkgPath() + "." + f.Type.Name()
		_, typeName := path.Split(fullTypeName) // split on last /
		switch typeName {
		case "text.Text":
			w = text.New()
		case "stringinput.String":
			w = stringinput.New(f.Name)
		case "numberinput.Number":
			w = numberinput.New(f.Name)
		case "boxinput.Box":
			w = boxinput.New(f.Name)
		case "percentinput.Percent":
			w = percentinput.New(f.Name)
		case "menuinput.MenuInt":
			w = menuinput.NewMenuInt(f.Name)
		default:
			fmt.Println("Block:CreateFromDefinition() not handling:", fullTypeName, "as", typeName)
			continue
		}
		// N.B. run when a valid widget has been created - note the use of continue above
		widget.SetFields(w, string(tag))
		block.widgets = append(block.widgets, w)
		definition.SetupWidget(defn, f.Name)
	}
	if block.TypeName == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH EMPTY ("") BLOCK TYPE`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	if block.Class != "" {
		block.Class = "quando-" + block.Class // always insert quando- infront of class
	}
	return
}

func (block *Block) Op(early action.Early) {
	if early == nil {
		fmt.Printf("Warning: block type '%s' has nil operation\n", block.TypeName)
	}
	block.Early = early
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
		s, ok := w.(script.Generator) // ignore widgets that are purely visual, i.e. do not provide parameters
		if ok {
			if asHtml != "" { // separate parameters with comma
				asHtml += ","
			}
			asHtml += s.Generate()
		}
	}
	return asHtml
}
