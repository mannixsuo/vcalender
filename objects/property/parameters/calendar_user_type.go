package parameters

import "fmt"

//   Parameter Name:  CUTYPE
//
//   Purpose:  To identify the type of calendar user specified by the
//      property.
//
//   Format Definition:  This property parameter is defined by the
//      following notation:
//
//       cutypeparam        = "CUTYPE" "="
//                          ("INDIVIDUAL"   ; An individual
//                         / "GROUP"        ; A group of individuals
//                         / "RESOURCE"     ; A physical resource
//                         / "ROOM"         ; A room resource
//                         / "UNKNOWN"      ; Otherwise not known
//                         / x-name         ; Experimental type
//                         / iana-token)    ; Other IANA-registered
//                                          ; type
//       ; Default is INDIVIDUAL
//
//   Description:  This parameter can be specified on properties with a
//      CAL-ADDRESS value type.  The parameter identifies the type of
//      calendar user specified by the property.  If not specified on a
//      property that allows this parameter, the default is INDIVIDUAL.
//      Applications MUST treat x-name and iana-token values they don't
//      recognize the same way as they would the UNKNOWN value.
//
//   Example:
//
//       ATTENDEE;CUTYPE=GROUP:mailto:ietf-calsch@example.org

type CuType struct {
	V string
}

var IndividualCuType = CuType{"INDIVIDUAL"}
var GroupCuType = CuType{"GROUP"}
var ResourceCuType = CuType{"RESOURCE"}
var RoomCuType = CuType{"ROOM"}
var UnknownCuType = CuType{"UNKNOWN"}

func (c *CuType) Parameter() string {
	return fmt.Sprintf("CUTYPE=%s", c.V)
}
