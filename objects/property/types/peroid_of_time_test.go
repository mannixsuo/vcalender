package types

import (
	"strings"
	"testing"
	"time"
)

func TestExplicitPeriod_Value(t *testing.T) {
	s := &strings.Builder{}

	//19970101T180000Z/19970102T070000Z
	e := ExplicitPeriod{
		Start: &DateTime{
			V:      time.Date(1997, 1, 1, 18, 0, 0, 0, time.Local),
			Format: UTCDateTimeFormat,
		},
		End: &DateTime{
			V:      time.Date(1997, 1, 2, 7, 0, 0, 0, time.Local),
			Format: UTCDateTimeFormat,
		},
	}
	e.WriteValueToStrBuilder(s)
	if s.String() != "19970101T180000Z/19970102T070000Z" {
		t.Error()
	}
}

func TestStartPeriod_Value(t *testing.T) {
	b := &strings.Builder{}

	//19970101T180000Z/PT5H30M
	s := StartPeriod{
		Start: &DateTime{
			V:      time.Date(1997, 1, 1, 18, 0, 0, 0, time.Local),
			Format: UTCDateTimeFormat,
		},
		Duration: &Duration{
			DurHour:   5,
			DurMinute: 30,
		},
	}
	s.WriteValueToStrBuilder(b)
	if b.String() != "19970101T180000Z/PT5H30M" {
		t.Errorf("value:%s != 19970101T180000Z/PT5H30M", b.String())
	}

}
