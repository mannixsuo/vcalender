package types

import (
	"testing"
	"time"
)

func TestDate_Date(t *testing.T) {
	d := Date{time.Date(2020,7, 7, 0, 0, 0, 0, time.Local)}
	if d.Value() != "20200707" {
		t.Error()
	}
}
