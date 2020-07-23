package components

import (
	"calendar/objects/property/components/properties/changemanage"
	"calendar/objects/property/components/properties/datetime"
	"calendar/objects/property/components/properties/descriptive"
	"calendar/objects/property/components/properties/relationship"
	"fmt"
	"strings"
	"testing"
)

func TestEvent_Event(t *testing.T) {
	//       BEGIN:VEVENT
	//       UID:19970901T130000Z-123401@example.com
	//       DTSTAMP:19970901T130000Z
	//       DTSTART:19970903T163000Z
	//       DTEND:19970903T190000Z
	//       SUMMARY:Annual Employee Review
	//       CLASS:PRIVATE
	//       CATEGORIES:BUSINESS,HUMAN RESOURCES
	//       END:VEVENT
	e := Event{
		DtStamp:    changemanage.NewDtStamp(1997, 9, 1, 13, 0, 0),
		Uid:        relationship.NewUid("19970901T130000Z-123401@example.com"),
		DtStart:    datetime.NewDateStartWithDatetime(1997, 9, 3, 16, 30, 0),
		DtEnd:      datetime.NewDateTimeDateEnd(1997, 9, 3, 19, 0, 0),
		Summary:    descriptive.NewSummary("Annual Employee Review"),
		Class:      &descriptive.PrivateClassification,
		Categories: []*descriptive.Categories{descriptive.NewCategories("BUSINESS", "HUMAN RESOURCES")},
	}
	v := &VTest{
		S: e.Event(),
		T: t,
	}
	v.ContainsOrError("BEGIN:VEVENT")
	v.ContainsOrError("UID:19970901T130000Z-123401@example.com")
	v.ContainsOrError("DTSTAMP:19970901T130000Z")
	v.ContainsOrError("DTSTART:19970903T163000Z")
	v.ContainsOrError("DTEND:19970903T190000Z")
	v.ContainsOrError("SUMMARY:Annual Employee Review")
	v.ContainsOrError("CLASS:PRIVATE")
	v.ContainsOrError("CATEGORIES:BUSINESS,HUMAN RESOURCES")
	v.ContainsOrError("END:VEVENT")

	//       BEGIN:VEVENT
	//       UID:19970901T130000Z-123402@example.com
	//       DTSTAMP:19970901T130000Z
	//       DTSTART:19970401T163000Z
	//       DTEND:19970402T010000Z
	//       SUMMARY:Laurel is in sensitivity awareness class.
	//       CLASS:PUBLIC
	//       CATEGORIES:BUSINESS,HUMAN RESOURCES
	//       TRANSP:TRANSPARENT
	//       END:VEVENT
	event2 := Event{
		DtStamp:     changemanage.NewDtStamp(1997, 9, 1, 13, 0, 0),
		Uid:         relationship.NewUid("19970901T130000Z-123402@example.com"),
		DtStart:     datetime.NewDateStartWithDatetime(1997, 4, 1, 16, 30, 0),
		Class:       &descriptive.PublicClassification,
		Summary:     descriptive.NewSummary("Laurel is in sensitivity awareness class."),
		Transparent: &datetime.TransparentTransparent,
		DtEnd:       datetime.NewDateTimeDateEnd(1997, 4, 2, 1, 0, 0),
		Categories:  []*descriptive.Categories{descriptive.NewCategories("BUSINESS", "HUMAN RESOURCES")},
	}
	v = &VTest{
		S: event2.Event(),
		T: t,
	}
	v.ContainsOrError("19970901T130000Z-123402@example.com")
	v.ContainsOrError("DTSTAMP:19970901T130000Z")
	v.ContainsOrError("DTSTART:19970401T163000Z")
	v.ContainsOrError("DTEND:19970402T010000Z")
	v.ContainsOrError("SUMMARY:Laurel is in sensitivity awareness class.")
	v.ContainsOrError("CLASS:PUBLIC")
	v.ContainsOrError("TRANSP:TRANSPARENT")
	v.ContainsOrError("CATEGORIES:BUSINESS,HUMAN RESOURCES")
}

type VTest struct {
	S string
	T *testing.T
}

func (v *VTest) ContainsOrError(sub string, ) {
	if !strings.Contains(v.S, sub) {
		fmt.Printf("%s not contain %s", v.S, sub)
		v.T.Error()
	}
}
