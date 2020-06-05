package model

import (
	"testing"
)

func TestBINARY_String(t *testing.T) {
	b := Binary{
		Data: []byte("abcdefghijklmnopqrstuvwxyz"),
	}
	got := b.String()
	if got != ";ENCODING=BASE64;VALUE=BINARY:YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=" {
		t.Errorf("error")
	}
}
