package descriptive

import (
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
	"fmt"
	"testing"
)

func TestCategories_Property(t *testing.T) {
	c := Categories{
		Parameters: []parameters.Parameter{&parameters.Language{V: "en"}},
		Values:     []types.Value{&types.Text{V: "APPOINTMENT"}, &types.Text{V: "EDUCATION"}},
	}
	p, _ := c.Property()
	if p != "CATEGORIES;LANGUAGE=en:APPOINTMENT,EDUCATION" {
		fmt.Println(p)
		t.Error()
	}
	c2 := Categories{
		Parameters: nil,
		Values:     []types.Value{&types.Text{V: "MEETING"}},
	}
	property, _ := c2.Property()
	if property != "CATEGORIES:MEETING" {
		fmt.Println(property)
		t.Error()
	}
}
