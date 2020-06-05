package model

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"time"
)

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

type DateTime struct {
	Date time.Time
}

func (d *DateTime) String() string {
	return d.Date.Format()
}

type DURATION time.Time
type FLOAT float32

type INTEGER int
type PERIOD string
type RECUR string
type TEXT string
type TIME time.Time
type URI string
type UTCOFFSET string
