package main

import (
	"fmt"
	"strings"
)

type Op func()

var ops map[string]Op

func hi() {
	fmt.Println("hi")
}

func hello(name string) Op {
	return func() {
		fmt.Printf("Hello %s!", name)
	}
}

func init() {
	ops = make(map[string]Op)
	ops["hi"] = hi
	ops["bob"] = hello("Bob")
	ops["jill"] = hello("Jill")
	ops["nl"] = func() { fmt.Println() }
}

func parse(line string) (txt string, ok bool) {
	var result string
	for _, word := range strings.Fields(line) {
		op, ok := ops[word]
		if ok {
			op()
		}
	}
	return result, false
}

func main() {
	parse("hi bob nl jill nl")
}
