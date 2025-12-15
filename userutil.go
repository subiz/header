package header

import (
	"encoding/json"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	apb "github.com/subiz/header/account"
	"github.com/tidwall/gjson"
)

const Tolerance = 0.000001
const TRUEB = byte('t')
const FALSEB = byte('f')

const STRDELIMITER = "\000"
const RUNEDELIMITER = '\000'

func applyTextTransform(str string, transforms []*TextTransform) string {
	if len(transforms) == 0 {
		return str
	}

	transform := transforms[0]
	if transform.GetName() == "nospace" {
		str = SpaceStringsBuilder(str)
	}

	if transform.GetName() == "trim" {
		str = strings.TrimSpace(str)
	}

	if transform.GetName() == "lower_case" {
		str = strings.ToLower(str)
	}

	if transform.GetName() == "upper_case" {
		str = strings.ToUpper(str)
	}

	return applyTextTransform(str, transforms[1:])
}

func applyNumberTransform(fl float64, transforms []*NumberTransform) float64 {
	return fl
}

func EvaluateText(has bool, str string, cond *TextCondition) bool {
	if len(cond.GetTransforms()) > 0 {
		str = applyTextTransform(str, cond.GetTransforms())
	}
	if !cond.GetCaseSensitive() {
		str = strings.ToLower(str)
	}
	if !cond.GetAccentSensitive() {
		str = Ascii(str)
	}

	switch cond.GetOp() {
	case "any":
		return true
	case "has_value":
		return has && str != ""
	case "is_empty", "empty":
		return strings.TrimSpace(str) == ""
	case "eq":
		if len(cond.GetEq()) == 0 {
			return true
		}
		if !has {
			return false
		}
		for _, cs := range cond.GetEq() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			if strings.TrimSpace(str) == strings.TrimSpace(cs) {
				return true
			}
		}
		return false
	case "neq":
		if len(cond.GetNeq()) == 0 {
			return true
		}

		for _, cs := range cond.GetNeq() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			if strings.TrimSpace(str) == strings.TrimSpace(cs) {
				return false
			}
		}
		return true
	case "regex":
		if !has {
			return false
		}

		regexp.MatchString(cond.GetRegex(), str)
	case "start_with":
		if !has {
			return false
		}
		for _, cs := range cond.GetStartWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			if strings.HasPrefix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
				return true
			}
		}
		return false

	case "end_with":
		if !has {
			return false
		}
		for _, cs := range cond.GetEndWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			if strings.HasSuffix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
				return true
			}
		}
		return false
	case "contain":
		if !has {
			return false
		}
		for _, cs := range cond.GetContain() {
			if strings.Contains(str, cs) {
				return true
			}
		}
		return false
	case "not_contain":
		if !has {
			return false
		}
		for _, cs := range cond.GetNotContain() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			if strings.Contains(strings.TrimSpace(str), strings.TrimSpace(cs)) {
				return false
			}
		}
		return true
	case "not_start_with":
		if !has {
			return false
		}
		for _, cs := range cond.GetNotStartWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			if strings.HasPrefix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
				return false
			}
		}
		return true
	case "not_end_with":
		if !has {
			return false
		}
		for _, cs := range cond.GetEndWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			if strings.HasSuffix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
				return false
			}
		}
		return true
	default:
		return true
	}

	return true
}

