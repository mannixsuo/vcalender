package datetime

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties"
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   Property Name:  DTEND
//
//   Purpose:  This property specifies the date and time that a calendar
//      component ends.
//
//   V Type:  The default value type is DATE-TIME.  The value type can
//      be set to a DATE value type.
//
//   Property Parameters:  IANA, non-standard, value data type, and time
//      zone identifier property parameters can be specified on this
//      property.
//
//   Conformance:  This property can be specified in "VEVENT" or
//      "VFREEBUSY" calendar components.
//
//   Description:  Within the "VEVENT" calendar component, this property
//      defines the date and time by which the event ends.  The value type
//      of this property MUST be the same as the "DTSTART" property, and
//      its value MUST be later in time than the value of the "DTSTART"
//      property.  Furthermore, this property MUST be specified as a date
//      with local time if and only if the "DTSTART" property is also
//      specified as a date with local time.
//
//      Within the "VFREEBUSY" calendar component, this property defines
//      the end date and time for the free or busy time information.  The
//      time MUST be specified in the UTC time format.  The value MUST be
//      later in time than the value of the "DTSTART" property.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       dtend      = "DTEND" dtendparam ":" dtendval CRLF
//
//       dtendparam = *(
//                  ;
//                  ; The following are OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  (";" "VALUE" "=" ("DATE-TIME" / "DATE")) /
//                  (";" tzidparam) /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  (";" other-param)
//                  ;
//                  )
//
//       dtendval   = date-time / date
//       ;V MUST match value type
//
//   Example:  The following is an example of this property:
//
//       DTEND:19960401T150000Z
//
//       DTEND;VALUE=DATE:19980704

type DateEnd struct {
	Parameters []parameters.Parameter
	Value      types.Value
}

func (d *DateEnd) WritePropertyToStrBuilder(s *strings.Builder) error {
	return properties.DefaultCreatePropertyFunc("DTEND", d.Parameters, d.Value, s)
}

func NewDateTimeDateEnd(year, month, day, hour, minute, seconds int) *DateEnd {
	return &DateEnd{
		Value: types.NewUTCDateTime(year, month, day, hour, minute, seconds),
	}
}
