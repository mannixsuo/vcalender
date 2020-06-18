package model

import (
	"testing"
)

func TestBINARY_String(t *testing.T) {
	b := Binary{
		Binary: []byte("abcdefghijklmnopqrstuvwxyz"),
	}
	got := b.String()
	if got != ";ENCODING=BASE64;VALUE=BINARY:YWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXo=" {
		t.Errorf("error")
	}
}

func TestSplitString(t *testing.T) {
	s := "DESCRIPTION:This is a long description that exists on a long line.This is a long description that exists on a long line\n"
	right := "DESCRIPTION:This is a long description that exists on a long line.This\n" +
		"  is a long description that exists on a long line\n"
	if right != SplitString(s) {
		t.Error()
	}
	s = "DESCRIPTION:This is a short description."
	right = s
	if right != SplitString(s) {
		t.Error()
	}
	s = "DESCRIPTION:This is a long description that exists on a long line.This is a long description that exists on a long lineists on a long linee\n"
	right = "DESCRIPTION:This is a long description that exists on a long line.This\n" +
		"  is a long description that exists on a long lineists on a long linee\n"
	if right != SplitString(s) {
		t.Error()
	}
}

func TestBoolean_Boolean(t *testing.T) {
	var b Boolean = false
	if b.Boolean() != "FALSE" {
		t.Error()
	}
	b = true
	if b.Boolean() != "TRUE" {
		t.Error()
	}
}
