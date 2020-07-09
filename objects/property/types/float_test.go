package types

import (
	"fmt"
	"testing"
)

func TestFloat_Value(t *testing.T) {
	f := Float{V: "1000000.0000001"}
	if f.Value() != "1000000.0000001" {
		fmt.Println(f.Value())
		t.Error()
	}
}
