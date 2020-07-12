package miscellaneous

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  An IANA-registered property name
//
//   Value Type:  The default value type is TEXT.  The value type can be
//      set to any value type.
//
//   Property Parameters:  Any parameter can be specified on this
//      property.
//
//   Description:  This specification allows other properties registered
//      with IANA to be specified in any calendar components.  Compliant
//      applications are expected to be able to parse these other IANA-
//      registered properties but can ignore them.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       iana-prop = iana-token *(";" icalparameter) ":" value CRLF
//
//   Example:  The following are examples of properties that might be
//      registered to IANA:
//
//       DRESSCODE:CASUAL
//
//       NON-SMOKING;VALUE=BOOLEAN:TRUE

type Iana struct {
	Name       string
	Parameters []parameters.Parameter
	Values     []types.Value
}

func (i *Iana) Property() (string, error) {
	return properties.DefaultCreateMultiplePropertyFunc(i.Name, i.Parameters, i.Values), nil
}
