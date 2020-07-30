package components

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties/changemanage"
	"github.com/mmsuo/vcalender/objects/property/components/properties/datetime"
	"github.com/mmsuo/vcalender/objects/property/components/properties/descriptive"
	"github.com/mmsuo/vcalender/objects/property/components/properties/relationship"
	"strings"
	"testing"
)

func TestJournal_Journal(t *testing.T) {
	//       BEGIN:VJOURNAL
	//       UID:19970901T130000Z-123405@example.com
	//       DTSTAMP:19970901T130000Z
	//       DTSTART;VALUE=DATE:19970317
	//       SUMMARY:Staff meeting minutes
	//       DESCRIPTION:1. Staff meeting: Participants include Joe\,
	//         Lisa\, and Bob. Aurora project plans were reviewed.
	//         There is currently no budget reserves for this project.
	//         Lisa will escalate to management. Next meeting on Tuesday.\n
	//        2. Telephone Conference: ABC Corp. sales representative
	//         called to discuss new printer. Promised to get us a demo by
	//         Friday.\n3. Henry Miller (Handsoff Insurance): Car was
	//         totaled by tree. Is looking into a loaner car. 555-2323
	//         (tel).
	//       END:VJOURNAL
	s := &strings.Builder{}

	j:=Journal{
		DtStamp:     changemanage.NewDtStamp(1997, 9, 1, 13, 0, 0),
		Uid:         relationship.NewUid("19970901T130000Z-123405@example.com"),
		DtStart:     datetime.NewDateStartWithDate(1997, 3, 17),
		Summary:     descriptive.NewSummary("Staff meeting minutes"),
		Description: []*descriptive.Description{descriptive.NewDescription("1. Staff meeting: Participants include Joe,Lisa, and Bob. Aurora project plans were reviewed.There is currently no budget reserves for this project.2. Telephone Conference: ABC Corp. sales representative called to discuss new printer. Promised to get us a demo by Friday.\n3. Henry Miller (Handsoff Insurance): Car was totaled by tree. Is looking into a loaner car. 555-2323(tel).")},
	}
	j.Journal(s)
	v:=VTest{
		S: s.String(),
		T: t,
	}
	v.ContainsOrError("UID:19970901T130000Z-123405@example.com")
	v.ContainsOrError("DTSTAMP:19970901T130000Z")
	v.ContainsOrError("DTSTART;VALUE=DATE:19970317")
	v.ContainsOrError("SUMMARY:Staff meeting minutes")
	v.ContainsOrError("1. Staff meeting: Participants include Joe,Lisa, and Bob. Aurora project plans were reviewed.There is currently no budget reserves for this project.2. Telephone Conference: ABC Corp. sales representative called to discuss new printer. Promised to get us a demo by Friday.\n3. Henry Miller (Handsoff Insurance): Car was totaled by tree. Is looking into a loaner car. 555-2323(tel).")
}