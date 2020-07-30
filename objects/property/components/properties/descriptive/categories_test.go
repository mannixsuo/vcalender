package descriptive

import (
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
	"testing"
)

func TestCategories_Property(t *testing.T) {
	s := &strings.Builder{}
	c := Categories{
		Parameters: []parameters.Parameter{&parameters.Language{V: "en"}},
		Values:     []types.Value{&types.Text{V: "APPOINTMENT"}, &types.Text{V: "EDUCATION"}},
	}
	_ = c.WritePropertyToStrBuilder(s)
	if s.String() != "CATEGORIES;LANGUAGE=en:APPOINTMENT,EDUCATION" {
		t.Error()
	}
	c2 := Categories{
		Parameters: nil,
		Values:     []types.Value{&types.Text{V: "MEETING"}},
	}
	s.Reset()
	_ = c2.WritePropertyToStrBuilder(s)
	if s.String() != "CATEGORIES:MEETING" {
		t.Error()
	}
}
