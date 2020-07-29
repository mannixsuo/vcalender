package properties

import (
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
	"strings"
)

type Property interface {
	WritePropertyToStrBuilder(s *strings.Builder) error
}

type CheckPropertyFunc func([]*parameters.Parameter, types.Value) error
type CreatePropertyFunc func(string, []parameters.Parameter, types.Value) string

func DefaultCheckPropertyFunc(p []parameters.Parameter, v types.Value) error {
	return nil
}

func DefaultCreatePropertyFunc(name string, p []parameters.Parameter, v types.Value, sb *strings.Builder) error {
	sb.WriteString(name)
	parameters.WriteParametersToStrBuilder(p, sb)
	sb.WriteString(":")
	v.WriteValueToStrBuilder(sb)
	sb.WriteString("\n")
	return nil
}

func DefaultCreateMultiplePropertyFunc(name string, p []parameters.Parameter, v []types.Value, sb *strings.Builder) error {
	sb.WriteString(name)
	if len(p) > 0 {
		parameters.WriteParametersToStrBuilder(p, sb)
	}
	sb.WriteString(":")
	last := len(v) - 1
	for index, value := range v {
		value.WriteValueToStrBuilder(sb)
		if index != last {
			sb.WriteString(",")
		}
	}
	sb.WriteString("\n")
	return nil
}
