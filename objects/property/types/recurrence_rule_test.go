package types

import (
	"fmt"
	"testing"
)

func TestRecurRule_Value(t *testing.T) {
	//FREQ=YEARLY;INTERVAL=2;BYMONTH=1;BYDAY=SU;BYHOUR=8,9;BYMINUTE=30
	r := RecurRule{Rules: []Rule{
		FreqYearly,
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
