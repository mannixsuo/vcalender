package model

import (
	"fmt"
	"strings"
	"time"
)

type Component interface {
	Build()
}

//    Component Name:  VEVENT
//
//   Purpose:  Provide a grouping of component properties that describe an
//      event.
//
//   Format Definition:  A "VEVENT" calendar component is defined by the
//      following notation:
//
//       eventc     = "BEGIN" ":" "VEVENT" CRLF
//                    eventprop *alarmc
//                    "END" ":" "VEVENT" CRLF
//
//       eventprop  = *(
//                  ;
//                  ; The following are REQUIRED,
//                  ; but MUST NOT occur more than once.
//                   ;
//                  dtstamp / uid /
//                  ;
//                  ; The following is REQUIRED if the component
//                  ; appears in an iCalendar object that doesn't
//                  ; specify the "METHOD" property; otherwise, it
//                  ; is OPTIONAL; in any case, it MUST NOT occur
//                  ; more than once.
//                  ;
//                  dtstart /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  class / created / description / geo /
//                  last-mod / location / organizer / priority /
//                  seq / status / summary / transp /
//                  url / recurid /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; but SHOULD NOT occur more than once.
//                  ;
//                  rrule /
//                  ;
//                  ; Either 'dtend' or 'duration' MAY appear in
//                  ; a 'eventprop', but 'dtend' and 'duration'
//                  ; MUST NOT occur in the same 'eventprop'.
//                  ;
//                  dtend / duration /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  attach / attendee / categories / comment /
//                  contact / exdate / rstatus / related /
//                  resources / rdate / x-prop / iana-prop
//                  ;
//                  )
type Eventc struct {
	dtstamp     string
	uid         string
	dtstart     time.Time
	class       string
	created     time.Time
	description string
	geo         string
	lastmod     string
	location    string
	organizer   string
	priority    string
	seq         string
	status      string
	summary     string
	transp      string
	url         string
	recurid     string
	rrule       string
	dtend       time.Time
	duration    string
	attach      []string
	attendee    []string
	categories  []string
	comment     []string
	contact     []string
	exdate      []string
	rstatus     []string
	related     []string
	resources   []string
	rdate       []string
	Xprop       []string
	ianaProp    []string
	Alarms      []Alarmc
}

func (e *Eventc) Build() {

}

//    Purpose:  Provide a grouping of calendar properties that describe a
//      to-do.
//    Format Definition:  A "VTODO" calendar component is defined by the
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
type Todoc struct {
	dtstamp     string
	uid         string
	class       string
	completed   string
	created     time.Time
	description string
	dtstart     time.Time
	geo         string
	lastmod     string
	location    string
	organizer   string
	percent     string
	priority    string

	recurid string
	seq     string
	status  string
	summary string
	url     string

	rrule string

	due      string
	duration string

	attach     []string
	attendee   []string
	categories []string
	comment    []string
	contact    []string
	exdate     []string
	rstatus    []string
	related    []string
	resources  []string
	rdate      []string
	Xprop      []string
	ianaProp   []string
	Alarms     []Alarmc
}

func (t *Todoc) Build() {

}

//    Purpose:  Provide a grouping of component properties that describe a
//      journal entry.
//    Format Definition:  A "VJOURNAL" calendar component is defined by the
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
type Journalc struct {
	dtstamp     string
	uid         string
	class       string
	created     string
	dtstart     string
	lastMod     string
	organizer   string
	recurid     string
	seq         string
	status      string
	summary     string
	url         string
	rrule       string
	attach      []string
	attendee    []string
	categories  []string
	comment     []string
	contact     []string
	description string
	exdate      []string
	related     []string
	rdate       []string
	rstatus     []string
	Xprop       []string
	ianaProp    []string
}

func (j *Journalc) Build() {

}

//    Purpose:  Provide a grouping of component properties that describe
//      either a request for free/busy time, describe a response to a
//      request for free/busy time, or describe a published set of busy
//      time.
//
//   Format Definition:  A "VFREEBUSY" calendar component is defined by
//      the following notation:
//
//       freebusyc  = "BEGIN" ":" "VFREEBUSY" CRLF
//                    fbprop
//                    "END" ":" "VFREEBUSY" CRLF
//
//       fbprop     = *(
//                  ;
//                  ; The following are REQUIRED,
//                   ; but MUST NOT occur more than once.
//                  ;
//                  dtstamp / uid /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  contact / dtstart / dtend /
//                  organizer / url /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  attendee / comment / freebusy / rstatus / x-prop /
//                  iana-prop
//                  ;
//                  )
type Freebusyc struct {
	dtstamp   string
	uid       string
	contact   string
	dtstart   string
	dtend     string
	organizer string
	url       string
	attendee  string
	comment   string
	freebusy  string
	rstatus   string
	XProp     string
	IanaProp  string
}

