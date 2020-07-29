package types

import (
	"strings"
	"testing"
	"time"
)

func TestDateTime_Value(t *testing.T) {
	s := &strings.Builder{}

	d := DateTime{
		V:      time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local),
		Format: UTCDateTimeFormat,
	}
	d.WriteValueToStrBuilder(s)
	if s.String() != "20060102T150405Z" {
		t.Error()
	}
	s.Reset()
	d.Format = LocalDateTimeFormat
	d.WriteValueToStrBuilder(s)
	if s.String() != "20060102T150405" {
		t.Error()
	}
}
