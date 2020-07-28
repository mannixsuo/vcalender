package timezone

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  TZOFFSETTO
//
//   Purpose:  This property specifies the offset that is in use in this
//      time zone observance.
//
//   V Type:  UTC-OFFSET
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  This property MUST be specified in "STANDARD" and
//      "DAYLIGHT" sub-components.
//
//   Description:  This property specifies the offset that is in use in
//      this time zone observance.  It is used to calculate the absolute
//      time for the new observance.  The property value is a signed
//      numeric indicating the number of hours and possibly minutes from
//      UTC.  Positive numbers represent time zones east of the prime
//      meridian, or ahead of UTC.  Negative numbers represent time zones
//      west of the prime meridian, or behind UTC.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       tzoffsetto = "TZOFFSETTO" toparam ":" utc-offset CRLF
//
//       toparam    = *(";" other-param)
//
//   Example:  The following are examples of this property:
//
//       TZOFFSETTO:-0400
//
//       TZOFFSETTO:+1245

type TzOffsetTo struct {
	Parameters []parameters.Parameter
	Value      *types.UTCOffset
}

func (t *TzOffsetTo) Property() (string, error) {
	return properties.DefaultCreatePropertyFunc("TZOFFSETTO", t.Parameters, t.Value), nil
}

func NewTzOffsetTo(positive bool, hour, minute, second int) *TzOffsetTo {
	return &TzOffsetTo{
		Value: types.NewUTCOffset(positive, hour, minute, second),
	}
}
