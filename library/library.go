package library

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"

	"github.com/andrewfstratton/quandoscript/action"
	"github.com/andrewfstratton/quandoscript/action/param"
	"github.com/andrewfstratton/quandoscript/block"
	"github.com/andrewfstratton/quandoscript/menu"
	"github.com/andrewfstratton/quandoscript/parse"
)

const (
	SERVER  = "server"
	UNKNOWN = ""
)

var blocks map[string]*block.Block // lookup for all block types
var menus map[string]*menu.Menu    // groups of blocks by 'class' for menu
var classes []string

func Block(defn any) (b *block.Block) { // call through block New
	b = block.CreateFromDefinition(defn)
	_, inuse := blocks[b.TypeName]
	if inuse {
		fmt.Println(`BLOCK "` + b.TypeName + `" ALREADY EXISTS`)
		if testing.Testing() {
			return
		}
		debug.PrintStack()
		os.Exit(99)
	}
	blocks[b.TypeName] = b
	bl, found := menus[b.Class]
	if !found {
		bl = menu.New(b.Class)
		menus[b.Class] = bl
		classes = append(classes, b.Class)
	}
	bl.Add(b)
	return
}

func FindBlock(block_type string) (block *block.Block, found bool) {
	block, found = blocks[block_type]
	return
}

func newAction(word string, early param.Params, late_params param.Params) *action.Action {
	bt, found := FindBlock(word)
	if !found {
		fmt.Printf("Error : Lookup failing for : '%s'\n", word)
		return nil
	}
	late := bt.Early(early)              // run the early binding
	return action.New(late, late_params) // return the late binding with the closure
}

func Classes() []string {
	return classes
}

func Menus(class string) *menu.Menu {
	return menus[class]
}

func Parse(lines string) { // setup the whole script as actions for calling - only does early setup
	parse.Lines(lines, newAction)
}

func init() {
	blocks = make(map[string]*block.Block)
	menus = make(map[string]*menu.Menu)
}
