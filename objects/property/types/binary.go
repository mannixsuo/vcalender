package types

import (
	"bytes"
	"encoding/base64"
	"fmt"
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
type Binary []byte

func (b *Binary) Binary() string {
	buf := new(bytes.Buffer)
	w := base64.NewEncoder(base64.StdEncoding, buf)
	_, err := w.Write(*b)
	if err != nil {
		return ""
	}
	_ = w.Close()
	bs := fmt.Sprintf("%s", buf.String())
	return bs
}