func EvaluateTexts(strs []string, cond *TextCondition) bool {
	if len(cond.GetTransforms()) > 0 {
		for i, str := range strs {
			strs[i] = applyTextTransform(str, cond.GetTransforms())
		}
	}

	if !cond.GetCaseSensitive() {
		for i, str := range strs {
			strs[i] = strings.ToLower(str)
		}
	}
	if !cond.GetAccentSensitive() {
		for i, str := range strs {
			strs[i] = Ascii(str)
		}
	}

	switch cond.GetOp() {
	case "any":
		return true
	case "has_value":
		return len(strs) != 0
	case "is_empty", "empty":
		return len(strs) == 0
	case "eq":
		if len(cond.GetEq()) == 0 {
			return true
		}
		for _, cs := range cond.GetEq() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			for _, str := range strs {
				if strings.TrimSpace(str) == strings.TrimSpace(cs) {
					return true
				}
			}
		}
		return false
	case "neq":
		if len(cond.GetNeq()) == 0 {
			return true
		}

		for _, cs := range cond.GetNeq() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			for _, str := range strs {
				if strings.TrimSpace(str) == strings.TrimSpace(cs) {
					return false
				}
			}
		}
		return true
	case "regex":
		for _, str := range strs {
			if b, _ := regexp.MatchString(cond.GetRegex(), str); b {
				return true
			}
		}
		return false
	case "start_with":
		for _, cs := range cond.GetStartWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			for _, str := range strs {
				if strings.HasPrefix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
					return true
				}
			}
		}
		return false

	case "end_with":
		for _, cs := range cond.GetEndWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			for _, str := range strs {
				if strings.HasSuffix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
					return true
				}
			}
		}
		return false
	case "contain", "con":
		for _, cs := range cond.GetContain() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			for _, str := range strs {
				if strings.Contains(strings.TrimSpace(str), strings.TrimSpace(cs)) {
					return true
				}
			}
		}
		return false
	case "not_contain":
		for _, cs := range cond.GetNotContain() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			for _, str := range strs {
				if strings.Contains(strings.TrimSpace(str), strings.TrimSpace(cs)) {
					return false
				}
			}
		}
		return true
	case "not_start_with":
		for _, cs := range cond.GetNotStartWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			for _, str := range strs {
				if strings.HasPrefix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
					return false
				}
			}
		}
		return true
	case "not_end_with":

		for _, cs := range cond.GetEndWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			for _, str := range strs {
				if strings.HasSuffix(strings.TrimSpace(str), strings.TrimSpace(cs)) {
					return false
				}
			}
		}
		return true
	default:
		return true
	}

	return true
}

func EvaluateNumber(found bool, fl float64, cond *NumberCondition) bool {
	fl = applyNumberTransform(fl, cond.GetTransforms())

	switch cond.GetOp() {
	case "has_value":
		return found
	case "is_empty", "empty":
		return !found
	case "eq":
		if len(cond.GetEq()) == 0 {
			return true
		}
		for _, cf := range cond.GetEq() {
			if math.Abs(cf-fl) < Tolerance {
				return true
			}
		}
		return false

	case "neq":
		if len(cond.GetNeq()) == 0 {
			return true
		}
		for _, cf := range cond.GetNeq() {
			if math.Abs(cf-fl) < Tolerance {
				return false
			}
		}
		return true
	case "gt":
		return fl > cond.GetGt()
	case "lt":
		return fl <= cond.GetLt()
	case "gte":
		return fl >= cond.GetGte() || math.Abs(fl-cond.GetGte()) < Tolerance
	case "lte":
		return fl <= cond.GetLte() || math.Abs(fl-cond.GetLte()) < Tolerance
	case "in_range":
		if len(cond.GetInRange()) < 2 {
			return false
		}
		return cond.GetInRange()[0] <= fl && fl <= cond.GetInRange()[1]
	case "not_in_range":
		if len(cond.GetNotInRange()) < 2 {
			return false
		}
		return fl <= cond.GetNotInRange()[0] || cond.GetNotInRange()[1] <= fl
	}

	return true
}

func EvaluateBoolean(found, boo bool, cond *BooleanCondition) bool {
	switch cond.GetOp() {
	// apply transform first
	case "has_value":
		return found
	case "true", "t":
		return boo
	case "false", "f":
		return !boo
	}
	return true
}

