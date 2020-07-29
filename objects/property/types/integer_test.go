package types

import (
	"strings"
	"testing"
)

func TestInteger_Value(t *testing.T) {
	s := &strings.Builder{}

	i := Integer{V: -1234567890}
	_ = i.WriteValueToStrBuilder(s)
	if s.String() != "-1234567890" {
		t.Error()
	}
}
