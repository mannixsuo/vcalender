package descriptive

import (
	"calendar/objects/property/types"
	"testing"
)

func TestClassification_Property(t *testing.T) {
	c := Classification{
		Value: &types.Text{V: "PUBLIC"},
	}
	property, _ := c.Property()
	if property != "CLASS:PUBLIC" {
		t.Error()
	}
}
