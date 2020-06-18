package model

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

const (
	UTCDateTimeFormat = "20060102T150405Z"
	UTCTimeFormat     = "150405Z"
)

type Attach interface {
	Attach() string
}

type Trigger struct {
	Related2Start bool
	Duration      *Duration
	DateTime      *DateTime
}

func (t *Trigger) Trigger() string {
	if t.Duration != nil {
		if t.Related2Start {
			return fmt.Sprintf(";VALUE=DURATION;RELATED=START:%s", t.Duration.String())
		}
		return fmt.Sprintf(";VALUE=DURATION;RELATED=END:%s", t.Duration.String())
	}
	return fmt.Sprintf(";VALUE=DATETIME:%s", t.DateTime.String())
}

//    Purpose:  This value type is used to identify properties that contain
//      a character encoding of inline binary data.  For example, an
//      inline attachment of a document might be included in an iCalendar
//      object.
//    Description:  Property values with this value type MUST also include
//      the inline encoding parameter sequence of ";ENCODING=BASE64".
//      That is, all inline binary data MUST first be character encoded
//      using the "BASE64" encoding method defined in [RFC2045].  No
//      additional content value encoding (i.e., BACKSLASH character
//      encoding, see Section 3.3.11) is defined for this value type.
type Binary struct {
	Data []byte
}

func (b *Binary) String() string {
	buf := new(bytes.Buffer)
	w := base64.NewEncoder(base64.StdEncoding, buf)
	_, err := w.Write(b.Data)
	if err != nil {
		return ""
	}
	_ = w.Close()
	bs := fmt.Sprintf(";ENCODING=BASE64;VALUE=BINARY:%s", buf.String())
	return bs
}
func (b *Binary) Attach() string {
	return b.String()
}

//    Purpose:  This value type is used to identify properties that contain
//      either a "TRUE" or "FALSE" Boolean value.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       boolean    = "TRUE" / "FALSE"
//
//   Description:  These values are case-insensitive text.  No additional
//      content value encoding (i.e., BACKSLASH character encoding, see
//      Section 3.3.11) is defined for this value type.
type Boolean struct {
	Data bool
}

func (b *Boolean) String() string {
	if b.Data {
		return "TRUE"
	}
	return "FALSE"
}

//    Purpose:  This value type is used to identify properties that contain
//      a calendar user address.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       cal-address        = uri
//
//   Description:  The value is a URI as defined by [RFC3986] or any other
//      IANA-registered form for a URI.  When used to address an Internet
//      email transport address for a calendar user, the value MUST be a
//      mailto URI, as defined by [RFC2368].  No additional content value
//      encoding (i.e., BACKSLASH character encoding, see Section 3.3.11)
//      is defined for this value type.
type CalAddress struct {
	Data string
}

func (c *CalAddress) String() string {
	return fmt.Sprintf("mailto:%s", c.Data)
}

//    Purpose:  This value type is used to identify values that contain a
//      calendar date.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       date               = date-value
//
//       date-value         = date-fullyear date-month date-mday
//       date-fullyear      = 4DIGIT
//       date-month         = 2DIGIT        ;01-12
//       date-mday          = 2DIGIT        ;01-28, 01-29, 01-30, 01-31
//                                          ;based on month/year
//
//   Description:  If the property permits, multiple "date" values are
//      specified as a COMMA-separated list of values.  The format for the
//      value type is based on the [ISO.8601.2004] complete
//      representation, basic format for a calendar date.  The textual
//      format specifies a four-digit year, two-digit month, and two-digit
//      day of the month.  There are no separator characters between the
//      year, month, and day component text.
//
//      No additional content value encoding (i.e., BACKSLASH character
//      encoding, see Section 3.3.11) is defined for this value type.
type Date struct {
	Date time.Time
}

func (d *Date) String() string {
	return d.Date.Format("20060102")
}

//    Purpose:  This value type is used to identify values that specify a
//      precise calendar date and time of day.
//    Format Definition:  This value type is defined by the following
//      notation:
//
//       date-time  = date "T" time ;As specified in the DATE and TIME
//                                  ;value definitions
type DateTime struct {
	Date time.Time
}

func (d *DateTime) String() string {
	// todo other date format style
	return d.Date.Format(UTCDateTimeFormat)
}

func (d *DateTime) Trigger() string {
	return fmt.Sprintf(";VALUE=DATETIME:%s", d.String())
}

func (d *DateTime) Related() string {
	return ""
}

