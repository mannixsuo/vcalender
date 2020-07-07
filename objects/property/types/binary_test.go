package types

import "testing"

func TestBINARY_Binary(t *testing.T) {
	var b Binary = []byte("abcdefghijklmnopqrstuvwxyz")
	got := b.Binary()
	if got != "YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=" {
		t.Errorf("error")
	}
}
