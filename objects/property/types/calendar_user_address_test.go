package types

import "testing"

func TestCalAddress_Value(t *testing.T) {
	u := CalAddress{&URI{"jane_doe@example.com"}}
	if u.Value() != "mailto:jane_doe@example.com" {
		t.Error()
	}
}
