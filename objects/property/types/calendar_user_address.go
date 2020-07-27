package types

import (
	"fmt"
)

//   V Name:  CAL-ADDRESS
//
//   Purpose:  This value type is used to identify properties that contain
//      a calendar user address.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       cal-address        = uri
//
//   Description:  The value is a URI as defined by [RFC3986] or any other
//      IANA-registered form for a URI.  When used to address an Internet
//      email transport address for a calendar user, the value MUST be a
//      mailto URI, as defined by [RFC2368].  No additional content value
//      encoding (i.e., BACKSLASH character encoding, see Section 3.3.11)
//      is defined for this value type.
//
//   Example:
//
//    mailto:jane_doe@example.com

type CalAddress struct {
	V *URI
}

func (c *CalAddress) Value() string {
	return fmt.Sprintf("mailto:%s", c.V.Value())
}

func NewCalAddress(address string) *CalAddress {
	return &CalAddress{V: &URI{V: address}}
}
