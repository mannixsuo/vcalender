package properties

import (
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
	"strings"
)

type Property interface {
	Property() (string, error)
}

type CheckPropertyFunc func([]*parameters.Parameter, types.Value) error
type CreatePropertyFunc func(string, []parameters.Parameter, types.Value) string

func DefaultCheckPropertyFunc(p []parameters.Parameter, v types.Value) error {
	return nil
}

func DefaultCreatePropertyFunc(name string, p []parameters.Parameter, v types.Value) string {
	sb := strings.Builder{}
	sb.WriteString(name)
	sb.WriteString(parameters.Parameters(p))
	sb.WriteString(":")
	sb.WriteString(v.Value())
	return sb.String()
}

func DefaultCreateMultiplePropertyFunc(name string, p []parameters.Parameter, v []types.Value) string {
	sb := strings.Builder{}
	sb.WriteString(name)
	sb.WriteString(parameters.Parameters(p))
	sb.WriteString(":")
	for index, value := range v {
		sb.WriteString(value.Value())
		if index != len(v)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}
