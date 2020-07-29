package parameters

import (
	"fmt"
	"strings"
)

//   WriteParameterToStrBuilder Name:  VALUE
//
//   Purpose:  To explicitly specify the value type format for a property
//      value.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//       valuetypeparam = "VALUE" "=" valuetype
//
//       valuetype  = ("BINARY"
//                  / "BOOLEAN"
//                  / "CAL-ADDRESS"
//                  / "DATE"
//                  / "DATE-TIME"
//                  / "DURATION"
//                  / "FLOAT"
//                  / "INTEGER"
//                  / "PERIOD"
//                  / "RECUR"
//                  / "TEXT"
//                  / "TIME"
//                  / "URI"
//                  / "UTC-OFFSET"
//                  / x-name
//                  ; Some experimental iCalendar value type.
//                  / iana-token)
//                  ; Some other IANA-registered iCalendar value type.
//
//   Description:  This parameter specifies the value type and format of
//      the property value.  The property values MUST be of a single value
//      type.  For example, a "RDATE" property cannot have a combination
//      of DATE-TIME and TIME value types.
//
//      If the property's value is the default value type, then this
//      parameter need not be specified.  However, if the property's
//      default value type is overridden by some other allowable value
//      type, then this parameter MUST be specified.
//
//      Applications MUST preserve the value data for x-name and iana-
//      token values that they don't recognize without attempting to
//      interpret or parse the value data.

type ValueType struct {
	V string
}

var (
	Binary     = ValueType{"BINARY"}
	Boolean    = ValueType{"BOOLEAN"}
	CalAddress = ValueType{"CAL-ADDRESS"}
	Date       = ValueType{"DATE"}
	DateTime   = ValueType{"DATE-TIME"}
	Duration   = ValueType{"DURATION"}
	Float      = ValueType{"FLOAT"}
	Integer    = ValueType{"INTEGER"}
	Period     = ValueType{"PERIOD"}
	Recur      = ValueType{"RECUR"}
	Text       = ValueType{"TEXT"}
	Time       = ValueType{"TIME"}
	Uri        = ValueType{"URI"}
	UtcOffset  = ValueType{"UTC-OFFSET"}
)

func (v *ValueType) WriteParameterToStrBuilder(s *strings.Builder) error {
	s.WriteString(fmt.Sprintf("VALUE=%s", v.V))
	return nil
}
