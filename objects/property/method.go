package property

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties"
	"github.com/mmsuo/vcalender/objects/property/parameters"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   Property Name:  METHOD
//
//   Purpose:  This property defines the iCalendar object method
//      associated with the calendar object.
//
//   V Type:  TEXT
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  This property can be specified once in an iCalendar
//      object.
//
//   Description:  When used in a MIME message entity, the value of this
//      property MUST be the same as the Content-Type "method" parameter
//      value.  If either the "METHOD" property or the Content-Type
//      "method" parameter is specified, then the other MUST also be
//      specified.
//
//      No methods are defined by this specification.  This is the subject
//      of other specifications, such as the iCalendar Transport-
//      independent Interoperability Protocol (iTIP) defined by [2446bis].
//
//      If this property is not present in the iCalendar object, then a
//      scheduling transaction MUST NOT be assumed.  In such cases, the
//      iCalendar object is merely being used to transport a snapshot of
//
//      some calendar information; without the intention of conveying a
//      scheduling semantic.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       method     = "METHOD" metparam ":" metvalue CRLF
//
//       metparam   = *(";" other-param)
//
//       metvalue   = iana-token
//
//   Example:  The following is a hypothetical example of this property to
//      convey that the iCalendar object is a scheduling request:
//
//       METHOD:REQUEST

type Method struct {
	Parameters []parameters.Parameter
	Value      *types.Text
}

func (a *Method) WritePropertyToStrBuilder(s *strings.Builder) error {
	return properties.DefaultCreatePropertyFunc("METHOD", a.Parameters, a.Value, s)
}
