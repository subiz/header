package header

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	pb "github.com/subiz/header/account"
)

func DuringBusinessHour(bh *pb.BusinessHours, date time.Time, tz string) (bool, error) {
	isholiday, err := IsHoliday(bh, date, tz)
	if err != nil {
		return false, err
	}
	if isholiday {
		return false, nil
	}

	if len(bh.GetWorkingDays()) == 0 {
		return true, nil
	}

	_, _, _, h, m, weekday, err := ConvertTimezone(date, tz)
	if err != nil {
		return false, err
	}
	currentmin := h*60 + m

	for _, wd := range bh.GetWorkingDays() {
		if wd.GetWeekday() == weekday || wd.GetWeekday() == "Everyday" {
			start, err := toMinute(wd.GetStartTime())
			if err != nil {
				return false, err
			}
			end, err := toMinute(wd.GetEndTime())
			if err != nil {
				return false, err
			}

			if start <= currentmin && currentmin <= end {
				return true, nil
			}
		}
	}
	return false, nil
}

// IsHoliday tells where ther date (within timezone tz) is in holiday list
func IsHoliday(bh *pb.BusinessHours, date time.Time, tz string) (bool, error) {
	y, m, d, _, _, _, err := ConvertTimezone(date, tz)
	if err != nil {
		return false, err
	}
	for _, h := range bh.GetHolidays() {
		if int32(y) == h.GetYear() && int32(m) == h.GetMonth() &&
			int32(d) == h.GetDay() {
			return true, nil
		}
	}
	return false, nil
}

// lack of testcases
func GetNextBusinessMili(bh *pb.BusinessHours, date time.Time, tz string) (int64, error) {
	currentMin := date.Unix() / 60
	isInBussinessHour, err := DuringBusinessHour(bh, date, tz)
	if err != nil {
		return 0, nil
	}
	if isInBussinessHour {
		return date.UnixMilli(), nil
	}

	year, month, day, h, m, weekday, err := ConvertTimezone(date, tz)
	if err != nil {
		return 0, err
	}
	startMinCurrentDay := currentMin - int64(h)*60 - int64(m)
	bussesPeriodMins := [][]int64{}
	for _, wd := range bh.GetWorkingDays() {
		// get day diff not match holiday

		daydiffs := []int64{}
		if wd.GetWeekday() == "Everyday" {
			daydiffs = append(daydiffs, 0)
			daydiffs = append(daydiffs, 1)
			daydiffs = append(daydiffs, 2)
			daydiffs = append(daydiffs, 3)
			daydiffs = append(daydiffs, 4)
			daydiffs = append(daydiffs, 5)
			daydiffs = append(daydiffs, 6)
			// append first day of next week it to handle edge case: same day check
			daydiffs = append(daydiffs, 7)
		} else {
			dayDiff := getDiffWeekDay(weekday, wd.GetWeekday())
			daydiffs = append(daydiffs, dayDiff)
			// append next week it self to handle edge case: only 1 weekday
			daydiffs = append(daydiffs, dayDiff+7)
		}

		for _, daydiff := range daydiffs {
			// make sure next day is not match holiday
			for i := 0; i <= 10; i++ {
				tempDay := time.Date(year, time.Month(month), day+int(daydiff), 0, 0, 0, 0, time.UTC)
				isHoliday, err := IsHoliday(bh, tempDay, tz)
				if err != nil {
					return 0, err
				}
				if !isHoliday {
					break
				}
				daydiff += 7
			}
			startMinofDay, err := toMinute(wd.GetStartTime())
			if err != nil {
				return 0, err
			}
			endMinOfDay, err := toMinute(wd.GetEndTime())
			if err != nil {
				return 0, err
			}

			startBussinessMin := int64(startMinofDay) + daydiff*24*60 + startMinCurrentDay
			endMin := int64(endMinOfDay) + daydiff*24*60 + startMinCurrentDay
			bussesPeriodMins = append(bussesPeriodMins, []int64{startBussinessMin, endMin})
		}

	}
	var nextBussinessMin int64 = math.MaxInt64
	for _, period := range bussesPeriodMins {
		start := period[0]
		if start >= currentMin && start <= nextBussinessMin {
			nextBussinessMin = start
		}

	}
	nextBussinessMili := nextBussinessMin * 60 * 1000
	return nextBussinessMili, nil
}

