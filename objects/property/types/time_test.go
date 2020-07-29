package types

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTime_Value(t *testing.T) {
	b := &strings.Builder{}

	ti := Time{
		V:      time.Date(0, 0, 0, 10, 52, 11, 0, time.Local),
		Format: UTCTimeFormat,
	}
	ti.WriteValueToStrBuilder(b)
	if b.String() != "105211Z" {
		fmt.Print(b.String())
		t.Error()
	}
	ti.Format = LocalTimeFormat
	b.Reset()
	ti.WriteValueToStrBuilder(b)

	if b.String() != "105211" {
		fmt.Print(b.String())
		t.Error()
	}
}
