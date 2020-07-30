package parameters

import (
	"github.com/mmsuo/vcalender/objects/property/types"
	"strings"
)

//   WriteParameterToStrBuilder Name:  DELEGATED-FROM
//
//   Purpose:  To specify the calendar users that have delegated their
//      participation to the calendar user specified by the property.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//       delfromparam       = "DELEGATED-FROM" "=" DQUOTE cal-address
//                             DQUOTE *("," DQUOTE cal-address DQUOTE)
//
//   Description:  This parameter can be specified on properties with a
//      CAL-ADDRESS value type.  This parameter specifies those calendar
//      users that have delegated their participation in a group-scheduled
//      event or to-do to the calendar user specified by the property.
//      The individual calendar address parameter values MUST each be
//      specified in a quoted-string.
//
//   Example:
//
//       ATTENDEE;DELEGATED-FROM="mailto:jsmith@example.com":mailto:
//        jdoe@example.com

type DelegatedFrom struct {
	V []*types.CalAddress
}

func (d *DelegatedFrom) WriteParameterToStrBuilder(s *strings.Builder) error {
	d.delegatedFrom(s)
	return nil
}

func (d *DelegatedFrom) delegatedFrom(s *strings.Builder) {
	s.WriteString("DELEGATED-FROM=")
	for index, addr := range d.V {
		if index != 0 {
			s.WriteString(",")
		}
		s.WriteString("\"")
		addr.WriteValueToStrBuilder(s)
		s.WriteString("\"")
	}
}
