package relationship

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties"
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   Property Name:  URL
//
//   Purpose:  This property defines a Uniform Resource Locator (URL)
//      associated with the iCalendar object.
//
//   V Type:  URI
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  This property can be specified once in the "VEVENT",
//      "VTODO", "VJOURNAL", or "VFREEBUSY" calendar components.
//
//   Description:  This property may be used in a calendar component to
//      convey a location where a more dynamic rendition of the calendar
//      information associated with the calendar component can be found.
//      This memo does not attempt to standardize the form of the URI, nor
//      the format of the resource pointed to by the property value.  If
//      the URL property and Content-Location MIME header are both
//      specified, they MUST point to the same resource.
//   Format Definition:  This property is defined by the following
//      notation:
//
//       url        = "URL" urlparam ":" uri CRLF
//
//       urlparam   = *(";" other-param)
//
//   Example:  The following is an example of this property:
//
//       URL:http://example.com/pub/calendars/jsmith/mytime.ics

type Url struct {
	Parameters []parameters.Parameter
	Value      *types.URI
}

func (r *Url) WritePropertyToStrBuilder(s *strings.Builder) error {
	return properties.DefaultCreatePropertyFunc("URL", r.Parameters, r.Value,s)
}

func NewUrl(url string) *Url {
	return &Url{
		Value: types.NewUri(url),
	}
}
