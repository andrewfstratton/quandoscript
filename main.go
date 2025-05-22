package main

import (
	"fmt"

	"github.com/andrewfstratton/quandoscript/parse"
)

func main() {
	id, word, params, err := parse.Line("1 bob nl jill nl")
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(id, word, params)
}
