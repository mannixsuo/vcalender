package parameters

import "strings"

type Parameter interface {
	Parameter() string
}

type Parameters []Parameter

func NewParameters(p ...Parameter) Parameters {
	return p
}

func (p *Parameters) Parameters() string {
	sb := strings.Builder{}
	for _, v := range *p {
		sb.WriteString(";")
		sb.WriteString(v.Parameter())
	}
	return sb.String()
}
