package datetime

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  DUE
//
//   Purpose:  This property defines the date and time that a to-do is
//      expected to be completed.
//
//   V Type:  The default value type is DATE-TIME.  The value type can
//      be set to a DATE value type.
//
//   Property Parameters:  IANA, non-standard, value data type, and time
//      zone identifier property parameters can be specified on this
//      property.
//
//   Conformance:  The property can be specified once in a "VTODO"
//      calendar component.
//
//   Description:  This property defines the date and time before which a
//      to-do is expected to be completed.  For cases where this property
//      is specified in a "VTODO" calendar component that also specifies a
//      "DTSTART" property, the value type of this property MUST be the
//      same as the "DTSTART" property, and the value of this property
//      MUST be later in time than the value of the "DTSTART" property.
//      Furthermore, this property MUST be specified as a date with local
//      time if and only if the "DTSTART" property is also specified as a
//      date with local time.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       due        = "DUE" dueparam ":" dueval CRLF
//
//       dueparam   = *(
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
//       dueval     = date-time / date
//       ;V MUST match value type
//
//   Example:  The following is an example of this property:
//
//       DUE:19980430T000000Z

type Due struct {
	Parameters []parameters.Parameter
	Value      types.Value
}

func (d *Due) Property() (string, error) {
	return properties.DefaultCreatePropertyFunc("DUE", d.Parameters, d.Value), nil
}

func NewDueWithDate(year, month, day int) *Due {
	return &Due{
		Parameters: []parameters.Parameter{&parameters.Date},
		Value:      types.NewDate(year, month, day),
	}
}
func NewDueWithDateTime(year, month, day, hour, minute, seconds int) *Due {
	return &Due{
		Parameters: []parameters.Parameter{&parameters.DateTime},
		Value:      types.NewUTCDateTime(year, month, day, hour, minute, seconds),
	}
}