func EvaluateDatetime(acc *apb.Account, found bool, unixms int64, cond *DatetimeCondition) bool {
	t := time.Unix(unixms/1000, 0)

	switch cond.GetOp() {
	case "any":
		return true
	case "unset":
		return !found
	case "has_value":
		return found
	// apply transform first
	case "in_business_hour":
		inbusinesshours, _ := DuringBusinessHour(acc.GetBusinessHours(), t, acc.GetTimezone())
		return inbusinesshours
	case "non_business_hour":
		inbusinesshours, _ := DuringBusinessHour(acc.GetBusinessHours(), t, acc.GetTimezone())
		return !inbusinesshours
	case "today":
		h, m, _ := SplitTzOffset(acc.GetTimezone())
		location := time.FixedZone("account_tz", h*3600+m*60)

		nowInLoc := time.Now().In(location)
		year, month, day := nowInLoc.Date()
		startoftoday := time.Date(year, month, day, 0, 0, 0, 0, location)
		endoftoday := time.Date(year, month, day, 23, 59, 59, 0, location)

		return startoftoday.Unix() <= t.Unix() && t.Unix() <= endoftoday.Unix()
	case "date_last_30mins":
		now := time.Now().Unix()
		last30mins := now - 1800
		return last30mins <= t.Unix() && t.Unix() <= now
	case "date_last_2hours":
		now := time.Now().Unix()
		last2hours := now - 7200
		return last2hours <= t.Unix() && t.Unix() <= now
	case "date_last_24h":
		now := time.Now().Unix()
		last1days := now - 86400
		return last1days <= t.Unix() && t.Unix() <= now
	case "date_last_7days":
		now := time.Now().Unix()
		last7days := now - 7*86400
		return last7days <= t.Unix() && t.Unix() <= now
	case "date_last_30days":
		now := time.Now().Unix()
		last30days := now - 30*86400
		return last30days <= t.Unix() && t.Unix() <= now
	case "yesterday":
		h, m, _ := SplitTzOffset(acc.GetTimezone())
		location := time.FixedZone("account_tz", h*3600+m*60)
		nowInLoc := time.Now().In(location)

		// Yesterday is the day before today
		yesterday := nowInLoc.AddDate(0, 0, -1)
		year, month, day := yesterday.Date()
		startofyesterday := time.Date(year, month, day, 0, 0, 0, 0, location)
		endofyesterday := time.Date(year, month, day, 23, 59, 59, 0, location)

		return startofyesterday.Unix() <= t.Unix() && t.Unix() <= endofyesterday.Unix()
	case "last_week":
		h, m, _ := SplitTzOffset(acc.GetTimezone())
		location := time.FixedZone("account_tz", h*3600+m*60)
		nowInLoc := time.Now().In(location)

		// Monday of the current week
		weekday := nowInLoc.Weekday()
		daysToSubtract := int(weekday - time.Monday)
		if daysToSubtract < 0 {
			daysToSubtract += 7
		}
		startofweek_date := nowInLoc.AddDate(0, 0, -daysToSubtract)
		year, month, day := startofweek_date.Date()
		startofthisweek := time.Date(year, month, day, 0, 0, 0, 0, location)

		// Last week is the 7 days before the start of this week
		startoflastweek := startofthisweek.AddDate(0, 0, -7)
		endoflastweek_date := startofthisweek.AddDate(0, 0, -1)
		year, month, day = endoflastweek_date.Date()
		endoflastweek := time.Date(year, month, day, 23, 59, 59, 0, location)

		return startoflastweek.Unix() <= t.Unix() && t.Unix() <= endoflastweek.Unix()
	case "this_week":
		h, m, _ := SplitTzOffset(acc.GetTimezone())
		location := time.FixedZone("account_tz", h*3600+m*60)
		nowInLoc := time.Now().In(location)

		// Monday of the current week
		weekday := nowInLoc.Weekday()
		daysToSubtract := int(weekday - time.Monday)
		if daysToSubtract < 0 {
			daysToSubtract += 7
		}
		startofweek_date := nowInLoc.AddDate(0, 0, -daysToSubtract)
		year, month, day := startofweek_date.Date()
		startofweek := time.Date(year, month, day, 0, 0, 0, 0, location)

		// Sunday of the current week
		endofweek_date := startofweek.AddDate(0, 0, 6)
		year, month, day = endofweek_date.Date()
		endofweek := time.Date(year, month, day, 23, 59, 59, 0, location)

		return startofweek.Unix() <= t.Unix() && t.Unix() <= endofweek.Unix()
	case "last_month":
		h, m, _ := SplitTzOffset(acc.GetTimezone())
		location := time.FixedZone("account_tz", h*3600+m*60)
		nowInLoc := time.Now().In(location)

		year, month, _ := nowInLoc.Date()
		startofthismonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
		startoflastmonth := startofthismonth.AddDate(0, -1, 0)

		endoflastmonth_date := startofthismonth.AddDate(0, 0, -1)
		year, month, day := endoflastmonth_date.Date()
		endoflastmonth := time.Date(year, month, day, 23, 59, 59, 0, location)

		return startoflastmonth.Unix() <= t.Unix() && t.Unix() <= endoflastmonth.Unix()
	case "this_month":
		h, m, _ := SplitTzOffset(acc.GetTimezone())
		location := time.FixedZone("account_tz", h*3600+m*60)
		nowInLoc := time.Now().In(location)

		year, month, _ := nowInLoc.Date()
		startofmonth := time.Date(year, month, 1, 0, 0, 0, 0, location)

		endofmonth_date := startofmonth.AddDate(0, 1, -1)
		year, month, day := endofmonth_date.Date()
		endofmonth := time.Date(year, month, day, 23, 59, 59, 0, location)

		return startofmonth.Unix() <= t.Unix() && t.Unix() <= endofmonth.Unix()
	case "last":
		a := time.Now().Unix() - cond.GetLast()
		b := time.Now().Unix()
		return a <= t.Unix() && t.Unix() <= b
	case "before_ago":
		return t.Unix() < time.Now().Unix()-cond.GetBeforeAgo()
	case "days_of_week":
		for _, weekday := range cond.GetDaysOfWeek() {
			if strings.EqualFold(weekday, t.Weekday().String()) {
				return true
			}
		}
		return false
	case "after":
		return time.Unix(cond.GetAfter()/1000, 0).Unix() <= t.Unix()
	case "before":
		return t.Unix() <= time.Unix(cond.GetBefore()/1000, 0).Unix()
	case "between":
		if len(cond.GetBetween()) != 2 {
			return true
		}
		a := cond.GetBetween()[0] / 1000
		b := cond.GetBetween()[1] / 1000
		return a <= t.Unix() && t.Unix() <= b
	case "outside":
		if len(cond.GetOutside()) != 2 {
			return true
		}
		a := cond.GetOutside()[0] / 1000
		b := cond.GetOutside()[1] / 1000
		return t.Unix() <= a || b <= t.Unix()
	}
	return true
}

