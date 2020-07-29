package components

import "strings"

type Component interface {
	WriteComponentToStrBuilder(s *strings.Builder) error
}
