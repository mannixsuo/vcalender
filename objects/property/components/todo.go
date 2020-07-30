package components

import (
	"github.com/mmsuo/vcalender/objects/property/components/properties/changemanage"
	"github.com/mmsuo/vcalender/objects/property/components/properties/datetime"
	"github.com/mmsuo/vcalender/objects/property/components/properties/descriptive"
	"github.com/mmsuo/vcalender/objects/property/components/properties/miscellaneous"
	"github.com/mmsuo/vcalender/objects/property/components/properties/recurrence"
	"github.com/mmsuo/vcalender/objects/property/components/properties/relationship"
	"strings"
)

//   Component Name:  VTODO
//
//   Purpose:  Provide a grouping of calendar properties that describe a
//      to-do.
//   Format Definition:  A "VTODO" calendar component is defined by the
//      following notation:
//
//       todoc      = "BEGIN" ":" "VTODO" CRLF
//                    todoprop *alarmc
//                    "END" ":" "VTODO" CRLF
//
//       todoprop   = *(
//                  ;
//                  ; The following are REQUIRED,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  dtstamp / uid /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  class / completed / created / description /
//                  dtstart / geo / last-mod / location / organizer /
//                  percent / priority / recurid / seq / status /
//                  summary / url /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; but SHOULD NOT occur more than once.
//                  ;
//                  rrule /
//                  ;
//                  ; Either 'due' or 'duration' MAY appear in
//                  ; a 'todoprop', but 'due' and 'duration'
//                  ; MUST NOT occur in the same 'todoprop'.
//                  ; If 'duration' appear in a 'todoprop',
//                  ; then 'dtstart' MUST also appear in
//                  ; the same 'todoprop'.
//                  ;
//                  due / duration /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  attach / attendee / categories / comment / contact /
//                  exdate / rstatus / related / resources /
//                  rdate / x-prop / iana-prop
//                  ;
//                  )
//
//   Description:  A "VTODO" calendar component is a grouping of component
//      properties and possibly "VALARM" calendar components that
//      represent an action-item or assignment.  For example, it can be
//
//      used to represent an item of work assigned to an individual; such
//      as "turn in travel expense today".
//
//      The "VTODO" calendar component cannot be nested within another
//      calendar component.  However, "VTODO" calendar components can be
//      related to each other or to a "VEVENT" or to a "VJOURNAL" calendar
//      component with the "RELATED-TO" property.
//
//      A "VTODO" calendar component without the "DTSTART" and "DUE" (or
//      "DURATION") properties specifies a to-do that will be associated
//      with each successive calendar date, until it is completed.
//
//   Examples:  The following is an example of a "VTODO" calendar
//      component that needs to be completed before May 1st, 2007.  On
//      midnight May 1st, 2007 this to-do would be considered overdue.
//
//       BEGIN:VTODO
//       UID:20070313T123432Z-456553@example.com
//       DTSTAMP:20070313T123432Z
//       DUE;VALUE=DATE:20070501
//       SUMMARY:Submit Quebec Income Tax Return for 2006
//       CLASS:CONFIDENTIAL
//       CATEGORIES:FAMILY,FINANCE
//       STATUS:NEEDS-ACTION
//       END:VTODO
//
//      The following is an example of a "VTODO" calendar component that
//      was due before 1:00 P.M. UTC on July 9th, 2007 and was completed
//      on July 7th, 2007 at 10:00 A.M. UTC.
//
//       BEGIN:VTODO
//       UID:20070514T103211Z-123404@example.com
//       DTSTAMP:20070514T103211Z
//       DTSTART:20070514T110000Z
//       DUE:20070709T130000Z
//       COMPLETED:20070707T100000Z
//       SUMMARY:Submit Revised Internet-Draft
//       PRIORITY:1
//       STATUS:NEEDS-ACTION
//       END:VTODO
type Todo struct {
	DtStamp      *changemanage.DtStamp
	Uid          *relationship.Uid
	Class        *descriptive.Classification
	Completed    *datetime.Completed
	Created      *changemanage.Created
	Description  *descriptive.Description
	DtStart      *datetime.DateStart
	Geo          *descriptive.Geographic
	LastModified *changemanage.LastModified
	Location     *descriptive.Location
	Organizer    *relationship.Organizer
	Percent      *descriptive.PercentComplete

	Priority *descriptive.Priority
	RecurId  *relationship.RecurrenceId

	Seq      *changemanage.Sequence
	Status   *descriptive.Status
	Summary  *descriptive.Summary
	Url      *relationship.Url
	RRule    *recurrence.RRule
	Due      *datetime.Due
	Duration *datetime.Duration

	Attach     []*descriptive.Attach
	Attendee   []*relationship.Attendee
	Categories []*descriptive.Categories
	Comment    []*descriptive.Comment
	Contact    []*relationship.Contact
	ExDate     []*recurrence.ExDate
	RStatus    []*miscellaneous.RequestStatus
	Related    []*relationship.RelatedTo
	Resources  []*descriptive.Resources
	RDate      []*recurrence.RDate
	Xprop      []*miscellaneous.NoStandard
	IanaProp   []*miscellaneous.Iana
	Alarm      *Alarm
}

func (t *Todo) Todo(b *strings.Builder) error {
	b.WriteString("BEGIN:VTODO\n")
	WriteProperty(b, t.DtStamp)
	WriteProperty(b, t.Uid)
	WriteProperty(b, t.DtStart)
	WriteProperty(b, t.Class)
	WriteProperty(b, t.Completed)
	WriteProperty(b, t.Created)
	WriteProperty(b, t.Description)
	WriteProperty(b, t.Geo)
	WriteProperty(b, t.LastModified)
	WriteProperty(b, t.Location)
	WriteProperty(b, t.Organizer)
	WriteProperty(b, t.Percent)
	WriteProperty(b, t.Priority)
	WriteProperty(b, t.Seq)
	WriteProperty(b, t.Status)
	WriteProperty(b, t.Summary)
	WriteProperty(b, t.Url)
	WriteProperty(b, t.RecurId)
	WriteProperty(b, t.RRule)
	WriteProperty(b, t.Due)
	WriteProperty(b, t.Duration)

	WriteProperties(b, t.Attach)
	WriteProperties(b, t.Attendee)
	WriteProperties(b, t.Categories)
	WriteProperties(b, t.Comment)
	WriteProperties(b, t.Contact)
	WriteProperties(b, t.ExDate)
	WriteProperties(b, t.RStatus)
	WriteProperties(b, t.Related)
	WriteProperties(b, t.Resources)
	WriteProperties(b, t.RDate)
	WriteProperties(b, t.Xprop)
	WriteProperties(b, t.IanaProp)
	if t.Alarm != nil {
		t.Alarm.Alarm(b)
	}
	b.WriteString("END:VTODO\n")
	return nil
}

func (t *Todo) WriteComponentToStrBuilder(s *strings.Builder) error {
	return t.Todo(s)
}
