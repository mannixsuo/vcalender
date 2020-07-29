package components

import (
	"calendar/objects/property/components/properties/changemanage"
	"calendar/objects/property/components/properties/datetime"
	"calendar/objects/property/components/properties/descriptive"
	"calendar/objects/property/components/properties/relationship"
	"strings"
	"testing"
)

func TestTodo_Todo(t *testing.T) {
	//       BEGIN:VTODO
	//       UID:20070313T123432Z-456553@example.com
	//       DTSTAMP:20070313T123432Z
	//       DUE;VALUE=DATE:20070501
	//       SUMMARY:Submit Quebec Income Tax Return for 2006
	//       CLASS:CONFIDENTIAL
	//       CATEGORIES:FAMILY,FINANCE
	//       STATUS:NEEDS-ACTION
	//       END:VTODO
	s:=&strings.Builder{}
	todo := Todo{
		DtStamp:    changemanage.NewDtStamp(2007, 3, 13, 12, 34, 32),
		Uid:        relationship.NewUid("20070313T123432Z-456553@example.com"),
		Class:      &descriptive.ConfidentialClassification,
		Status:     descriptive.NewStatus("NEEDS-ACTION"),
		Summary:    descriptive.NewSummary("Submit Quebec Income Tax Return for 2006"),
		Due:        datetime.NewDueWithDate(2007, 5, 1),
		Categories: []*descriptive.Categories{descriptive.NewCategories("FAMILY","FINANCE")},
	}
	todo.Todo(s)
	v:=VTest{
		S: s.String(),
		T: t,
	}
	v.ContainsOrError("BEGIN:VTODO")
	v.ContainsOrError("UID:20070313T123432Z-456553@example.com")
	v.ContainsOrError("DTSTAMP:20070313T123432Z")
	v.ContainsOrError("DUE;VALUE=DATE:20070501")
	v.ContainsOrError("SUMMARY:Submit Quebec Income Tax Return for 2006")
	v.ContainsOrError("CLASS:CONFIDENTIAL")
	v.ContainsOrError("CATEGORIES:FAMILY,FINANCE")
	v.ContainsOrError("STATUS:NEEDS-ACTION")
	v.ContainsOrError("END:VTODO")
}
