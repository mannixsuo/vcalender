package types

import (
	"strings"
	"testing"
)

func TestCalAddress_Value(t *testing.T) {
	s := &strings.Builder{}

	u := CalAddress{&URI{"jane_doe@example.com"}}
	u.WriteValueToStrBuilder(s)
	if s.String() != "mailto:jane_doe@example.com" {
		t.Error()
	}
}
