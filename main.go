package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Op func() string

var ops map[string]Op

func log(prefix string) Op {
	return func() string {
		// see https://cs.opensource.google/go/go/+/go1.24.2:src/time/format.go;l=639
		fmt.Println(prefix + " " + time.Now().Format("15:04:05.00000"))
		return ""
	}
}

func init() {
	ops = make(map[string]Op)
	addOp("log", log("Log"))
}

func parseCall(call string) (Op, error) {
	var result Op
	var err error
	if call != "" {
		op, ok := ops[call]
		if ok {
			result = op
		} else {
			err = errors.New("Failed to parse call '" + call + "'")
		}
	}
	return result, err
}

func parseLine(line string) (string, error) {
	var result string
	var err error
	for _, word := range strings.Fields(line) {
		var op Op
		op, err = parseCall(word)
		if err == nil {
			result += op()
		} else {
			break // i.e. bail out early
		}
	}
	return result, err
}

func addOp(lookup string, op Op) (err error) {
	if lookup == "" {
		err = errors.New("ignoring new operation with empty lookup")
	} else if op == nil {
		err = errors.New("ignoring nil operation for '" + lookup + "'")
	} else {
		ops[lookup] = op
	}
	return err
}

func main() {
	result, err := parseLine("hi bob nl jill nl")
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(result)
}
