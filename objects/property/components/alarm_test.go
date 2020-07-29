package components

import (
	"calendar/objects/property/components/properties/alarm"
	"calendar/objects/property/components/properties/descriptive"
	"calendar/objects/property/components/properties/relationship"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
	"fmt"
	"strings"
	"testing"
	"time"
)

//       BEGIN:VALARM
//       TRIGGER;VALUE=DATE-TIME:19970317T133000Z
//       REPEAT:4
//       DURATION:PT15M
//       ACTION:AUDIO
//       ATTACH;FMTTYPE=audio/basic:ftp://example.com/pub/
//        sounds/bell-01.aud
//       END:VALARM
func TestAlarm_Audio(t *testing.T) {
	s := &strings.Builder{}
	audio := Alarm{
		Trigger: &alarm.Trigger{
			Parameters: []parameters.Parameter{&parameters.DateTime},
			Value: &types.DateTime{
				V:      time.Date(1997, 3, 17, 13, 30, 0, 0, time.Local),
				Format: types.UTCDateTimeFormat,
			},
		},
		Repeat: &types.Integer{V: 4},
		Duration: &types.Duration{
			DurMinute: 15,
		},
		Action: &alarm.Audio,
		Attach: []*descriptive.Attach{{
			Parameters: []parameters.Parameter{&parameters.FmtType{V: "audio/basic"}},
			Value:      &types.URI{V: "ftp://example.com/pub/sounds/bell-01.aud"},
		}},
	}
	audio.Alarm(s)
	fmt.Println(s.String())
}

//       BEGIN:VALARM
//       TRIGGER:-PT30M
//       REPEAT:2
//       DURATION:PT15M
//       ACTION:DISPLAY
//       DESCRIPTION:Breakfast meeting with executive\n
//        team at 8:30 AM EST.
//       END:VALARM
func TestAlarm_Display(t *testing.T) {
	s := &strings.Builder{}
	dis := Alarm{
		Action: &alarm.Display,
		Trigger: &alarm.Trigger{
			Value: &types.Duration{
				DurMinute: 15,
			},
		},
		Description: &descriptive.Description{Value: &types.Text{V: "Breakfast meeting with executive team at 8:30 AM EST."}},
		Repeat:      &types.Integer{V: 2},
	}
	dis.Alarm(s)
	fmt.Println(s.String())
}

//       BEGIN:VALARM
//       TRIGGER;RELATED=END:-P2D
//       ACTION:EMAIL
//       ATTENDEE:mailto:john_doe@example.com
//       SUMMARY:*** REMINDER: SEND AGENDA FOR WEEKLY STAFF MEETING ***
//       DESCRIPTION:A draft agenda needs to be sent out to the attendees
//         to the weekly managers meeting (MGR-LIST). Attached is a
//         pointer the document template for the agenda file.
//       ATTACH;FMTTYPE=application/msword:http://example.com/
//        templates/agenda.doc
//       END:VALARM

func TestAlarm_Email(t *testing.T) {
	s := &strings.Builder{}

	dis := Alarm{
		Action: &alarm.Email,
		Trigger: &alarm.Trigger{
			Parameters: []parameters.Parameter{&parameters.EndRelated},
			Value: &types.Duration{
				Negative: true,
				DurDay:   2,
			},
		},
		Description: &descriptive.Description{Value: &types.Text{V: "A draft agenda needs to be sent out to the attendees to the weekly managers meeting (MGR-LIST). Attached is a pointer the document template for the agenda file."}},
		Attendee:    []*relationship.Attendee{{Value: &types.CalAddress{V: &types.URI{V: "john_doe@example.com"}}}},
		Attach: []*descriptive.Attach{{
			Parameters: []parameters.Parameter{&parameters.FmtType{V: "application/msword"}},
			Value:      &types.URI{V: "http://example.com/templates/agenda.doc"},
		}},
		Summary: &descriptive.Summary{
			Value: &types.Text{V: "*** REMINDER: SEND AGENDA FOR WEEKLY STAFF MEETING ***"},
		},
	}
	dis.Alarm(s)
	fmt.Println(s.String())
}
