package main

import (
	"testing"

	"github.com/andrewfstratton/quandoscript/assert"
)

func TestParse(t *testing.T) {
	result, ok := parse("")
	assert.Eq(t, result, "")
	assert.Eq(t, ok, true)
}
