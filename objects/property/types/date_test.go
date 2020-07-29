package types

import (
	"strings"
	"testing"
	"time"
)

func TestDate_Date(t *testing.T) {
	s := &strings.Builder{}

	d := Date{time.Date(2020, 7, 7, 0, 0, 0, 0, time.Local)}
	d.WriteValueToStrBuilder(s)
	if s.String() != "20200707" {
		t.Error()
	}
}
