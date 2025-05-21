package main

import (
	"fmt"
	"time"

	"github.com/andrewfstratton/quandoscript/op"
	"github.com/andrewfstratton/quandoscript/parse"
	"github.com/andrewfstratton/quandoscript/run/param"
)

func log(prefix string) op.Op {
	return func() string {
		// see https://cs.opensource.google/go/go/+/go1.24.2:src/time/format.go;l=639
		fmt.Println(prefix + " " + time.Now().Format("15:04:05.00000"))
		return ""
	}
}

func init() {
	op.Add("log", log("Log"))
	// block := block.New("")
	// fmt.Println(block.Html())
}

func parseLine(line string) (id int, word string, params param.Params, err error) {
	id, word, params, err = parse.Line(line)
	return
}

func main() {
	id, word, params, err := parse.Line("1 bob nl jill nl")
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(id, word, params)
}

// result := `<div data-quando-block-type="` +
// 	block.qid +
// 	`" class="quando-block"` +
// 	//` data-quando-id="true"` + // Removed since always used now
// 	// Note that quandoscript needs unique block id (data-quando-id) at start of line + space when generating quandoscript call
// 	` data-quando-quandoscript='` + // single quote to allow quandoscript double quote string embedding
// 	block.qid + `(` +
// 	block.params() +
// 	`)'>` +
// 	`<div class="quando-left quando-` + block.class + `"></div>` +
// 	`<div class="quando-right">` +
// 	`<div class="quando-row quando-` + block.class + `">`
// for _, w := range block.widgets {
// 	result += w.Html()
// }
// result += `</div></div></div>`
// return result
