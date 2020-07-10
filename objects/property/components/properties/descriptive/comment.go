package descriptive

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  COMMENT
//
//   Purpose:  This property specifies non-processing information intended
//      to provide a comment to the calendar user.
//
//   Value Type:  TEXT
//
//   Property Parameters:  IANA, non-standard, alternate text
//      representation, and language property parameters can be specified
//      on this property.
//   Conformance:  This property can be specified multiple times in
//      "VEVENT", "VTODO", "VJOURNAL", and "VFREEBUSY" calendar components
//      as well as in the "STANDARD" and "DAYLIGHT" sub-components.
//
//   Description:  This property is used to specify a comment to the
//      calendar user.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       comment    = "COMMENT" commparam ":" text CRLF
//
//       commparam  = *(
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
//   Example:  The following is an example of this property:
//
//       COMMENT:The meeting really needs to include both ourselves
//         and the customer. We can't hold this meeting without them.
//         As a matter of fact\, the venue for the meeting ought to be at
//         their site. - - John
type Comment struct {
	Parameters []parameters.Parameter
	Value      types.Value
}

func (c *Comment) Property() (string, error) {
	return properties.DefaultCreatePropertyFunc("COMMENT", c.Parameters, c.Value), nil
}
