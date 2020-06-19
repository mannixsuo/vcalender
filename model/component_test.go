package model

import (
	"fmt"
	"strings"
	"testing"
)

const (
	EmailPropString = "BEGIN:VALARM\n" +
		"TRIGGER;RELATED=END:-P2D\n" +
		"ACTION:EMAIL\n" +
		"ATTENDEE:mailto:john_doe@example.com\n" +
		"SUMMARY:*** REMINDER: SEND AGENDA FOR WEEKLY STAFF MEETING ***\n" +
		"DESCRIPTION:A draft agenda needs to be sent out to the attendees to the weekly man\n" +
		" agers meeting (MGR-LIST). Attached is a pointer the document template \n" +
		" or the agenda file.\n" +
		"ATTACH;FMTTYPE=application/msword:http://example.com/\n" +
		"templates/agenda.doc\n" +
		"END:VALARM"
)

func TestEmailProp_AlarmProp(t *testing.T) {

	e := EmailProp{
		Description: "fA draft agenda needs to be sent out to the attendees to the weekly managers meeting (MGR-LIST). Attached is a pointer the document template or the agenda file.",
		Trigger: &Trigger{
			Related2Start: false,
			Duration: &Duration{
				Negative: true,
				DurDay:   2,
			},
		},
		Summary:  "*** REMINDER: SEND AGENDA FOR WEEKLY STAFF MEETING ***",
		Attendee: []CalAddress{{URI{Uri: "john_doe@example.com"}}},
		Duration: nil,
		Repeat:   0,
		Attach: []Attach{{
			FmtType: "application/msword",
			URI:     &URI{Uri: "http://example.com/templates/agenda.doc"},
		}},
		XProp:    nil,
		IanaProp: nil,
	}
	sb := strings.Builder{}
	a := Alarmc{Prop: &e}
	_ = a.Alarmc(&sb)
	fmt.Print(sb.String())
}
