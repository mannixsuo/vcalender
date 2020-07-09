package types

import "testing"

func TestInteger_Value(t *testing.T) {
	i := Integer{V: -1234567890}
	if i.Value() != "-1234567890" {
		t.Error()
	}
}
