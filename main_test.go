package main

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
	"github.com/andrewfstratton/quandoscript/op"
	"github.com/andrewfstratton/quandoscript/parse"
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
	id, word, params, err := parseLine("")
	assert.Eq(t, err, nil)
	assert.Eq(t, id, 0)
	assert.Eq(t, word, "")
	assert.True(t, params == nil)
}

func TestParseMissing(t *testing.T) {
	id, word, params, err := parseLine("_")
	assert.Neq(t, err, nil)
	assert.Eq(t, id, 0)
	assert.Eq(t, word, "")
	assert.True(t, params == nil)
}

func TestParseWhiteSpace(t *testing.T) {
	id, word, params, err := parseLine(" \n\t")
	assert.Neq(t, err, nil)
	assert.Eq(t, id, 0)
	assert.Eq(t, word, "")
	assert.True(t, params == nil)
}

func TestLogNoParams(t *testing.T) {
	id, word, params, err := parseLine("12 log()")
	assert.Eq(t, err, nil)
	assert.Eq(t, id, 12)
	assert.Eq(t, word, "log")
	assert.Eq(t, len(params), 0)
}

func TestLogParams(t *testing.T) {
	id, word, params, err := parseLine(`12 log(v#12,message"hello")`)
	assert.Eq(t, err, nil)
	assert.Eq(t, id, 12)
	assert.Eq(t, word, "log")
	assert.Eq(t, len(params), 2)
	assert.Eq(t, params["v"].Qtype, parse.NUMBER)
	assert.Eq(t, params["v"].Val, 12.0)
	assert.Eq(t, params["message"].Qtype, parse.STRING)
	assert.Eq(t, params["message"].Val, "hello")
}

func init() {
	// first below shouldn't be added
	op.Add("", func() string { return "Should never allow map to empty string" })
	op.Add("hi", func() string { return "hi" })
	op.Add("nl", func() string { return "\n" })
}
