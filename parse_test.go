package main

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestAddEmptyStringWordOp(t *testing.T) {
	err := addOp("", func() string { return "Should never allow map to empty string" })
	assert.Neq(t, err, nil)
}

func TestParse(t *testing.T) {
	op, err := parseCall("")
	assert.Eq(t, err, nil)
	assert.True(t, op == nil)
}

func TestParseMissing(t *testing.T) {
	op, err := parseCall("_")
	assert.Neq(t, err, nil)
	assert.True(t, op == nil)
}

func TestParseWhiteSpace(t *testing.T) {
	txt, err := parseLine(" ")
	assert.Eq(t, err, nil)
	assert.Eq(t, txt, "")
}

func init() {
	// first below shouldn't be added
	addOp("", func() string { return "Should never allow map to empty string" })
	addOp("hi", func() string { return "hi" })
	addOp("nl", func() string { return "\n" })
}
