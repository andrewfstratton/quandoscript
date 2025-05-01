package main

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/op"
)

func TestAddEmptyStringWordOp(t *testing.T) {
	err := op.Add("", func() string { return "Should never allow map to empty string" })
	assert.Neq(t, err, nil)
}

func TestAddWordOp(t *testing.T) {
	err := op.Add("word", func() string { return "word string" })
	assert.Eq(t, err, nil)
}

func TestParseEmpty(t *testing.T) {
	fn, err := parseLine("")
	assert.Eq(t, err, nil)
	assert.True(t, fn == nil)
}

func TestParseMissing(t *testing.T) {
	fn, err := parseLine("_")
	assert.Neq(t, err, nil)
	assert.True(t, fn == nil)
}

func TestParseWhiteSpace(t *testing.T) {
	fn, err := parseLine(" \n\t")
	assert.Eq(t, err, nil)
	assert.True(t, fn == nil)
}

func TestLogStandard(t *testing.T) {
	fn, err := parseLine("log")
	assert.Eq(t, err, nil)
	assert.True(t, fn != nil)
	fn()
}

func init() {
	// first below shouldn't be added
	op.Add("", func() string { return "Should never allow map to empty string" })
	op.Add("hi", func() string { return "hi" })
	op.Add("nl", func() string { return "\n" })
}
