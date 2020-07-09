package descriptive

import (
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
	"testing"
)

func TestAttach_Properties(t *testing.T) {
	a := Attach{
		Parameters: parameters.NewParameters(parameters.NewFmtType("application/postscript")),
		Value:      &types.URI{V: "ftp://example.com/pub/reports/r-960812.ps"},
	}
	if a.Properties() != "ATTACH;FMTTYPE=application/postscript:ftp://example.com/pub/reports/r-960812.ps" {
		t.Error()
	}
	b := Attach{
		Parameters: parameters.NewParameters(parameters.NewValueType(parameters.Binary),
			parameters.NewEncoding(parameters.BASE64)),
		Value: &types.Binary{V: []byte("abcdefghijklmnopqrstuvwxyz")},
	}
}
