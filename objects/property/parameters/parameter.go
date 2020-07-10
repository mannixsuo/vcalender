package parameters

import "strings"

type Parameter interface {
	Parameter() string
}

func Parameters(p []Parameter) string {
	sb := strings.Builder{}
	for _, v := range p {
		sb.WriteString(";")
		sb.WriteString(v.Parameter())
	}
	return sb.String()
}
