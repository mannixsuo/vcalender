package parameters

import (
	"calendar/objects/property/types"
	"fmt"
	"strings"
)

//   Parameter Name:  MEMBER
//
//   Purpose:  To specify the group or list membership of the calendar
//      user specified by the property.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//       memberparam        = "MEMBER" "=" DQUOTE cal-address DQUOTE
//                            *("," DQUOTE cal-address DQUOTE)
//   Description:  This parameter can be specified on properties with a
//      CAL-ADDRESS value type.  The parameter identifies the groups or
//      list membership for the calendar user specified by the property.
//      The parameter value is either a single calendar address in a
//      quoted-string or a COMMA-separated list of calendar addresses,
//      each in a quoted-string.  The individual calendar address
//      parameter values MUST each be specified in a quoted-string.
//
//   Example:
//
//       ATTENDEE;MEMBER="mailto:ietf-calsch@example.org":mailto:
//        jsmith@example.com
//
//       ATTENDEE;MEMBER="mailto:projectA@example.com","mailto:pr
//        ojectB@example.com":mailto:janedoe@example.com

type Member struct {
	V []*types.CalAddress
}

func (m *Member) Parameter() string {
	sb := strings.Builder{}
	sb.WriteString("MEMBER=")

	for index, addr := range m.V {
		if index != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("\"%s\"", addr.Value()))
	}
	return sb.String()
}