// lack of testcases
func GetNextNonBusinessMili(bh *pb.BusinessHours, date time.Time, tz string) (int64, error) {
	currentMin := date.Unix() / 60
	isInBussinessHour, err := DuringBusinessHour(bh, date, tz)
	if err != nil {
		return 0, nil
	}
	if !isInBussinessHour {
		return date.UnixMilli(), nil
	}

	year, month, day, h, m, weekday, err := ConvertTimezone(date, tz)
	if err != nil {
		return 0, err
	}
	startMinCurrentDay := currentMin - int64(h)*60 - int64(m)
	bussesPeriodMins := [][]int64{}
	var maxEndMin int64 = 0
	for _, wd := range bh.GetWorkingDays() {
		daydiffs := []int64{}
		if wd.GetWeekday() == "Everyday" {
			daydiffs = append(daydiffs, 0)
			daydiffs = append(daydiffs, 1)
			daydiffs = append(daydiffs, 2)
			daydiffs = append(daydiffs, 3)
			daydiffs = append(daydiffs, 4)
			daydiffs = append(daydiffs, 5)
			daydiffs = append(daydiffs, 6)
			// append first day of next week it to handle edge case: same day check
			daydiffs = append(daydiffs, 7)
		} else {
			dayDiff := getDiffWeekDay(weekday, wd.GetWeekday())
			daydiffs = append(daydiffs, dayDiff)
			// append next week it self to handle edge case: only 1 weekday
			daydiffs = append(daydiffs, dayDiff+7)
		}

		for _, daydiff := range daydiffs {

			startMinofDay, err := toMinute(wd.GetStartTime())
			if err != nil {
				return 0, err
			}
			endMinOfDay, err := toMinute(wd.GetEndTime())
			if err != nil {
				return 0, err
			}
			tempDay := time.Date(year, time.Month(month), day+int(daydiff), 0, 0, 0, 0, time.UTC)
			isHoliday, err := IsHoliday(bh, tempDay, tz)
			if err != nil {
				return 0, err
			}
			if isHoliday {
				startMinofDay = 0
				endMinOfDay = 23*60 + 59
			}
			startBussinessMin := int64(startMinofDay) + daydiff*24*60 + startMinCurrentDay
			endBussinessMin := int64(endMinOfDay) + daydiff*24*60 + startMinCurrentDay
			endMin := 23*60 + 59 + daydiff*24*60 + startMinCurrentDay
			if maxEndMin < endMin {
				maxEndMin = endMin
			}
			bussesPeriodMins = append(bussesPeriodMins, []int64{startBussinessMin, endBussinessMin})
		}

	}
	var nextNonBussinessMili int64 = math.MaxInt64
	for min := currentMin; min <= maxEndMin; min++ {
		isNonBussiness := true
		for _, period := range bussesPeriodMins {
			start := period[0]
			end := period[1]

			if min >= start && min <= end {
				isNonBussiness = false
				// this period is alway false
				min = end
				break
			}
		}
		if isNonBussiness {
			nextNonBussinessMili = min * 60 * 1000
			break
		}
	}
	return nextNonBussinessMili, nil
}

// toMinute from string 10:25 convert to number of minutes: 675
func toMinute(t string) (int, error) {
	var mins, hours int
	var err error

	parts := strings.SplitN(t, ":", 2)
	switch len(parts) {
	case 1:
		mins, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
	case 2:
		hours, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}

		mins, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("invalid time: %s", t)
	}

	if mins > 59 || mins < 0 || hours > 23 || hours < 0 {
		return 0, fmt.Errorf("invalid time: %s", t)
	}

	return hours*60 + mins, nil
}

func ConvertTimezone(t time.Time, tz string) (year, mon, day, hour, min int, weekday string, err error) {
	tzhour, tzmin, err := SplitTzOffset(tz)
	if err != nil {
		return 0, 0, 0, 0, 0, "", err
	}

	t = t.UTC().Add(time.Hour*time.Duration(tzhour) +
		time.Minute*time.Duration(tzmin))
	return t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(),
		t.Weekday().String(), nil
}

// SplitTzOffset splits timezone offset (e.g: +07:00) to hour and minute
// pair (e.g: 7, 0)
// timezone offset must follow +hh:mm or -hh:mm, otherwise the function
// will return an invalid timezone offset error
// examples:
//
//	SplitTzOffset("+07:00") => 7, 0
//	SplitTzOffset("-07:30") => -7, -30
//	SplitTzOffset("-00:30") => -7, -30
func SplitTzOffset(offset string) (int, int, error) {
	offset = strings.TrimSpace(offset)
	if offset == "" || offset == "0" || offset == "00:00" || offset == "Z" {
		return 0, 0, nil
	}
	sign := 1
	if offset[0] == '-' {
		sign = -1
	} else if offset[0] == '+' {
	} else {
		offset = "+" + offset
	}
	if len(offset) != 6 {
		return 0, 0, fmt.Errorf("invalid timezone offset %s", offset)
	}

	if offset[3] != ':' {
		return 0, 0, fmt.Errorf("invalid timezone offset %s", offset)
	}

	if !isNumeric(offset[1]) || !isNumeric(offset[2]) ||
		!isNumeric(offset[4]) || !isNumeric(offset[5]) {
		return 0, 0, fmt.Errorf("invalid timezone offset %s", offset)
	}
	hour := (offset[1]-'0')*10 + offset[2] - '0'
	min := (offset[4]-'0')*10 + offset[5] - '0'

	// why not 24?, see https://en.wikipedia.org/wiki/List_of_UTC_time_offsets
	// for the list of all possible timezone
	if hour > 14 {
		return 0, 0, fmt.Errorf("invalid timezone offset %s", offset)
	}

	if min > 60 {
		return 0, 0, fmt.Errorf("invalid timezone offset %s", offset)
	}

	if sign == -1 {
		return -int(hour), -int(min), nil
	}
	return int(hour), int(min), nil
}

func isNumeric(r byte) bool { return '0' <= r && r <= '9' }

func getDiffWeekDay(weekDayLhs, weekDayRhs string) int64 {
	diffDay := weekDayM[weekDayLhs] - weekDayM[weekDayRhs]
	if diffDay < 0 {
		return -diffDay
	}
	return diffDay

}

var weekDayM map[string]int64

func init() {
	weekDayM = map[string]int64{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}
}
