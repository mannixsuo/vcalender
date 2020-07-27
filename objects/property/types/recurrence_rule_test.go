package types

import (
	"fmt"
	"testing"
)

func TestRecurRule_Value(t *testing.T) {
	//FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1
	r := RecurRule{
		Frequency: FreqMonthly,
		Rules: []Rule{
			&ByDay{V: []*WeekDayNum{{WeekDay: Monday,}, {WeekDay: Tuesday},
				{WeekDay: Wednesday}, {WeekDay: Thursday}, {WeekDay: Friday}}},
			&BySetpos{V: []*YearDayNum{{Operator: Minus, OrdYrDay: 1}}}},
	}
	if r.Value() != "FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1" {
		fmt.Println(r.Value())
		t.Error()
	}

	//FREQ=YEARLY;INTERVAL=2;BYMONTH=1;BYDAY=SU;BYHOUR=8,9;BYMINUTE=30
	r = RecurRule{
		Frequency: FreqYearly,
		Rules: []Rule{
			&Interval{V: 2},
			&ByMonth{V: []int{1}},
			&ByDay{V: []*WeekDayNum{{
				WeekDay: Sunday,
			}}},
			&ByHour{V: []int{8, 9}},
			&ByMinute{V: []int{30}},
		}}
	if r.Value() != "FREQ=YEARLY;INTERVAL=2;BYMONTH=1;BYDAY=SU;BYHOUR=8,9;BYMINUTE=30" {
		fmt.Println(r.Value())
		t.Error()
	}
}
