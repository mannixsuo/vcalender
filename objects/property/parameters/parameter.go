package parameters

import "strings"

type Parameter interface {
	WriteParameterToStrBuilder(s *strings.Builder) error
}

func WriteParametersToStrBuilder(p []Parameter, sb *strings.Builder) error {
	for _, v := range p {
		sb.WriteString(";")
		v.WriteParameterToStrBuilder(sb)
	}
	return nil
}
