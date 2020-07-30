package changemanage

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties"
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
	"time"
)

//   Property Name:  CREATED
//
//   Purpose:  This property specifies the date and time that the calendar
//      information was created by the calendar user agent in the calendar
//      store.
//
//         Note: This is analogous to the creation date and time for a
//         file in the file system.
//
//   V Type:  DATE-TIME
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  The property can be specified once in "VEVENT",
//      "VTODO", or "VJOURNAL" calendar components.  The value MUST be
//      specified as a date with UTC time.
//
//   Description:  This property specifies the date and time that the
//      calendar information was created by the calendar user agent in the
//      calendar store.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       created    = "CREATED" creaparam ":" date-time CRLF
//
//       creaparam  = *(";" other-param)
//
//   Example:  The following is an example of this property:
//
//       CREATED:19960329T133000Z

type Created struct {
	Parameters []parameters.Parameter
	Value      *types.DateTime
}

func (c *Created) WritePropertyToStrBuilder(s *strings.Builder) error {
	return properties.DefaultCreatePropertyFunc("CREATED", c.Parameters, c.Value, s)
}

func NewCreated(year, month, day, hour, minute, seconds int) *Created {
	return &Created{
		Value: &types.DateTime{
			V:      time.Date(year, time.Month(month), day, hour, minute, seconds, 0, time.Local),
			Format: types.UTCDateTimeFormat,
		},
	}
}