func (f *Freebusyc) Build() {

}

type Timezonec struct {
}

func (t *Timezonec) build() {

}

type IanaComp struct {
}

func (i *IanaComp) Build() {

}

type Xcomp struct {
}

func (x *Xcomp) Build() {

}

//    Purpose:  Provide a grouping of component properties that define an
//      alarm.
//
//   Format Definition:  A "VALARM" calendar component is defined by the
//      following notation:
//
//       alarmc     = "BEGIN" ":" "VALARM" CRLF
//                    (audioprop / dispprop / emailprop)
//                    "END" ":" "VALARM" CRLF
//
//       audioprop  = *(
//                  ;
//                  ; 'action' and 'trigger' are both REQUIRED,
//                   ; but MUST NOT occur more than once.
//                  ;
//                  action / trigger /
//                  ;
//                  ; 'duration' and 'repeat' are both OPTIONAL,
//                  ; and MUST NOT occur more than once each;
//                  ; but if one occurs, so MUST the other.
//                  ;
//                  duration / repeat /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  attach /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  x-prop / iana-prop
//                  ;
//                  )
//
//       dispprop   = *(
//                  ;
//                  ; The following are REQUIRED,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  action / description / trigger /
//                  ;
//                  ; 'duration' and 'repeat' are both OPTIONAL,
//                  ; and MUST NOT occur more than once each;
//                  ; but if one occurs, so MUST the other.
//                  ;
//                  duration / repeat /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  x-prop / iana-prop
//                  ;
//                  )
//
//       emailprop  = *(
//                  ;
//                  ; The following are all REQUIRED,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  action / description / trigger / summary /
//                   ;
//                  ; The following is REQUIRED,
//                  ; and MAY occur more than once.
//                  ;
//                  attendee /
//                  ;
//                  ; 'duration' and 'repeat' are both OPTIONAL,
//                  ; and MUST NOT occur more than once each;
//                  ; but if one occurs, so MUST the other.
//                  ;
//                  duration / repeat /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  attach / x-prop / iana-prop
//                  ;
//                  )
type Alarmc struct {
	Prop AlarmProp
}

func (a *Alarmc) Build(s *strings.Builder) error {
	ok, err := a.Prop.Validate()
	if ok {
		s.WriteString("BEGIN:VALARM")
		a.Prop.Build(s)
		s.WriteString("END:VALARM")
		return nil

	}
	return err
}

type AlarmProp interface {
	Build(s *strings.Builder)
	Validate() (bool, error)
}

type Audioprop struct {
	action   string
	trigger  string
	duration string
	repeat   string
	attach   string
	xProp    string
	ianaProp string
}

func (a *Audioprop) Build(s *strings.Builder) {
	s.WriteString(fmt.Sprintln("ACTION:AUDIO"))
	s.WriteString(fmt.Sprintf("TRIGGER:%s\n", a.trigger))
	if a.duration != "" {
		s.WriteString(fmt.Sprintf("DURATION:%s\n", a.duration))
		s.WriteString(fmt.Sprintf("REPEAT:%s\n", a.repeat))
	}
	if a.attach != "" {
		s.WriteString(fmt.Sprintf("ATTACH:%s\n", a.attach))
	}
	if a.xProp != "" {
		s.WriteString(fmt.Sprintf("XPROP:%s\n", a.xProp))
	}
	if a.ianaProp != "" {
		s.WriteString(fmt.Sprintf("IANAPROP:%s\n", a.ianaProp))

	}
}

func (a *Audioprop) Validate() (bool, error) {
	if a.action == "" || a.trigger == "" {
		return false, fmt.Errorf("audioprop's action or trigger properity must not be null")
	}
	if a.duration != "" && a.repeat == "" {
		return false, fmt.Errorf("audioprop's duration and repeat must appear both")
	}
	if a.duration == "" && a.repeat != "" {
		return false, fmt.Errorf("audioprop's duration and repeat must appear both")
	}
	return true, nil
}

