package types

import (
	"testing"
	"time"
)

func TestDateTime_Value(t *testing.T) {
	d := DateTime{
		V:      time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local),
		Format: UTCDateTimeFormat,
	}
	if d.Value() != "20060102T150405Z" {
		t.Error()
	}
	d.Format = LocalDateTimeFormat
	if d.Value() != "20060102T150405" {
		t.Error()
	}
}
