package assert

import (
	"testing"
)

// from https://dev.to/yawaramin/why-i-dont-use-a-third-party-assertion-library-in-go-unit-tests-1gak
func Eq[V comparable](t *testing.T, got, expected V) {
	t.Helper()

	if expected != got {
		t.Errorf("assert.Eq(got:%v, expected:%v)", got, expected)
	}
}
