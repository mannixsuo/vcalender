package types

import "strings"

//   V Name:  FLOAT
//
//   Purpose:  This value type is used to identify properties that contain
//      a real-number value.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       float      = (["+"] / "-") 1*DIGIT ["." 1*DIGIT]
//
//   Description:  If the property permits, multiple "float" values are
//      specified by a COMMA-separated list of values.
//
//      No additional content value encoding (i.e., BACKSLASH character
//      encoding, see Section 3.3.11) is defined for this value type.
//
//   Example:
//
//       1000000.0000001
//       1.333
//       -3.14

type Float struct {
	V string
}

func (f *Float) WriteValueToStrBuilder(s *strings.Builder) error {
	s.WriteString(f.V)
	return nil
}
