package header

import (
	"testing"

	apb "github.com/subiz/header/account"
)

func TestEvaluateDatetime(t *testing.T) {
	acc := &apb.Account{Timezone: new("+07:00")}
	// 2024-05-10 10:00:00 +07:00
	// UTC: 2024-05-10 03:00:00
	// Unix: 1715310000
	// May is month 5
	// 2024-05-10 is Friday (Weekday 5)
	unixms := int64(1715310000000)

	// Test month_eq
	cond := &DatetimeCondition{
		Op:      "month_eq",
		MonthEq: []int64{5},
	}
	if !EvaluateDatetime(acc, true, unixms, cond, 0) {
		t.Errorf("month_eq: expected true for month 5")
	}

	cond.MonthEq = []int64{6}
	if EvaluateDatetime(acc, true, unixms, cond, 0) {
		t.Errorf("month_eq: expected false for month 6")
	}

	// Test weekday_eq
	cond = &DatetimeCondition{
		Op:        "weekday_eq",
		WeekdayEq: []int64{5}, // Friday
	}
	if !EvaluateDatetime(acc, true, unixms, cond, 0) {
		t.Errorf("weekday_eq: expected true for Friday (5)")
	}

	cond.WeekdayEq = []int64{0} // Sunday
	if EvaluateDatetime(acc, true, unixms, cond, 0) {
		t.Errorf("weekday_eq: expected false for Sunday (0)")
	}
}
