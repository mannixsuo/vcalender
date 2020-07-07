package types

import "testing"

func Test(t *testing.T) {
	var b Boolean
	b = true
	if b.Boolean() != "TRUE" {
		t.Error()
	}
	b = false
	if b.Boolean() != "FALSE" {
		t.Error()
	}
}
