package types

import (
	"strings"
	"testing"
)

func TestFloat_Value(t *testing.T) {
	s := &strings.Builder{}

	f := Float{V: "1000000.0000001"}
	f.WriteValueToStrBuilder(s)
	if s.String() != "1000000.0000001" {
		t.Error()
	}
}
