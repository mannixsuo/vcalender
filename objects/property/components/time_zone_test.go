package components

import (
	"calendar/objects/property/components/properties/changemanage"
	"calendar/objects/property/components/properties/datetime"
	"calendar/objects/property/components/properties/recurrence"
	"calendar/objects/property/components/properties/timezone"
	"calendar/objects/property/types"
	"testing"
)

func TestTimeZone_TimeZone(t *testing.T) {
	//       BEGIN:VTIMEZONE
	//       TZID:America/New_York
	//       LAST-MODIFIED:20050809T050000Z
	//       BEGIN:DAYLIGHT
	//       DTSTART:19670430T020000
	//       RRULE:FREQ=YEARLY;BYMONTH=4;BYDAY=-1SU;UNTIL=19730429T070000Z
	//       TZOFFSETFROM:-0500
	//       TZOFFSETTO:-0400
	//       TZNAME:EDT
	//       END:DAYLIGHT
	//       BEGIN:STANDARD
	//       DTSTART:19671029T020000
	//       RRULE:FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU;UNTIL=20061029T060000Z
	//       TZOFFSETFROM:-0400
	//       TZOFFSETTO:-0500
	//       TZNAME:EST
	//       END:STANDARD
	//       BEGIN:DAYLIGHT
	//       DTSTART:19740106T020000
	//       RDATE:19750223T020000
	//       TZOFFSETFROM:-0500
	//       TZOFFSETTO:-0400
	//       TZNAME:EDT
	//       END:DAYLIGHT
	//       BEGIN:DAYLIGHT
	//       DTSTART:19760425T020000
	//       RRULE:FREQ=YEARLY;BYMONTH=4;BYDAY=-1SU;UNTIL=19860427T070000Z
	//       TZOFFSETFROM:-0500
	//       TZOFFSETTO:-0400
	//       TZNAME:EDT
	//       END:DAYLIGHT
	//       BEGIN:DAYLIGHT
	//       DTSTART:19870405T020000
	//       RRULE:FREQ=YEARLY;BYMONTH=4;BYDAY=1SU;UNTIL=20060402T070000Z
	//       TZOFFSETFROM:-0500
	//       TZOFFSETTO:-0400
	//       TZNAME:EDT
	//       END:DAYLIGHT
	//       BEGIN:DAYLIGHT
	//       DTSTART:20070311T020000
	//       RRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=2SU
	//       TZOFFSETFROM:-0500
	//       TZOFFSETTO:-0400
	//       TZNAME:EDT
	//       END:DAYLIGHT
	//       BEGIN:STANDARD
	//       DTSTART:20071104T020000
	//       RRULE:FREQ=YEARLY;BYMONTH=11;BYDAY=1SU
	//       TZOFFSETFROM:-0400
	//       TZOFFSETTO:-0500
	//       TZNAME:EST
	//       END:STANDARD
	//       END:VTIMEZONE
	z := TimeZone{
		TzId:    timezone.NewTzId("America/New_York"),
		LastMod: changemanage.NewLastModified(2005, 8, 9, 5, 0, 0),
		Standard: []*Standard{
			{&TzProp{
				DtStart:      datetime.NewDateStartWithDatetimeAndFormat(1967, 10, 29, 2, 0, 0, types.LocalDateTimeFormat),
				TzOffsetTo:   timezone.NewTzOffsetTo(false, 5, 0, 0),
				TzOffsetFrom: timezone.NewTzOffsetFrom(false, 4, 0, 0),
				RRule: &recurrence.RRule{
					V: &types.RecurRule{
						Frequency: types.FreqYearly,
						Rules: []types.Rule{
							&types.ByMonth{V: []int{10}},
							&types.ByDay{V: []*types.WeekDayNum{{
								Operator: types.Minus,
								OrdWk:    1,
								WeekDay:  types.Sunday,
							}}},
							&types.Until{Time: types.NewUTCDateTime(2006, 10, 29, 6, 0, 0)},
						},
					},
				},
				TzName: []*timezone.TzName{timezone.NewTzName("EST")},
			}},
			{&TzProp{
				DtStart:      datetime.NewDateStartWithDatetimeAndFormat(2007, 11, 4, 2, 0, 0, types.LocalDateTimeFormat),
				TzOffsetTo:   timezone.NewTzOffsetTo(false, 5, 0, 0),
				TzOffsetFrom: timezone.NewTzOffsetFrom(false, 4, 0, 0),
				RRule: &recurrence.RRule{
					V: &types.RecurRule{
						Frequency: types.FreqYearly,
						Rules: []types.Rule{
							&types.ByMonth{V: []int{11}},
							&types.ByDay{V: []*types.WeekDayNum{{
								OrdWk:   1,
								WeekDay: types.Sunday,
							}}},
						},
					},
				},
				TzName: []*timezone.TzName{timezone.NewTzName("EST")},
			}},
		},
		DayLight: []*DayLight{
			{&TzProp{
				DtStart: datetime.NewDateStartWithDatetimeAndFormat(1967, 4, 30, 2, 0, 0, types.LocalDateTimeFormat),
				RRule: &recurrence.RRule{
					V: &types.RecurRule{
						Frequency: types.FreqYearly,
						Rules: []types.Rule{
							&types.ByMonth{V: []int{4}},
							&types.ByDay{V: []*types.WeekDayNum{{
								Operator: types.Minus,
								OrdWk:    1,
								WeekDay:  types.Sunday,
							}}},
							&types.Until{Time: types.NewUTCDateTime(1973, 4, 29, 7, 0, 0)},
						},
					},
				},
				TzOffsetFrom: timezone.NewTzOffsetFrom(false, 5, 0, 0),
				TzOffsetTo:   timezone.NewTzOffsetTo(false, 4, 0, 0),
				TzName:       []*timezone.TzName{timezone.NewTzName("EDT")},
			}},
			{&TzProp{
				DtStart: datetime.NewDateStartWithDatetimeAndFormat(1974, 1, 6, 2, 0, 0, types.LocalDateTimeFormat),
				RDate: []*recurrence.RDate{
					{
						Values: []types.Value{types.NewLocalDateTime(1975, 2, 23, 2, 0, 0)},
					},
				},
				RRule: &recurrence.RRule{
					V: &types.RecurRule{
						Frequency: types.FreqYearly,
						Rules: []types.Rule{
							&types.ByMonth{V: []int{4}},
							&types.ByDay{V: []*types.WeekDayNum{{
								Operator: types.Minus,
								OrdWk:    1,
								WeekDay:  types.Sunday,
							}}},
							&types.Until{Time: types.NewUTCDateTime(1973, 4, 29, 7, 0, 0)},
						},
					},
				},
				TzOffsetFrom: timezone.NewTzOffsetFrom(false, 5, 0, 0),
				TzOffsetTo:   timezone.NewTzOffsetTo(false, 4, 0, 0),
				TzName:       []*timezone.TzName{timezone.NewTzName("EDT")},
			}},
			{&TzProp{
				DtStart: datetime.NewDateStartWithDatetimeAndFormat(1976, 4, 25, 2, 0, 0, types.LocalDateTimeFormat),
				RRule: &recurrence.RRule{
					V: &types.RecurRule{
						Frequency: types.FreqYearly,
						Rules: []types.Rule{
							&types.ByMonth{V: []int{4}},
							&types.ByDay{V: []*types.WeekDayNum{{
								Operator: types.Minus,
								OrdWk:    1,
								WeekDay:  types.Sunday,
							}}},
							&types.Until{Time: types.NewUTCDateTime(1986, 4, 27, 7, 0, 0)},
						},
					},
				},
				TzOffsetFrom: timezone.NewTzOffsetFrom(false, 5, 0, 0),
				TzOffsetTo:   timezone.NewTzOffsetTo(false, 4, 0, 0),
				TzName:       []*timezone.TzName{timezone.NewTzName("EDT")},
			}},
			{&TzProp{
				DtStart: datetime.NewDateStartWithDatetimeAndFormat(1987, 4, 5, 2, 0, 0, types.LocalDateTimeFormat),
				RRule: &recurrence.RRule{
					V: &types.RecurRule{
						Frequency: types.FreqYearly,
						Rules: []types.Rule{
							&types.ByMonth{V: []int{4}},
							&types.ByDay{V: []*types.WeekDayNum{{
								OrdWk:   1,
								WeekDay: types.Sunday,
							}}},
							&types.Until{Time: types.NewUTCDateTime(2006, 4, 2, 7, 0, 0)},
						},
					},
				},
				TzOffsetFrom: timezone.NewTzOffsetFrom(false, 5, 0, 0),
				TzOffsetTo:   timezone.NewTzOffsetTo(false, 4, 0, 0),
				TzName:       []*timezone.TzName{timezone.NewTzName("EDT")},
			}},
			{&TzProp{
				DtStart: datetime.NewDateStartWithDatetimeAndFormat(2007, 3, 1, 2, 0, 0, types.LocalDateTimeFormat),
				RRule: &recurrence.RRule{
					V: &types.RecurRule{
						Frequency: types.FreqYearly,
						Rules: []types.Rule{
							&types.ByMonth{V: []int{3}},
							&types.ByDay{V: []*types.WeekDayNum{{
								OrdWk:   2,
								WeekDay: types.Sunday,
							}}},
						},
					},
				},
				TzOffsetFrom: timezone.NewTzOffsetFrom(false, 5, 0, 0),
				TzOffsetTo:   timezone.NewTzOffsetTo(false, 4, 0, 0),
				TzName:       []*timezone.TzName{timezone.NewTzName("EDT")},
			}},
		},
	}
	v := VTest{
		S: z.TimeZone(),
		T: t,
	}
	v.ContainsOrError("BEGIN:VTIMEZONE")
	v.ContainsOrError("TZID:America/New_York")
	v.ContainsOrError("LAST-MODIFIED:20050809T050000Z")
	v.ContainsOrError("BEGIN:DAYLIGHT")
	v.ContainsOrError("DTSTART:19670430T020000")
	v.ContainsOrError("RRULE:FREQ=YEARLY;BYMONTH=4;BYDAY=-1SU;UNTIL=19730429T070000Z")
	v.ContainsOrError("TZOFFSETFROM:-0500")
	v.ContainsOrError("TZOFFSETTO:-0400")
	v.ContainsOrError("TZNAME:EDT")
	v.ContainsOrError("END:DAYLIGHT")
	v.ContainsOrError("BEGIN:STANDARD")
	v.ContainsOrError("DTSTART:19671029T020000")
	v.ContainsOrError("RRULE:FREQ=YEARLY;BYMONTH=10;BYDAY=-1SU;UNTIL=20061029T060000Z")
	v.ContainsOrError("TZOFFSETFROM:-0400")
	v.ContainsOrError("TZOFFSETTO:-0500")
	v.ContainsOrError("TZNAME:EST")
	v.ContainsOrError("END:STANDARD")
	v.ContainsOrError("BEGIN:DAYLIGHT")
	v.ContainsOrError("DTSTART:19740106T020000")
	v.ContainsOrError("RDATE:19750223T020000")
	v.ContainsOrError("TZOFFSETFROM:-0500")
	v.ContainsOrError("TZOFFSETTO:-0400")
	v.ContainsOrError("TZNAME:EDT")
	v.ContainsOrError("END:DAYLIGHT")
	v.ContainsOrError("BEGIN:DAYLIGHT")
	v.ContainsOrError("DTSTART:19760425T020000")
	v.ContainsOrError("RRULE:FREQ=YEARLY;BYMONTH=4;BYDAY=-1SU;UNTIL=19860427T070000Z")
	v.ContainsOrError("TZOFFSETFROM:-0500")
	v.ContainsOrError("TZOFFSETTO:-0400")
	v.ContainsOrError("TZNAME:EDT")
	v.ContainsOrError("END:DAYLIGHT")
	v.ContainsOrError("BEGIN:DAYLIGHT")
	v.ContainsOrError("DTSTART:20071104T020000")
	v.ContainsOrError("RRULE:FREQ=YEARLY;BYMONTH=3;BYDAY=2SU")
	v.ContainsOrError("TZOFFSETFROM:-0500")
	v.ContainsOrError("TZOFFSETTO:-0400")
	v.ContainsOrError("TZNAME:EDT")
	v.ContainsOrError("END:DAYLIGHT")
	v.ContainsOrError("BEGIN:STANDARD")
	v.ContainsOrError("DTSTART:20071104T020000")
	v.ContainsOrError("RRULE:FREQ=YEARLY;BYMONTH=11;BYDAY=1SU")
	v.ContainsOrError("TZOFFSETFROM:-0400")
	v.ContainsOrError("TZOFFSETTO:-0500")
	v.ContainsOrError("TZNAME:EST")
	v.ContainsOrError("END:STANDARD")
	v.ContainsOrError("END:VTIMEZONE")

}
