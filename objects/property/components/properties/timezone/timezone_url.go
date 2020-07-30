package timezone

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties"
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   Property Name:  TZURL
//
//   Purpose:  This property provides a means for a "VTIMEZONE" component
//      to point to a network location that can be used to retrieve an up-
//      to-date version of itself.
//
//   V Type:  URI
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  This property can be specified in a "VTIMEZONE"
//      calendar component.
//
//   Description:  This property provides a means for a "VTIMEZONE"
//      component to point to a network location that can be used to
//      retrieve an up-to-date version of itself.  This provides a hook to
//      handle changes government bodies impose upon time zone
//      definitions.  Retrieval of this resource results in an iCalendar
//      object containing a single "VTIMEZONE" component and a "METHOD"
//      property set to PUBLISH.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       tzurl      = "TZURL" tzurlparam ":" uri CRLF
//
//       tzurlparam = *(";" other-param)
//
//   Example:  The following is an example of this property:
//
//    TZURL:http://timezones.example.org/tz/America-Los_Angeles.ics

type TzUrl struct {
	Parameters []parameters.Parameter
	Value      *types.URI
}

func (t *TzUrl) WritePropertyToStrBuilder(s *strings.Builder) error {
	return properties.DefaultCreatePropertyFunc("TZURL", t.Parameters, t.Value, s)
}
