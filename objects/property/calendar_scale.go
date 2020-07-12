package property

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  CALSCALE
//
//   Purpose:  This property defines the calendar scale used for the
//      calendar information specified in the iCalendar object.
//
//   Value Type:  TEXT
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  This property can be specified once in an iCalendar
//      object.  The default value is "GREGORIAN".
//   Description:  This memo is based on the Gregorian calendar scale.
//      The Gregorian calendar scale is assumed if this property is not
//      specified in the iCalendar object.  It is expected that other
//      calendar scales will be defined in other specifications or by
//      future versions of this memo.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       calscale   = "CALSCALE" calparam ":" calvalue CRLF
//
//       calparam   = *(";" other-param)
//
//       calvalue   = "GREGORIAN"
//
//   Example:  The following is an example of this property:
//
//       CALSCALE:GREGORIAN

type CalendarScale struct {
	Parameters []parameters.Parameter
	Value      *types.Text
}

func (a *CalendarScale) Property() (string, error) {
	return properties.DefaultCreatePropertyFunc("CALSCALE", a.Parameters, a.Value), nil
}
