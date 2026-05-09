package header

import (
	"testing"

	apb "github.com/subiz/header/account"
)

func TestToWorkflowCondition(t *testing.T) {
	cond := &UserViewCondition{
		Id:   "test-id",
		Key:  "test-key",
		Type: "text",
		Text: &TextCondition{
			Op: "eq",
			Eq: []string{"value"},
		},
		All: []*UserViewCondition{
			{
				Id:   "sub-id",
				Key:  "sub-key",
				Type: "boolean",
				Boolean: &BooleanCondition{
					Op: "true",
				},
			},
		},
	}

	wcond := ToWorkflowCondition(cond)

	if wcond.Id != cond.Id {
		t.Errorf("Id mismatch: expected %s, got %s", cond.Id, wcond.Id)
	}
	if wcond.Key != cond.Key {
		t.Errorf("Key mismatch: expected %s, got %s", cond.Key, wcond.Key)
	}
	if wcond.Type != cond.Type {
		t.Errorf("Type mismatch: expected %s, got %s", cond.Type, wcond.Type)
	}

	if wcond.Text == nil || wcond.Text.Op != cond.Text.Op || wcond.Text.Eq[0] != cond.Text.Eq[0] {
		t.Errorf("Text condition mismatch")
	}

	if len(wcond.All) != 1 {
		t.Errorf("All slice length mismatch: expected 1, got %d", len(wcond.All))
	} else {
		sub := wcond.All[0]
		expectedSub := cond.All[0]
		if sub.Id != expectedSub.Id {
			t.Errorf("Sub Id mismatch: expected %s, got %s", expectedSub.Id, sub.Id)
		}
		if sub.Boolean == nil || sub.Boolean.Op != expectedSub.Boolean.Op {
			t.Errorf("Sub Boolean condition mismatch")
		}
	}

	// Verify deep copy
	cond.Text.Eq[0] = "changed"
	if wcond.Text.Eq[0] == "changed" {
		t.Errorf("Expected deep copy, but wcond.Text changed when cond.Text changed")
	}
}

func strPtr(s string) *string {
	return &s
}

func TestEvaluateDatetime(t *testing.T) {
	acc := &apb.Account{Timezone: strPtr("+07:00")}
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