//       dispprop   = *(
//                  ;
//                  ; The following are REQUIRED,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  action / description / trigger /
//                  ;
//                  ; 'duration' and 'repeat' are both OPTIONAL,
//                  ; and MUST NOT occur more than once each;
//                  ; but if one occurs, so MUST the other.
//                  ;
//                  duration / repeat /
//                  ;
//                  ; The following is OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  x-prop / iana-prop
//                  ;
//                  )
type Disprop struct {
	action      string
	description string
	trigger     string
	duration    string
	repeat      int
	xProp       string
	ianaProp    string
}

func (d *Disprop) Build(s *strings.Builder) {
	s.WriteString(fmt.Sprintln("ACTION:DISPLY"))
	s.WriteString(fmt.Sprintf("DESCRIPTION:%s\n", d.description))
	s.WriteString(fmt.Sprintf("TRIGGER:%s", d.trigger))
	if d.duration != "" {
		s.WriteString(fmt.Sprintf("DURATION:%s", d.duration))
		s.WriteString(fmt.Sprintf("REPEAT:%d", d.repeat))
	}
	if d.xProp != "" {
		s.WriteString(fmt.Sprintf("XPROP:%s\n", d.xProp))
	}
	if d.ianaProp != "" {
		s.WriteString(fmt.Sprintf("IANAPROP:%s\n", d.ianaProp))

	}
}

func (d *Disprop) Validate() (bool, error) {
	if d.action == "" || d.trigger == "" || d.description == "" {
		return false, fmt.Errorf("disprop's action and trigger and description properity must not be null")
	}
	if d.duration != "" && d.repeat == 0 {
		return false, fmt.Errorf("disprop's duration and repeat must appear both")
	}
	if d.duration == "" && d.repeat != 0 {
		return false, fmt.Errorf("disprop's duration and repeat must appear both")
	}
	return true, nil
}

//       emailprop  = *(
//                  ;
//                  ; The following are all REQUIRED,
//                  ; but MUST NOT occur more than once.
//                  ;
//                  action / description / trigger / summary /
//                   ;
//                  ; The following is REQUIRED,
//                  ; and MAY occur more than once.
//                  ;
//                  attendee /
//                  ;
//                  ; 'duration' and 'repeat' are both OPTIONAL,
//                  ; and MUST NOT occur more than once each;
//                  ; but if one occurs, so MUST the other.
//                  ;
//                  duration / repeat /
//                  ;
//                  ; The following are OPTIONAL,
//                  ; and MAY occur more than once.
//                  ;
//                  attach / x-prop / iana-prop
//                  ;
//                  )
type Emailprop struct {
	action      string
	description string
	trigger     string
	summary     string
	attendee    []string
	duration    string
	repeat      int
	attach      []string
	xProp       []string
	ianaProp    []string
}

func (e *Emailprop) Build(s *strings.Builder) {
	s.WriteString(fmt.Sprintln("ACTION:EMAIL"))
	s.WriteString(fmt.Sprintf("DESCRIPTION:%s\n", e.description))
	s.WriteString(fmt.Sprintf("TRIGGER:%s\n", e.trigger))
	s.WriteString(fmt.Sprintf("SUMMARY:%s\n", e.summary))
	for _, a := range e.attendee {
		s.WriteString(fmt.Sprintf("ATTENDEE:%s\n", a))
	}
	if e.duration != "" {
		s.WriteString(fmt.Sprintf("DURATION:%s\n", e.duration))
		s.WriteString(fmt.Sprintf("REPEAT:%d\n", e.repeat))
	}
	if e.xProp != nil {
		s.WriteString(fmt.Sprintf("XPROP:%s\n", e.xProp))
	}
	if e.ianaProp != nil {
		s.WriteString(fmt.Sprintf("IANAPROP:%s\n", e.ianaProp))

	}
}
func (e *Emailprop) Validate() (bool, error) {
	if e.action == "" || e.trigger == "" || e.description == "" || e.summary == "" {
		return false, fmt.Errorf("emailprop's action and trigger and description and sunnary properity must not be null")
	}
	if e.attendee == nil {
		return false, fmt.Errorf("emailprop's attendee properties must not be specified")
	}
	if e.duration != "" && e.repeat == 0 {
		return false, fmt.Errorf("emailprop's duration and repeat must appear both")
	}
	if e.duration == "" && e.repeat != 0 {
		return false, fmt.Errorf("emailprop's duration and repeat must appear both")
	}
	return true, nil
}
