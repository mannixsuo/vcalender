package model

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestEmailProp_AlarmProp(t *testing.T) {
	emailPropString :=
		"BEGIN:VALARM\n" +
			"ACTION:EMAIL\n" +
			"DESCRIPTION:A draft agenda needs to be sent out to the attendees to the weekly managers meeting (MGR-LIST). Attached is a pointer the document template or the agenda file.\n" +
			"TRIGGER;RELATED=END:-P2D\n" +
			"SUMMARY:*** REMINDER: SEND AGENDA FOR WEEKLY STAFF MEETING ***\n" +
			"ATTENDEE:mailto:john_doe@example.com\n" +
			"ATTACH;FMTTYPE=application/msword:http://example.com/templates/agenda.doc\n" +
			"END:VALARM\n"
	e := EmailProp{
		Description: "A draft agenda needs to be sent out to the attendees to the weekly managers meeting (MGR-LIST). Attached is a pointer the document template or the agenda file.",
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
	s := sb.String()
	fmt.Print(s)
	if s != emailPropString {
		t.Error()
	}
}

func TestAudioProp_AlarmProp(t *testing.T) {
	right := "BEGIN:VALARM\n" +
		"TRIGGER;VALUE=DATE-TIME:19970317T133000Z\n" +
		"REPEAT:4\n" +
		"DURATION:PT15M\n" +
		"ACTION:AUDIO\n" +
		"ATTACH;FMTTYPE=audio/basic:ftp://example.com/pub/sounds/bell-01.aud\n" +
		"END:VALARM\n"
	audio := AudioProp{
		Trigger: &Trigger{
			Related2Start: true,
			DateTime:      &DateTime{Date: time.Date(1997, 3, 17, 13, 30, 00, 00, time.Local)},
		},
		Repeat: 4,
		Duration: &Duration{
			DurMinute: 15,
		},
		Attach: &Attach{
			FmtType: "audio/basic",
			URI:     &URI{Uri: "ftp://example.com/pub/sounds/bell-01.aud"},
		},
	}
	sb := strings.Builder{}
	a := Alarmc{Prop: &audio}
	_ = a.Alarmc(&sb)
	s := sb.String()
	fmt.Print(s)
	if s != right {
		t.Error()
	}
}
