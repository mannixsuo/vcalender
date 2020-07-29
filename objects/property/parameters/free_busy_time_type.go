package parameters

import (
	"fmt"
	"strings"
)

//   WriteParameterToStrBuilder Name:  FBTYPE
//
//   Purpose:  To specify the free or busy time type.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//       fbtypeparam        = "FBTYPE" "=" ("FREE" / "BUSY"
//                          / "BUSY-UNAVAILABLE" / "BUSY-TENTATIVE"
//                          / x-name
//                ; Some experimental iCalendar free/busy type.
//                          / iana-token)
//                ; Some other IANA-registered iCalendar free/busy type.
//
//   Description:  This parameter specifies the free or busy time type.
//      The value FREE indicates that the time interval is free for
//      scheduling.  The value BUSY indicates that the time interval is
//      busy because one or more events have been scheduled for that
//      interval.  The value BUSY-UNAVAILABLE indicates that the time
//      interval is busy and that the interval can not be scheduled.  The
//      value BUSY-TENTATIVE indicates that the time interval is busy
//      because one or more events have been tentatively scheduled for
//      that interval.  If not specified on a property that allows this
//      parameter, the default is BUSY.  Applications MUST treat x-name
//      and iana-token values they don't recognize the same way as they
//      would the BUSY value.
//
//   Example:  The following is an example of this parameter on a
//      "FREEBUSY" property.
//
//       FREEBUSY;FBTYPE=BUSY:19980415T133000Z/19980415T170000Z

type FreeBusyType struct {
	V string
}

var Free = FreeBusyType{"FREE"}
var Busy = FreeBusyType{"BUSY"}
var BusyUnavailable = FreeBusyType{"BUSY-UNAVAILABLE"}
var BusyTentative = FreeBusyType{"BUSY-TENTATIVE"}

func (f *FreeBusyType) WriteParameterToStrBuilder(s *strings.Builder) error {
	s.WriteString(fmt.Sprintf("FBTYPE=%s", f.V))
	return nil
}