//    Purpose:  This value type is used to identify properties that contain
//      a duration of time.
//    Format Definition:  This value type is defined by the following
//      notation:
//
//       dur-value  = (["+"] / "-") "P" (dur-date / dur-time / dur-week)
//
//       dur-date   = dur-day [dur-time]
//       dur-time   = "T" (dur-hour / dur-minute / dur-second)
//       dur-week   = 1*DIGIT "W"
//       dur-hour   = 1*DIGIT "H" [dur-minute]
//       dur-minute = 1*DIGIT "M" [dur-second]
//       dur-second = 1*DIGIT "S"
//       dur-day    = 1*DIGIT "D"

type Duration struct {
	Negative  bool
	DurWeek   int
	DurDay    int
	DurHour   int
	DurMinute int
	DurSecond int
}

func (d *Duration) Trigger() string {
	return fmt.Sprintf("%s%s", d.Related(), d.String())
}

func (d *Duration) Related() string {
	return fmt.Sprintf(";VALUE=DURATION;RELATED=START:")
}

func (d *Duration) String() string {
	s := strings.Builder{}
	if d.Negative {
		s.WriteString("-")
	}
	s.WriteString("P")
	if d.DurWeek != 0 {
		s.WriteString(fmt.Sprint(d.DurWeek))
		s.WriteString("W")
		return s.String()
	}
	if d.DurDay != 0 {
		s.WriteString(fmt.Sprint(d.DurDay))
		s.WriteString("D")
	}
	if d.DurHour != 0 || d.DurMinute != 0 || d.DurSecond != 0 {
		s.WriteString("T")
	}
	if d.DurHour != 0 {
		s.WriteString(fmt.Sprint(d.DurHour))
		s.WriteString("H")
	}
	if d.DurMinute != 0 {
		s.WriteString(fmt.Sprint(d.DurMinute))
		s.WriteString("M")
	}
	if d.DurSecond != 0 {
		s.WriteString(fmt.Sprint(d.DurSecond))
		s.WriteString("S")
	}
	return s.String()
}

//    Purpose:  This value type is used to identify properties that contain
//      a real-number value.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       float      = (["+"] / "-") 1*DIGIT ["." 1*DIGIT]
type Float float32

//    Purpose:  This value type is used to identify properties that contain
//      a signed integer value.
//
//   Format Definition:  This value type is defined by the following
//      notation:
type Integer int

//    Purpose:  This value type is used to identify values that contain a
//      precise period of time.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       period     = period-explicit / period-start
//
//       period-explicit = date-time "/" date-time
//       ; [ISO.8601.2004] complete representation basic format for a
//       ; period of time consisting of a start and end.  The start MUST
//       ; be before the end.
//
//       period-start = date-time "/" dur-value
//       ; [ISO.8601.2004] complete representation basic format for a
//       ; period of time consisting of a start and positive duration
//       ; of time.
// 
type Period interface {
	Check() bool
	Build() string
}
type ExplicitPeriod struct {
	Start time.Time
	End   time.Time
}

func (e *ExplicitPeriod) Check() bool {
	return e.Start.Before(e.End)
}

func (e *ExplicitPeriod) Build() string {
	return e.Start.Format(UTCDateTimeFormat) + "/" + e.End.Format(UTCDateTimeFormat)
}

type StartPeriod struct {
	Start    time.Time
	Duration Duration
}

func (s *StartPeriod) Check() bool {
	return true
}

func (s *StartPeriod) Build() string {
	return s.Start.Format(UTCDateTimeFormat) + "/" + s.Duration.String()
}

type Recur string
type Text string

//    Purpose:  This value type is used to identify values that contain a
//      time of day.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       time         = time-hour time-minute time-second [time-utc]
//
//       time-hour    = 2DIGIT        ;00-23
//       time-minute  = 2DIGIT        ;00-59
//       time-second  = 2DIGIT        ;00-60
//       ;The "60" value is used to account for positive "leap" seconds.
//
//       time-utc     = "Z"
type Time struct {
	Data time.Time
}

func (t *Time) String() string {
	return t.Data.Format(UTCTimeFormat)
}

//    Purpose:  This value type is used to identify values that contain a
//      uniform resource identifier (URI) type of reference to the
//      property value.
type URI struct {
	Uri string
}

func (u *URI) String() string {
	return u.Uri
}

func (u *URI) Attach() string {
	return fmt.Sprintf(":%s", u.String())
}

//    Purpose:  This value type is used to identify properties that contain
//      an offset from UTC to local time.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//
//       utc-offset = time-numzone
//
//       time-numzone = ("+" / "-") time-hour time-minute [time-second]
type UTCOffset string

func SplitString(s string) string {
	if len(s) <= 75 {
		return s
	}
	sb := strings.Builder{}
	var i = 1
	for ; i*70 < len(s); i++ {
		sb.WriteString(s[(i-1)*70 : i*70])
		sb.WriteString("\n ")
	}
	i--
	sb.WriteString(s[i*70:])
	return sb.String()
}
