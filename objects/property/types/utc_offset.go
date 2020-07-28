package types

import (
	"fmt"
	"strings"
)

//   V Name:  UTC-OFFSET
//
//   Purpose:  This value type is used to identify properties that contain
//      an offset from UTC to local time.
//
//   Format Definition:  This value type is defined by the following
//      notation:
//       utc-offset = time-numzone
//       time-numzone = ("+" / "-") time-hour time-minute [time-second]
//   Description:  The PLUS SIGN character MUST be specified for positive
//      UTC offsets (i.e., ahead of UTC).  The HYPHEN-MINUS character MUST
//      be specified for negative UTC offsets (i.e., behind of UTC).  The
//      value of "-0000" and "-000000" are not allowed.  The time-second,
//      if present, MUST NOT be 60; if absent, it defaults to zero.
//
//      No additional content value encoding (i.e., BACKSLASH character
//      encoding, see Section 3.3.11) is defined for this value type.
//
//   Example:  The following UTC offsets are given for standard time for
//      New York (five hours behind UTC) and Geneva (one hour ahead of
//      UTC):
//
//       -0500
//
//       +0100

type UTCOffset struct {
	Positive bool
	Hour     int
	Minute   int
	Second   int
}

func (u *UTCOffset) Value() string {
	s := strings.Builder{}
	if !u.Positive {
		s.WriteString("-")
	}
	s.WriteString(fmt.Sprintf("%02d", u.Hour))
	s.WriteString(fmt.Sprintf("%02d", u.Minute))
	if u.Second != 0 {
		s.WriteString(fmt.Sprintf("%02d", u.Second))
	}
	return s.String()
}

func NewUTCOffset(positive bool, hour, minute, second int) *UTCOffset {
	return &UTCOffset{
		Positive: positive,
		Hour:     hour,
		Minute:   minute,
		Second:   second,
	}
}
