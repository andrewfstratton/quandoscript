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
)

type Block struct {
	TypeName string
	Class    string
	widgets  []widget.Widget
	Early    action.Early
}

func CreateFromDefinition(defn any) (block *Block) {
	// 	N.B. TypeName and Class exist in Defn - not in widgets
	block = &Block{}
	block.setupWidgets(defn)
	if block.TypeName == "" {
		fmt.Println(`ATTEMPT TO CREATE BLOCK WITH EMPTY ("") BLOCK TYPE`)
		if testing.Testing() {
			return nil
		}
		debug.PrintStack()
		os.Exit(99)
	}
	return
}

func (block *Block) setupWidgets(defn any) {
	typeDefn := reflect.TypeOf(defn)
	for i := range typeDefn.NumField() {
		field := typeDefn.Field(i)
		tag := field.Tag
		underscoreTag := tag.Get("_")
		if field.Name == "TypeName" {
			block.TypeName = underscoreTag
			continue
		}
		if field.Name == "Class" && underscoreTag != "" {
			block.Class = "quando-" + underscoreTag // always insert quando- infront of class
			continue
		}
		// Otherwise check the type
		var w widget.Widget
		fullTypeName := field.Type.PkgPath() + "." + field.Type.Name()
		if fullTypeName == "." {
			fmt.Printf("setupWidgets field '%s' has no type -- may be pointer", field.Name)
		}
		_, typeName := path.Split(fullTypeName) // split on last /
		switch typeName {
		case "text.Text":
			w = text.New()
		case "stringinput.String":
			w = stringinput.New(field.Name)
		case "numberinput.Number":
			w = numberinput.New(field.Name)
		case "boxinput.Box":
			w = boxinput.New(field.Name, block.Class)
		case "percentinput.Percent":
			w = percentinput.New(field.Name)
		case "menuinput.MenuInt":
			w = menuinput.NewMenuInt(field.Name)
		case "menuinput.MenuStr":
			w = menuinput.NewMenuStr(field.Name)
		default:
			fmt.Println("Block:setupWidgets() not handling:", fullTypeName, "as", typeName, "with kind", field.Type.Kind())
			continue
		}
		// N.B. below runs when a valid widget has been created - note the use of continue above
		widget.SetFields(w, string(tag))
		block.widgets = append(block.widgets, w)
	}
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
