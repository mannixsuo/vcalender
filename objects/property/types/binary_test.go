package types

import (
	"strings"
	"testing"
)

func TestBINARY_Binary(t *testing.T) {
	s := &strings.Builder{}
	var b Binary = Binary{V: []byte("abcdefghijklmnopqrstuvwxyz")}
	_ = b.WriteValueToStrBuilder(s)
	if s.String() != "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=" {
		t.Errorf("error")
	}
}