func HasDeletedCond(cond *UserViewCondition) bool {
	if len(cond.GetOne()) > 0 {
		for _, c := range cond.GetOne() {
			if HasDeletedCond(c) {
				return true
			}
		}
		return false
	}

	if len(cond.GetAll()) > 0 {
		for _, c := range cond.GetAll() {
			if HasDeletedCond(c) {
				return true
			}
		}
		return false
	}
	return cond.GetKey() == "deleted" && cond.GetBoolean().GetOp() == "true"
}

// must pass in primary user and its secondaries
func RsCheck(acc *apb.Account, users []*User, cond *UserViewCondition, deleted bool) bool {
	if len(users) == 0 {
		return true
	}

	if len(cond.GetOne()) > 0 {
		for _, c := range cond.GetOne() {
			if RsCheck(acc, users, c, deleted) {
				return true
			}
		}
		return false
	}

	if len(cond.GetAll()) > 0 {
		for _, c := range cond.GetAll() {
			if !RsCheck(acc, users, c, deleted) {
				return false
			}
		}
		return true
	}

	primary := users[0]
	for _, u := range users {
		if u.PrimaryId == "" || u.PrimaryId == u.Id {
			primary = u
			break
		}
	}

	if IsPrimaryCond(cond) {
		return evaluateSingleCond(acc, primary, cond, deleted)
	}

	if IsPositiveCond(cond) {
		// positive condition like equal, just one match -> group match
		for _, u := range users {
			if pass := evaluateSingleCond(acc, u, cond, deleted); pass {
				return true
			}
		}
		return false
	}

	// negative condition like not not_eq -> all element must pass
	for _, u := range users {
		if pass := evaluateSingleCond(acc, u, cond, deleted); !pass {
			return false
		}
	}
	return true
}

func IsPrimaryCond(cond *UserViewCondition) bool {
	key := cond.GetKey()
	if strings.HasPrefix(key, "attr:") || strings.HasPrefix(key, "attr.") {
		key = key[5:]
	}

	return key == "lifecycle_stage" || key == "owner"
}

func IsPositiveCond(cond *UserViewCondition) bool {
	switch cond.GetType() {
	case "number":
		switch cond.GetNumber().GetOp() {
		case "is_empty", "not_end_with", "not_start_with", "not_contain", "neq", "not_in_range", "outside":
			return false
		}
		return true

	case "datetime":
		switch cond.GetDatetime().GetOp() {
		case "is_empty", "not_end_with", "not_start_with", "not_contain", "neq", "not_in_range", "false", "non_business_hour", "outside":
			return false
		}
		return true

	case "bool", "boolean":
		switch cond.GetBoolean().GetOp() {
		case "is_empty", "false":
			return false
		}
		return true
	}

	// fallback to text
	switch cond.GetText().GetOp() {
	case "is_empty", "not_end_with", "not_start_with", "not_contain", "neq", "not_in_range", "outside":
		return false
	}

	return true
}

