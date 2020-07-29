package parameters

import (
	"strings"
)

//   WriteParameterToStrBuilder Name:  RELATED
//
//   Purpose:  To specify the relationship of the alarm trigger with
//      respect to the start or end of the calendar component.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//       trigrelparam       = "RELATED" "="
//                           ("START"       ; Trigger off of start
//                          / "END")        ; Trigger off of end
//
//   Description:  This parameter can be specified on properties that
//      specify an alarm trigger with a "DURATION" value type.  The
//      parameter specifies whether the alarm will trigger relative to the
//      start or end of the calendar component.  The parameter value START
//      will set the alarm to trigger off the start of the calendar
//      component; the parameter value END will set the alarm to trigger
//      off the end of the calendar component.  If the parameter is not
//      specified on an allowable property, then the default is START.
//
//   Example:
//
//       TRIGGER;RELATED=END:PT5M

type Related struct {
	V string
}

var StartRelated = Related{V: "START"}

var EndRelated = Related{V: "END"}

func (r *Related) WriteParameterToStrBuilder(s *strings.Builder) error {
	s.WriteString("RELATED=")
	s.WriteString(r.V)
	return nil
}
