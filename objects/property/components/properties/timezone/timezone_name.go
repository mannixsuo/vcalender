package timezone

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties"
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   Property Name:  TZNAME
//
//   Purpose:  This property specifies the customary designation for a
//      time zone description.
//
//   V Type:  TEXT
//
//   Property Parameters:  IANA, non-standard, and language property
//      parameters can be specified on this property.
//
//   Conformance:  This property can be specified in "STANDARD" and
//      "DAYLIGHT" sub-components.
//
//   Description:  This property specifies a customary name that can be
//      used when displaying dates that occur during the observance
//      defined by the time zone sub-component.
//
//   Format Definition:  This property is defined by the following
//      notation:
//       tzname     = "TZNAME" tznparam ":" text CRLF
//
//       tznparam   = *(
//                  ;
//                  ; The following is OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  (";" languageparam) /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  (";" other-param)
//                  ;
//                  )
//
//   Example:  The following are examples of this property:
//
//       TZNAME:EST
//
//       TZNAME;LANGUAGE=fr-CA:HNE

type TzName struct {
	Parameters []parameters.Parameter
	Value      *types.Text
}

func (t *TzName) WritePropertyToStrBuilder(s *strings.Builder) error {
	return properties.DefaultCreatePropertyFunc("TZNAME", t.Parameters, t.Value,s)
}

func NewTzName(name string) *TzName {
	return &TzName{
		Value: types.NewText(name),
	}
}
