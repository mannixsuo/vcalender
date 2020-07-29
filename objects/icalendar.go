package objects

import (
	"calendar/objects/property"
	"calendar/objects/property/components"
	"calendar/objects/property/components/properties/miscellaneous"
	"strings"
)

type Calendar struct {
	ProdId     *property.ProductIdentifier
	Version    *property.Version
	CalScale   *property.CalendarScale
	Method     *property.Method
	Components []components.Component
	Xprop      []*miscellaneous.NoStandard
	IanaProp   []*miscellaneous.Iana
}

func (c *Calendar) Calendar() (string, error) {
	s := &strings.Builder{}
	s.WriteString("BEGIN:VCALENDAR\n")
	components.WriteProperty(s, c.ProdId)
	components.WriteProperty(s, c.Version)
	components.WriteProperty(s, c.CalScale)
	components.WriteProperty(s, c.Method)
	for _, comp := range c.Components {
		err := comp.WriteComponentToStrBuilder(s)
		if err != nil {
			return "", err
		}
	}
	components.WriteProperties(s, c.Xprop)
	components.WriteProperties(s, c.IanaProp)
	s.WriteString("END:VCALENDAR\n")
	return s.String(), nil
}
