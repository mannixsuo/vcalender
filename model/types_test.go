package model

import (
	"log"
	"testing"
	"time"
)

func TestBINARY_String(t *testing.T) {
	var b Binary = []byte("abcdefghijklmnopqrstuvwxyz")
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

func TestAttach_Attach(t *testing.T) {
	cid := "ATTACH:CID:jsmith.part3.960817T083000.xyzMail@example.com"
	acid := Attach{
		URI: &URI{Uri: "CID:jsmith.part3.960817T083000.xyzMail@example.com"},
	}
	if acid.Attach() != cid {
		log.Print(acid.Attach())
		t.Error()
	}
	ps := "ATTACH;FMTTYPE=application/postscript:ftp://example.com/pub/reports/r-960812.ps"
	aps := Attach{
		FmtType: "application/postscript",
		URI:     &URI{Uri: "ftp://example.com/pub/reports/r-960812.ps"},
	}
	if ps != aps.Attach() {
		log.Print(aps.Attach())
		t.Error()
	}
}

func TestTrigger_Trigger(t *testing.T) {
	start := "TRIGGER:-PT15M"
	trigger := Trigger{
		Related2Start: true,
		Duration: &Duration{
			Negative:  true,
			DurMinute: 15,
		},
	}
	if start != trigger.Trigger() {
		log.Print(trigger.Trigger())
		t.Error()
	}
	end := "TRIGGER;RELATED=END:PT5M"
	trigger = Trigger{
		Related2Start: false,
		Duration: &Duration{
			DurMinute: 5,
		},
	}
	if end != trigger.Trigger() {
		log.Print(trigger.Trigger())
		t.Error()
	}
	dateTime := "TRIGGER;VALUE=DATE-TIME:19980101T050000Z"
	trigger = Trigger{
		DateTime: &DateTime{Date: time.Date(1998, 1, 1, 5, 0, 0, 0, time.Local)},
	}
	if dateTime != trigger.Trigger() {
		log.Print(trigger.Trigger())
		t.Error()
	}
}
