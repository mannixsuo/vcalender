package property

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  PRODID
//
//   Purpose:  This property specifies the identifier for the product that
//      created the iCalendar object.
//
//   V Type:  TEXT
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  The property MUST be specified once in an iCalendar
//      object.
//
//   Description:  The vendor of the implementation SHOULD assure that
//      this is a globally unique identifier; using some technique such as
//      an FPI value, as defined in [ISO.9070.1991].
//
//      This property SHOULD NOT be used to alter the interpretation of an
//      iCalendar object beyond the semantics specified in this memo.  For
//      example, it is not to be used to further the understanding of non-
//      standard properties.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       prodid     = "PRODID" pidparam ":" pidvalue CRLF
//
//       pidparam   = *(";" other-param)
//       pidvalue   = text
//       ;Any text that describes the product and version
//       ;and that is generally assured of being unique.
//
//   Example:  The following is an example of this property.  It does not
//      imply that English is the default language.
//
//       PRODID:-//ABC Corporation//NONSGML My Product//EN

type ProductIdentifier struct {
	Parameters []parameters.Parameter
	Value      *types.Text
}

func (a *ProductIdentifier) Property() (string, error) {
	return properties.DefaultCreatePropertyFunc("PRODID", a.Parameters, a.Value), nil
}

func NewProductIdentifier(id string) *ProductIdentifier {
	return &ProductIdentifier{
		Value: types.NewText(id),
	}
}
