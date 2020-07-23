package components

import (
	"calendar/objects/property/components/properties/changemanage"
	"calendar/objects/property/components/properties/datetime"
	"calendar/objects/property/components/properties/descriptive"
	"calendar/objects/property/components/properties/relationship"
	"fmt"
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
	event := e.Event()
	fmt.Println(event)

}