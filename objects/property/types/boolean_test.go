package types

import "testing"

func Test(t *testing.T) {
	b1 := Boolean{true}
	b2 := Boolean{false}
	if b1.Value() != "TRUE" || b2.Value() != "FALSE" {
		t.Error()
	}
}
