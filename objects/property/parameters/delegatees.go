package parameters

import (
	"fmt"
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   WriteParameterToStrBuilder Name:  DELEGATED-TO
//
//   Purpose:  To specify the calendar users to whom the calendar user
//      specified by the property has delegated participation.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//       deltoparam = "DELEGATED-TO" "=" DQUOTE cal-address DQUOTE
//                    *("," DQUOTE cal-address DQUOTE)
//
//   Description:  This parameter can be specified on properties with a
//      CAL-ADDRESS value type.  This parameter specifies those calendar
//      users whom have been delegated participation in a group-scheduled
//      event or to-do by the calendar user specified by the property.
//      The individual calendar address parameter values MUST each be
//      specified in a quoted-string.
//   Example:
//
//       ATTENDEE;DELEGATED-TO="mailto:jdoe@example.com","mailto:jqpublic
//        @example.com":mailto:jsmith@example.com

type DelegatedTo struct {
	V []*types.CalAddress
}

func (d *DelegatedTo) WriteParameterToStrBuilder(s *strings.Builder) error {
	d.delegatedTo(s)
	return nil
}

func (d *DelegatedTo) delegatedTo(s *strings.Builder) string {
	s.WriteString("DELEGATED-TO=")
	for index, addr := range d.V {
		if index != 0 {
			s.WriteString(",")
		}
		s.WriteString(fmt.Sprintf("\""))
		addr.WriteValueToStrBuilder(s)
		s.WriteString(fmt.Sprintf("\""))
	}
	return s.String()
}