func evaluateSingleCond(acc *apb.Account, u *User, cond *UserViewCondition, deleted bool) bool {
	if deleted && u.Deleted == 0 {
		return false
	}

	if !deleted && u.Deleted > 0 {
		return false
	}

	if cond.GetKey() == "is_primary" {
		isprimary := u.GetPrimaryId() == "" || u.GetPrimaryId() == u.GetId()
		return EvaluateBoolean(true, isprimary, cond.GetBoolean())
	}

	if cond.GetKey() == "id" {
		id := u.GetId()
		return EvaluateText(true, id, cond.GetText())
	}

	if cond.GetKey() == "deleted" {
		return EvaluateBoolean(true, u.Deleted > 0, cond.GetBoolean())
	}

	if cond.GetKey() == "channel" {
		return EvaluateText(true, u.Channel, cond.GetText())
	}

	if cond.GetKey() == "channel_source" {
		text := cond.GetText()
		if cond.GetText().GetOp() == "contain" {
			newContains := []string{}
			for _, c := range cond.GetText().GetContain() {
				containsplit := strings.Split(c, ".")
				if len(containsplit) > 2 {
					newContains = append(newContains, containsplit[1])
				} else {
					newContains = append(newContains, c)
				}
			}
			text = &TextCondition{Op: "contain", Contain: newContains}
		}
		return EvaluateText(true, u.ChannelSource, text)
	}

	if cond.GetKey() == "keyword" && len(cond.GetText().GetContain()) > 0 { // email phone or name
		// remove space
		keyword := Ascii(SpaceStringsBuilder(strings.ToLower(cond.GetText().GetContain()[0])))

		for _, attr := range u.Attributes {
			if attr.Text != "" {
				if strings.Contains(Ascii(strings.ToLower(SpaceStringsBuilder(attr.Text))), keyword) {
					return true
				}
			}
		}

		return strings.Contains(strings.TrimSpace(strings.ToLower(u.Id)), keyword)
	}

	if cond.GetKey() == "lead_owners" {
		for _, owner := range u.GetLeadOwners() {
			if EvaluateText(true, owner, cond.GetText()) {
				return true
			}
		}

		if len(u.GetLeadOwners()) == 0 {
			if EvaluateText(false, "", cond.Text) {
				return true
			}
		}
		return false
	}

	if cond.GetKey() == "lead_conversion_bys" {
		for _, by := range u.GetLeadConversionBys() {
			if EvaluateText(true, by, cond.GetText()) {
				return true
			}
		}

		if len(u.GetLeadConversionBys()) == 0 {
			if EvaluateText(false, "", cond.Text) {
				return true
			}
		}
		return false
	}

	if cond.GetKey() == "start_content_view:by:device:ip" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetIp(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:language" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetLanguage(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:page_title" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetPageTitle(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:page_url" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetPageUrl(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:platform" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetPlatform(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:referrer" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetReferrer(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:screen_resolution" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetScreenResolution(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:source" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetSource(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:type" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetType(), cond.Text)
	}
	if cond.GetKey() == "start_content_view:by:device:user_agent" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetUserAgent(), cond.Text)
	}

	if cond.GetKey() == "start_content_view:by:device:utm:name" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetUtm().GetName(), cond.Text)
	}

	if cond.GetKey() == "start_content_view:by:device:utm:source" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetUtm().GetSource(), cond.Text)
	}

	if cond.GetKey() == "start_content_view:by:device:utm:medium" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetUtm().GetMedium(), cond.Text)
	}

	if cond.GetKey() == "start_content_view:by:device:utm:term" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetUtm().GetTerm(), cond.Text)
	}

	if cond.GetKey() == "start_content_view:by:device:utm:content" {
		return EvaluateText(u.StartContentView != nil, u.GetStartContentView().GetBy().GetDevice().GetUtm().GetContent(), cond.Text)
	}

	if cond.GetKey() == "first_content_view:by:device:ip" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetIp(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:language" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetLanguage(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:page_title" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetPageTitle(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:page_url" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetPageUrl(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:platform" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetPlatform(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:referrer" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetReferrer(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:screen_resolution" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetScreenResolution(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:source" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetSource(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:type" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetType(), cond.Text)
	}
	if cond.GetKey() == "first_content_view:by:device:user_agent" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetUserAgent(), cond.Text)
	}

	if cond.GetKey() == "first_content_view:by:device:utm:name" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetUtm().GetName(), cond.Text)
	}

	if cond.GetKey() == "first_content_view:by:device:utm:source" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetUtm().GetSource(), cond.Text)
	}

	if cond.GetKey() == "first_content_view:by:device:utm:medium" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetUtm().GetMedium(), cond.Text)
	}

	if cond.GetKey() == "first_content_view:by:device:utm:term" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetUtm().GetTerm(), cond.Text)
	}

	if cond.GetKey() == "first_content_view:by:device:utm:content" {
		return EvaluateText(u.FirstContentView != nil, u.GetFirstContentView().GetBy().GetDevice().GetUtm().GetContent(), cond.Text)
	}

	if cond.GetKey() == "labels" {
		labels := []string{}
		for _, label := range u.Labels {
			labels = append(labels, label.Label)
		}

		return EvaluateTexts(labels, cond.Text)
	}

	if strings.HasPrefix(cond.GetKey(), "start_content_view.") ||
		strings.HasPrefix(cond.GetKey(), "first_content_view.") ||
		strings.HasPrefix(cond.GetKey(), "latest_content_view.") {
		prefix := strings.Split(cond.GetKey(), "_")[0]
		path := strings.Split(cond.GetKey(), ".")[1]
		ub, _ := json.Marshal(u)
		haspageview := u.FirstContentView != nil || u.StartContentView != nil || u.LatestContentView != nil
		if path == "url" {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.referrer").String(), cond.Text)
		} else if path == "source" {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.source").String(), cond.Text)
		} else if path == "domain" {
			// _content_view.by.device.referrer.host ->
			// parsedUrl, _ := url.Parse(contentView.GetBy().GetDevice().GetReferrer())
			// domain := parsedUrl.Host
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.referrer.host").String(), cond.Text)
		} else if path == "referrer_url" {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.source_referrer").String(), cond.Text)
		} else if path == "referrer_domain" {
			// _content_view.by.device.source_referrer.host ->
			// parsedUrl, _ := url.Parse(contentView.GetBy().GetDevice().GetSourceReferrer())
			// domain := parsedUrl.Host
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.source_referrer.host").String(), cond.Text)
		} else if strings.HasSuffix(cond.GetKey(), "_content_view.utm.name") {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.utm.name").String(), cond.Text)
		} else if strings.HasSuffix(cond.GetKey(), "_content_view.utm.source") {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.utm.source").String(), cond.Text)
		} else if strings.HasSuffix(cond.GetKey(), "_content_view.utm.medium") {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.utm.medium").String(), cond.Text)
		} else if strings.HasSuffix(cond.GetKey(), "_content_view.utm.term") {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.utm.term").String(), cond.Text)
		} else if strings.HasSuffix(cond.GetKey(), "_content_view.utm.id") {
			return EvaluateText(haspageview, gjson.GetBytes(ub, prefix+"_content_view.by.device.utm.id").String(), cond.Text)
		}
		return false
	}

	if strings.HasPrefix(cond.GetKey(), "attr:") || strings.HasPrefix(cond.GetKey(), "attr.") {
		key := cond.GetKey()[5:]

		defType := cond.GetType()
		if defType == "list" || defType == "" {
			defType = "text"
		}
		text, num, date, boo, found := FindAttr(u, key, defType)
		if defType == "text" {
			return EvaluateText(found, text, cond.GetText())
		}

		if defType == "number" {
			return EvaluateNumber(found, num, cond.GetNumber())
		}
		if defType == "boolean" {
			return EvaluateBoolean(found, boo, cond.GetBoolean())
		}
		if defType == "datetime" { // consider number in ms
			return EvaluateDatetime(acc, found, date, cond.Datetime)
		}
	}
	return true
}

