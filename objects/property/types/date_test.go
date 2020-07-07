package types

import (
	"testing"
)

func TestDate_Date(t *testing.T) {
	d := NewDate(2020, 7, 7).Date()
	if d != "20200707" {
		t.Error()
	}
}
