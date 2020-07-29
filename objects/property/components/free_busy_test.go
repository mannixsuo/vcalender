package components

import (
	"calendar/objects/property/components/properties/changemanage"
	"calendar/objects/property/components/properties/datetime"
	"calendar/objects/property/components/properties/relationship"
	"strings"
	"testing"
)

func TestFreeBusy_FreeBusy(t *testing.T) {
	//       BEGIN:VFREEBUSY
	//       UID:19970901T082949Z-FA43EF@example.com
	//       ORGANIZER:mailto:jane_doe@example.com
	//       ATTENDEE:mailto:john_public@example.com
	//       DTSTART:19971015T050000Z
	//       DTEND:19971016T050000Z
	//       DTSTAMP:19970901T083000Z
	//       END:VFREEBUSY
	s := &strings.Builder{}
	b:=FreeBusy{
		DtStamp:   changemanage.NewDtStamp(1997, 9, 1, 8, 30, 0),
		Uid:       relationship.NewUid("19970901T082949Z-FA43EF@example.com"),
		DtStart:   datetime.NewDateStartWithDatetime(1997,10,15,5,0,0),
		DtEnd:     datetime.NewDateTimeDateEnd(1997,10,16,5,0,0),
		Organizer: relationship.NewOrganizer("jane_doe@example.com"),
		Attendee:  []*relationship.Attendee{relationship.NewAttendee("john_public@example.com")},
	}
	b.FreeBusy(s)
	v:=VTest{
		S: s.String(),
		T: t,
	}
	v.ContainsOrError("BEGIN:VFREEBUSY")
	v.ContainsOrError("UID:19970901T082949Z-FA43EF@example.com")
	v.ContainsOrError("ORGANIZER:mailto:jane_doe@example.com")
	v.ContainsOrError("ATTENDEE:mailto:john_public@example.com")
	v.ContainsOrError("DTSTART:19971015T050000Z")
	v.ContainsOrError("DTEND:19971016T050000Z")
	v.ContainsOrError("DTSTAMP:19970901T083000Z")
	v.ContainsOrError("END:VFREEBUSY")

}