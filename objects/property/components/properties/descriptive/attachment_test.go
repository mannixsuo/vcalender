package descriptive

import (
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"fmt"
	"strings"
	"testing"
)

func TestAttach_Properties(t *testing.T) {
	//ATTACH;FMTTYPE=application/postscript:ftp://example.com/pub/reports/r-960812.ps
	s := &strings.Builder{}
	a := Attach{
		Parameters: []parameters.Parameter{&parameters.FmtType{V: "application/postscript"}},
		Value:      &types.URI{V: "ftp://example.com/pub/reports/r-960812.ps"},
	}
	_ = a.WritePropertyToStrBuilder(s)
	if s.String() != "ATTACH;FMTTYPE=application/postscript:ftp://example.com/pub/reports/r-960812.ps" {
		fmt.Println(s.String())
		t.Error()
	}
	b := Attach{
		Value: &types.URI{V: "CID:jsmith.part3.960817T083000.xyzMail@example.com"},
	}
	s.Reset()
	_ = b.WritePropertyToStrBuilder(s)
	if s.String() != "ATTACH:CID:jsmith.part3.960817T083000.xyzMail@example.com" {
		t.Error()
	}

}
