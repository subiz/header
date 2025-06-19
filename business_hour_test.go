package header

import (
	"fmt"
	pb "github.com/subiz/header/account"
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	tcs := []struct {
		holidays  []*pb.BusinessHours_Holiday
		now       string
		tz        string
		isholiday bool
	}{
		{
			[]*pb.BusinessHours_Holiday{{
				Year:  PI32(2018),
				Month: PI32(8),
				Day:   PI32(21),
			}}, "2018-08-21T15:04:00Z", "+07:00", true,
		}, {
			[]*pb.BusinessHours_Holiday{{
				Year:  PI32(2018),
				Month: PI32(8),
				Day:   PI32(21),
			}}, "2018-08-21T23:04:00Z", "+07:00", false,
		}, {
			[]*pb.BusinessHours_Holiday{{
				Year:  PI32(2018),
				Month: PI32(8),
				Day:   PI32(22),
			}, {
				Year:  PI32(2018),
				Month: PI32(8),
				Day:   PI32(21),
			}}, "2018-08-21T23:04:00Z", "+07:00", true,
		},
	}

	for _, tc := range tcs {
		tim, err := time.Parse(time.RFC3339, tc.now)
		if err != nil {
			t.Fatalf("%s: %v", tc.now, err)
		}

		isholiday, err := IsHoliday(&pb.BusinessHours{Holidays: tc.holidays}, tim, tc.tz)
		if err != nil {
			t.Fatalf("%s: %v", tc.now, err)
		}

		if isholiday != tc.isholiday {
			t.Errorf("%s: should be %v, got %v", tc.now, tc.isholiday, isholiday)
		}
	}
}

func TestDuringBusinessHour(t *testing.T) {
	tcs := []struct {
		holidays    []*pb.BusinessHours_Holiday
		workingdays []*pb.BusinessHours_WorkingDay
		now         string
		tz          string
		in          bool
	}{
		{
			[]*pb.BusinessHours_Holiday{{
				Year:  PI32(2018),
				Month: PI32(8),
				Day:   PI32(21),
			}}, []*pb.BusinessHours_WorkingDay{{}}, "2018-08-21T15:04:00Z", "+07:00", false,
		},
		{
			nil, []*pb.BusinessHours_WorkingDay{{
				Weekday:   S("Wednesday"),
				StartTime: S("08:30"),
				EndTime:   S("23:30"),
			}}, "2018-08-21T05:04:00Z", "+07:00", false, // Tuesday
		},
		{
			nil, []*pb.BusinessHours_WorkingDay{{
				Weekday:   S("Tuesday"),
				StartTime: S("08:30"),
				EndTime:   S("23:30"),
			}}, "2018-08-21T05:04:00Z", "+07:00", true, // Tuesday
		},
		{
			nil, []*pb.BusinessHours_WorkingDay{{
				Weekday:   S("Tuesday"),
				StartTime: S("08:30"),
				EndTime:   S("23:30"),
			}}, "2018-08-21T23:04:00Z", "+07:00", false, // Tuesday
		},
		{
			nil, []*pb.BusinessHours_WorkingDay{{
				Weekday:   S("Thursday"),
				StartTime: S("10:30"),
				EndTime:   S("11:00"),
			}}, "2019-02-28T04:30:39Z", "+07:00", false, // Thursday
		},
	}

	for _, tc := range tcs {
		tim, err := time.Parse(time.RFC3339, tc.now)
		if err != nil {
			t.Fatalf("%s: %v", tc.now, err)
		}

		duringbusiness, err := DuringBusinessHour(&pb.BusinessHours{
			WorkingDays: tc.workingdays,
			Holidays:    tc.holidays,
		}, tim, tc.tz)
		if err != nil {
			t.Fatalf("%s: %v", tc.now, err)
		}

		if duringbusiness != tc.in {
			t.Errorf("%s: should be %v, got %v", tc.now, tc.in, duringbusiness)
		}
	}
}

func PI32(i int) *int32 {
	i32 := int32(i)
	return &i32
}

func S(s interface{}) *string {
	if s == nil {
		return S("")
	}
	switch v := s.(type) {
	case []byte:
		b := string(v)
		return &b
	case string:
		return &v
	case fmt.Stringer:
		str := v.String()
		return &str
	default:
		str := fmt.Sprintf("%v", v)
		return &str
	}
}
