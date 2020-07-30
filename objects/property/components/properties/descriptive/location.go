package descriptive

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties"
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   Property Name:  LOCATION
//
//   Purpose:  This property defines the intended venue for the activity
//      defined by a calendar component.
//
//   V Type:  TEXT
//
//   Property Parameters:  IANA, non-standard, alternate text
//      representation, and language property parameters can be specified
//      on this property.
//
//   Conformance:  This property can be specified in "VEVENT" or "VTODO"
//      calendar component.
//
//   Description:  Specific venues such as conference or meeting rooms may
//      be explicitly specified using this property.  An alternate
//      representation may be specified that is a URI that points to
//      directory information with more structured specification of the
//      location.  For example, the alternate representation may specify
//      either an LDAP URL [RFC4516] pointing to an LDAP server entry or a
//      CID URL [RFC2392] pointing to a MIME body part containing a
//      Virtual-Information Card (vCard) [RFC2426] for the location.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       location   = "LOCATION"  locparam ":" text CRLF
//
//       locparam   = *(
//                  ;
//                  ; The following are OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  (";" altrepparam) / (";" languageparam) /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  (";" other-param)
//                  ;
//                  )
//
//   Example:  The following are some examples of this property:
//
//       LOCATION:Conference Room - F123\, Bldg. 002
//
//       LOCATION;ALTREP="http://xyzcorp.com/conf-rooms/f123.vcf":
//        Conference Room - F123\, Bldg. 002

type Location struct {
	Parameters []parameters.Parameter
	Values     *types.Text
}

func (l *Location) WritePropertyToStrBuilder(s *strings.Builder) error {
	return  properties.DefaultCreatePropertyFunc("LOCATION", l.Parameters, l.Values,s)
}

func NewLocation(location string) *Location {
	return &Location{
		Values: &types.Text{V: location},
	}
}
