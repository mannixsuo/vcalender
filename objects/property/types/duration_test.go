package types

import (
	"fmt"
	"testing"
)

func TestDuration_Value(t *testing.T) {
	// P15DT5H0M20S
	d := Duration{
		Negative:  false,
		DurWeek:   0,
		DurDay:    15,
		DurHour:   5,
		DurMinute: 0,
		DurSecond: 20,
	}
	if d.Value() != "P15DT5H0M20S" {
		fmt.Print(d.Value())
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
	if w.Value() != "P7W" {
		t.Error()
	}
}
