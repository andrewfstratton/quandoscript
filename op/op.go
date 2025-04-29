package op

import (
	"errors"
)

type Op func() string

var ops map[string]Op

func Add(lookup string, op Op) (err error) {
	if lookup == "" {
		err = errors.New("ignoring new operation with empty lookup")
	} else if op == nil {
		err = errors.New("ignoring nil operation for '" + lookup + "'")
	} else {
		ops[lookup] = op
	}
	return err
}

func Get(word string) (fn Op, err error) {
	if word == "" {
		err = errors.New("empty string function lookup - ignored")
	} else {
		var ok bool
		fn, ok = ops[word]
		if !ok {
			err = errors.New("Failed to find operation '" + word + "'")
		}
	}
	return fn, err
}

func init() {
	ops = make(map[string]Op)
}
