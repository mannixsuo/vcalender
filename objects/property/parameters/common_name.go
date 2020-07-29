package parameters

import (
	"fmt"
	"strings"
)

//   WriteParameterToStrBuilder Name:  CN
//
//   Purpose:  To specify the common name to be associated with the
//      calendar user specified by the property.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//     cnparam    = "CN" "=" param-value
//
//   Description:  This parameter can be specified on properties with a
//      CAL-ADDRESS value type.  The parameter specifies the common name
//      to be associated with the calendar user specified by the property.
//      The parameter value is text.  The parameter value can be used for
//      display text to be associated with the calendar address specified
//      by the property.
//
//   Example:
//
//       ORGANIZER;CN="John Smith":mailto:jsmith@example.com

type CommonName struct {
	V string
}

func (c *CommonName) WriteParameterToStrBuilder(s *strings.Builder) error {
	s.WriteString(fmt.Sprintf("CN=\"%s\"", c.V))
	return nil
}
