package descriptive

import (
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
	"fmt"
	"testing"
)

func TestAttach_Properties(t *testing.T) {
	//ATTACH;FMTTYPE=application/postscript:ftp://example.com/pub/reports/r-960812.ps
	a := Attach{
		Parameters: []parameters.Parameter{&parameters.FmtType{V: "application/postscript"}},
		Value:      &types.URI{V: "ftp://example.com/pub/reports/r-960812.ps"},
	}
	ap, _ := a.Property()
	if ap != "ATTACH;FMTTYPE=application/postscript:ftp://example.com/pub/reports/r-960812.ps" {
		fmt.Println(a.Property())
		t.Error()
	}
	b := Attach{
		Value: &types.URI{V: "CID:jsmith.part3.960817T083000.xyzMail@example.com"},
	}
	bp, _ := b.Property()
	if bp != "ATTACH:CID:jsmith.part3.960817T083000.xyzMail@example.com" {
		t.Error()
	}

}