func FindAttr(u *User, key string, typ string) (string, float64, int64, bool, bool) {
	for _, a := range u.Attributes {
		if a.Key != key {
			continue
		}
		t, err := time.Parse(time.RFC3339, a.GetDatetime())
		if err != nil {
			t = time.Unix(0, 0)
		}

		text := a.Text
		return text, a.Number, t.UnixMilli(), a.Boolean, true
	}
	return "", 0, 0, false, false
}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func GetTimeAttr(u *User, key string) (time.Time, bool) {
	key = strings.ToLower(strings.TrimSpace(key))
	has := false
	for _, a := range u.GetAttributes() {
		if key != a.GetKey() {
			continue
		}

		has = true
		t, err := time.Parse(time.RFC3339, a.GetDatetime())
		if err == nil {
			return t, has
		}
	}

	return time.Unix(0, 0), has
}

func GetAttr(u *User, key string, typ string) interface{} {
	key = strings.ToLower(strings.TrimSpace(key))
	for _, a := range u.GetAttributes() {
		if key != a.GetKey() {
			continue
		}
		switch typ {
		case AttributeDefinition_text.String():
			return a.GetText()
		case AttributeDefinition_number.String():
			return a.GetNumber()
		case AttributeDefinition_boolean.String():
			return a.GetBoolean()
		case "list":
			return a.GetOtherValues()
		case AttributeDefinition_datetime.String():
			t, err := time.Parse(time.RFC3339, a.GetDatetime())
			if err != nil {
				return time.Now()
			}
			return t
		}
		return nil
	}
	return nil
}

