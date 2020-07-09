package types

import "testing"

func TestBINARY_Binary(t *testing.T) {
	var b Binary = Binary{V: []byte("abcdefghijklmnopqrstuvwxyz")}
	got := b.Value()
	if got != "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=" {
		t.Errorf("error")
	}
}
