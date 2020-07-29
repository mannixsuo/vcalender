package types

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	s := &strings.Builder{}
	b1 := Boolean{true}
	_ = b1.WriteValueToStrBuilder(s)
	if s.String() != "TRUE" {
		t.Error()
	}
	b1.V = false
	s.Reset()
	_ = b1.WriteValueToStrBuilder(s)
	if s.String() != "FALSE" {
		t.Error()
	}
}
