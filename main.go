package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/andrewfstratton/quandoscript/op"
	"github.com/andrewfstratton/quandoscript/parse"
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
}

func parseLine(line string) (fn op.Op, err error) {
	for _, word := range strings.Fields(line) {
		fn, err = op.Get(word)
		if err != nil {
			break // i.e. bail out early
		}
	}
	return fn, err
}

func main() {
	fn, err := parse.Line("1 bob nl jill nl")
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(fn)
}
