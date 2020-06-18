package model

import (
	"testing"
)

const (
	EmailPropString = "BEGIN:VALARM\n" +
		"TRIGGER;RELATED=END:-P2D\n" +
		"ACTION:EMAIL\n" +
		"ATTENDEE:mailto:john_doe@example.com\n" +
		"SUMMARY:*** REMINDER: SEND AGENDA FOR WEEKLY STAFF MEETING ***\n" +
		"DESCRIPTION:A draft agenda needs to be sent out to the attendees\n" +
		"  to the weekly managers meeting (MGR-LIST). Attached is a\n" +
		"  pointer the document template for the agenda file.\n" +
		"ATTACH;FMTTYPE=application/msword:http://example.com/\n" +
		"templates/agenda.doc\n" +
		"END:VALARM"
)

func TestEmailProp_AlarmProp(t *testing.T) {
	_ = EmailProp{
		Description: "A draft agenda needs to be sent out to the attendees to the weekly managers meeting (MGR-LIST). Attached is a pointer the document template for the agenda file.",
		Trigger:     nil,
		Summary:     "",
		Attendee:    nil,
		Duration:    nil,
		Repeat:      0,
		Attach:      nil,
		XProp:       nil,
		IanaProp:    nil,
	}
}
