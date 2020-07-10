package alarm

import (
	"calendar/objects/property/components/properties"
	"calendar/objects/property/parameters"
	"calendar/objects/property/types"
)

//   Property Name:  ACTION
//
//   Purpose:  This property defines the action to be invoked when an
//      alarm is triggered.
//
//   Value Type:  TEXT
//
//   Property Parameters:  IANA and non-standard property parameters can
//      be specified on this property.
//
//   Conformance:  This property MUST be specified once in a "VALARM"
//      calendar component.
//
//   Description:  Each "VALARM" calendar component has a particular type
//      of action with which it is associated.  This property specifies
//      the type of action.  Applications MUST ignore alarms with x-name
//      and iana-token values they don't recognize.
//
//   Format Definition:  This property is defined by the following
//      notation:
//
//       action      = "ACTION" actionparam ":" actionvalue CRLF
//
//       actionparam = *(";" other-param)
//
//
//       actionvalue = "AUDIO" / "DISPLAY" / "EMAIL"
//                   / iana-token / x-name
//
//   Example:  The following are examples of this property in a "VALARM"
//      calendar component:
//       ACTION:AUDIO
//
//       ACTION:DISPLAY

type Action struct {
	Parameters []parameters.Parameter
	Value      *types.Text
}

func (a *Action) Property() (string, error) {
	return properties.DefaultCreatePropertyFunc("ACTION", a.Parameters, a.Value), nil
}

var Audio = Action{Value: &types.Text{V: "AUDIO"},}
var Display = Action{Value: &types.Text{V: "DISPLAY"},}
var Email = Action{Value: &types.Text{V: "EMAIL"},}
