package types

import (
	"strings"
	"testing"
)

func TestDuration_Value(t *testing.T) {
	s := &strings.Builder{}
	// P15DT5H0M20S
	d := Duration{
		Negative:  false,
		DurWeek:   0,
		DurDay:    15,
		DurHour:   5,
		DurMinute: 0,
		DurSecond: 20,
	}
	d.WriteValueToStrBuilder(s)
	if s.String() != "P15DT5H0M20S" {
		t.Error()
	}
	// P7W
	w := Duration{
		Negative:  false,
		DurWeek:   7,
		DurDay:    0,
		DurHour:   0,
		DurMinute: 0,
		DurSecond: 0,
	}
	s.Reset()
	w.WriteValueToStrBuilder(s)
	if s.String() != "P7W" {
		t.Error()
	}
}
