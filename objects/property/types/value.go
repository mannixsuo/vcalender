package types

import "strings"

// Value represent the basic value types in calender object
type Value interface {
	// WriteValueToStrBuilder write it's value into string builder
	WriteValueToStrBuilder(s *strings.Builder) error
}