func SetAttr(u *User, key string, typ string, val interface{}) {
	key = strings.ToLower(strings.TrimSpace(key))
	if u == nil || val == nil || key == "" || typ == "" {
		return
	}
	a := &Attribute{}
	a.Key = key
	switch typ {
	case AttributeDefinition_text.String():
		v, _ := val.(string)
		a.Text = v
	case AttributeDefinition_number.String():
		vb, _ := json.Marshal(val)
		v, _ := strconv.ParseFloat(string(vb), 64)
		a.Number = v
	case AttributeDefinition_boolean.String():
		v, _ := val.(bool)
		a.Boolean = v
	case "list":
		ss, _ := val.([]string)
		if len(ss) > 0 {
			a.OtherValues = ss[1:]
			a.Text = ss[0]
		}
	case AttributeDefinition_datetime.String():
		t, _ := val.(time.Time)
		a.Datetime = t.Format(time.RFC3339)
	}
	for _, i := range u.GetAttributes() {
		if i.GetKey() != key {
			continue
		}

		i.Text, i.Number, i.Boolean, i.Datetime =
			a.Text, a.Number, a.Boolean, a.Datetime
		if typ == "list" {
			i.OtherValues = a.OtherValues
		}
		return
	}

	u.Attributes = append(u.Attributes, a)
}

func GetTextAttr(u *User, key string) string {
	key = strings.ToLower(strings.TrimSpace(key))
	for _, a := range u.GetAttributes() {
		if key != a.GetKey() {
			continue
		}
		return a.GetText()
	}
	return ""
}

func NormalizeCond(cond *UserViewCondition) {
	if len(cond.GetOne()) > 0 {
		for _, c := range cond.GetOne() {
			NormalizeCond(c)
		}
	}

	if len(cond.GetAll()) > 0 {
		for _, c := range cond.GetAll() {
			NormalizeCond(c)
		}
	}

	NormalizeTextCond(cond.Text)
}

func NormalizeTextCond(cond *TextCondition) {
	if cond == nil {
		return
	}
	switch cond.GetOp() {
	case "eq":
		for i, cs := range cond.GetEq() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			cond.Eq[i] = strings.TrimSpace(cs)
		}
	case "neq":
		for i, cs := range cond.GetNeq() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			cond.Neq[i] = strings.TrimSpace(cs)
		}
	case "start_with":
		for i, cs := range cond.GetStartWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			cond.StartWith[i] = strings.TrimSpace(cs)
		}
	case "end_with":
		for i, cs := range cond.GetEndWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			cond.EndWith[i] = strings.TrimSpace(cs)
		}
	case "contain":
		for i, cs := range cond.GetContain() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			cond.Contain[i] = strings.TrimSpace(cs)
		}
	case "not_contain":
		for i, cs := range cond.GetNotContain() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			cond.NotContain[i] = strings.TrimSpace(cs)
		}
	case "not_start_with":
		for i, cs := range cond.GetNotStartWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}
			cond.NotStartWith[i] = strings.TrimSpace(cs)
		}
	case "not_end_with":
		for i, cs := range cond.GetEndWith() {
			if !cond.GetCaseSensitive() {
				cs = strings.ToLower(cs)
			}
			if !cond.GetAccentSensitive() {
				cs = Ascii(cs)
			}

			cond.NotEndWith[i] = strings.TrimSpace(cs)
		}
	}
}
