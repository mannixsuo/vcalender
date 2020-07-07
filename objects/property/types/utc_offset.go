package types

import "fmt"

//   Value Name:  UTC-OFFSET
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

func (u *UTCOffset) UTCOffset() (string, error) {
	err := u.check()
	if err != nil {
		return "", err
	}
	if u.Positive {
		return fmt.Sprintf("%2d%2d%2d", u.Hour, u.Minute, u.Second), nil
	}
	return fmt.Sprintf("-%2d%2d%2d", u.Hour, u.Minute, u.Second), nil
}

func (u *UTCOffset) check() error {
	if u.Second >= 60 || u.Minute >= 60 || u.Hour >= 60 {
		return fmt.Errorf("utc offset time must less than 60")
	}
	if u.Second == 0 && u.Minute == 0 && u.Hour == 0 && u.Positive == false {
		return fmt.Errorf("-000000 utc offset are not allowed")
	}
	return nil
}
