package components

import (
	"calendar/objects/property/components/properties/changemanage"
	"calendar/objects/property/components/properties/datetime"
	"calendar/objects/property/components/properties/descriptive"
	"calendar/objects/property/components/properties/miscellaneous"
	"calendar/objects/property/components/properties/recurrence"
	"calendar/objects/property/components/properties/relationship"
	"strings"
)

//   Component Name:  VJOURNAL
//
//   Purpose:  Provide a grouping of component properties that describe a
//      journal entry.
//
//   Format Definition:  A "VJOURNAL" calendar component is defined by the
//      following notation:
//
//       journalc   = "BEGIN" ":" "VJOURNAL" CRLF
//                    jourprop
//                    "END" ":" "VJOURNAL" CRLF
//
//       jourprop   = *(
//                  ;
//                  ; The following are REQUIRED,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  dtstamp / uid /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  class / created / dtstart /
//                  last-mod / organizer / recurid / seq /
//                  status / summary / url /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; but SHOULD NOT occur more than once.
//                  ;
//                  rrule /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  attach / attendee / categories / comment /
//                  contact / description / exdate / related / rdate /
//                  rstatus / x-prop / iana-prop
//                  ;
//                  )
//
//   Description:  A "VJOURNAL" calendar component is a grouping of
//      component properties that represent one or more descriptive text
//      notes associated with a particular calendar date.  The "DTSTART"
//      property is used to specify the calendar date with which the
//      journal entry is associated.  Generally, it will have a DATE value
//      data type, but it can also be used to specify a DATE-TIME value
//      data type.  Examples of a journal entry include a daily record of
//      a legislative body or a journal entry of individual telephone
//      contacts for the day or an ordered list of accomplishments for the
//      day.  The "VJOURNAL" calendar component can also be used to
//      associate a document with a calendar date.
//      The "VJOURNAL" calendar component does not take up time on a
//      calendar.  Hence, it does not play a role in free or busy time
//      searches -- it is as though it has a time transparency value of
//      TRANSPARENT.  It is transparent to any such searches.
//
//      The "VJOURNAL" calendar component cannot be nested within another
//      calendar component.  However, "VJOURNAL" calendar components can
//      be related to each other or to a "VEVENT" or to a "VTODO" calendar
//      component, with the "RELATED-TO" property.
//
//   Example:  The following is an example of the "VJOURNAL" calendar
//      component:
//
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

type Journal struct {
	DtStamp      *changemanage.DtStamp
	Uid          *relationship.Uid
	Class        *descriptive.Classification
	Created      *changemanage.Created
	DtStart      *datetime.DateStart
	LastModified *changemanage.LastModified
	Organizer    *relationship.Organizer
	RecurId      *relationship.RecurrenceId
	Seq          *changemanage.Sequence
	Status       *descriptive.Status
	Summary      *descriptive.Summary
	Url          *relationship.Url
	RRule        *recurrence.RRule
	Attach       []*descriptive.Attach
	Attendee     []*relationship.Attendee
	Categories   []*descriptive.Categories
	Comment      []*descriptive.Comment
	Contact      []*relationship.Contact
	Description  []*descriptive.Description
	ExDate       []*recurrence.ExDate
	Related      []*relationship.RelatedTo
	RDate        []*recurrence.RDate
	RStatus      []*miscellaneous.RequestStatus
	Xprop        []*miscellaneous.NoStandard
	IanaProp     []*miscellaneous.Iana
}

func (j *Journal) Journal() string {
	b := strings.Builder{}
	b.WriteString("BEGIN:VJOURNAL\n")
	WriteProperty(&b, j.DtStamp)
	WriteProperty(&b, j.Uid)
	WriteProperty(&b, j.DtStart)
	WriteProperty(&b, j.Class)
	WriteProperty(&b, j.Created)
	WriteProperty(&b, j.LastModified)
	WriteProperty(&b, j.Organizer)
	WriteProperty(&b, j.Seq)
	WriteProperty(&b, j.Status)
	WriteProperty(&b, j.Summary)
	WriteProperty(&b, j.Url)
	WriteProperty(&b, j.RecurId)
	WriteProperty(&b, j.RRule)

	WriteProperties(&b, j.Description)

	WriteProperties(&b, j.Attach)
	WriteProperties(&b, j.Attendee)
	WriteProperties(&b, j.Categories)
	WriteProperties(&b, j.Comment)
	WriteProperties(&b, j.Contact)
	WriteProperties(&b, j.ExDate)
	WriteProperties(&b, j.RStatus)
	WriteProperties(&b, j.Related)
	WriteProperties(&b, j.RDate)
	WriteProperties(&b, j.Xprop)
	WriteProperties(&b, j.IanaProp)
	b.WriteString("END:VJOURNAL\n")
	return b.String()
}
