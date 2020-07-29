package objects

import (
	"calendar/objects/property/components/properties"
)

//     contentline   = name *(";" param ) ":" value CRLF
//     name          = iana-token / x-name
//     iana-token    = 1*(ALPHA / DIGIT / "-")
//     ; iCalendar identifier registered with IANA
//     x-name        = "X-" [vendorid "-"] 1*(ALPHA / DIGIT / "-")
//     ; Reserved for experimental use.
//     vendorid      = 3*(ALPHA / DIGIT)
//     ; Vendor identification
//     param         = param-name "=" param-value *("," param-value)
//     ; Each property defines the specific ABNF for the parameters
//     ; allowed on the property.  Refer to specific properties for
//     ; precise parameter ABNF.
//     param-name    = iana-token / x-name
//
//     param-value   = paramtext / quoted-string
//
//     paramtext     = *SAFE-CHAR
//     value         = *VALUE-CHAR
//
//     quoted-string = DQUOTE *QSAFE-CHAR DQUOTE
//     QSAFE-CHAR    = WSP / %x21 / %x23-7E / NON-US-ASCII
//     ; Any character except CONTROL and DQUOTE
//
//     SAFE-CHAR     = WSP / %x21 / %x23-2B / %x2D-39 / %x3C-7E
//                   / NON-US-ASCII
//     ; Any character except CONTROL, DQUOTE, ";", ":", ","
//     VALUE-CHAR    = WSP / %x21-7E / NON-US-ASCII
//     ; Any textual character
//
//     NON-US-ASCII  = UTF8-2 / UTF8-3 / UTF8-4
//     ; UTF8-2, UTF8-3, and UTF8-4 are defined in [RFC3629]
//
//     CONTROL       = %x00-08 / %x0A-1F / %x7F
//     ; All the controls except HTAB

func ToContentLine(p properties.Property) string {
	return ""
}
