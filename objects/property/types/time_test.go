package types

import (
	"fmt"
	"testing"
	"time"
)

func TestTime_Value(t *testing.T) {

	ti := Time{
		V:      time.Date(0, 0, 0, 10, 52, 11, 0, time.Local),
		Format: UTCTimeFormat,
	}
	if ti.Value() != "105211Z" {
		fmt.Print(ti.Value())
		t.Error()
	}
	ti.Format = LocalTimeFormat

	if ti.Value() != "105211" {
		fmt.Print(ti.Value())
		t.Error()
	}
}
