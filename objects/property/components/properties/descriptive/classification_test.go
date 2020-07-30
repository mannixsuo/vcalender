package descriptive

import (
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
	"testing"
)

func TestClassification_Property(t *testing.T) {
	s := &strings.Builder{}

	c := Classification{
		Value: &types.Text{V: "PUBLIC"},
	}
	_ = c.WritePropertyToStrBuilder(s)
	if s.String() != "CLASS:PUBLIC" {
		t.Error()
	}
}
