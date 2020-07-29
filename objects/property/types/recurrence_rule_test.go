package types

import (
	"fmt"
	"strings"
	"testing"
)

func TestRecurRule_Value(t *testing.T) {
	b := &strings.Builder{}

	//FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1
	r := RecurRule{
		Frequency: FreqMonthly,
		Rules: []Rule{
			&ByDay{V: []*WeekDayNum{{WeekDay: Monday,}, {WeekDay: Tuesday},
				{WeekDay: Wednesday}, {WeekDay: Thursday}, {WeekDay: Friday}}},
			&BySetpos{V: []*YearDayNum{{Operator: Minus, OrdYrDay: 1}}}},
	}
	r.WriteValueToStrBuilder(b)
	if b.String() != "FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1" {
		fmt.Println(b.String())
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
	b.Reset()
	r.WriteValueToStrBuilder(b)
	if b.String() != "FREQ=YEARLY;INTERVAL=2;BYMONTH=1;BYDAY=SU;BYHOUR=8,9;BYMINUTE=30" {
		fmt.Println(b.String())
		t.Error()
	}
}
