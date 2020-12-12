package header

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func GetAttr(u *User, key string, typ string) interface{} {
	key = strings.ToLower(strings.TrimSpace(key))
	for _, a := range u.GetAttributes() {
		if !SameKey(key, a.GetKey()) {
			continue
		}
		switch typ {
		case AttributeDefinition_text.String():
			return a.GetText()
		case AttributeDefinition_number.String():
			return a.GetNumber()
		case AttributeDefinition_boolean.String():
			return a.GetBoolean()
		case AttributeDefinition_list.String():
			return a.GetList()
		case AttributeDefinition_datetime.String():
			t, err := time.Parse(time.RFC3339Nano, a.GetDatetime())
			if err != nil {
				return time.Now()
			}
			return t
		}
		return nil
	}
	return nil
}

func SameKey(k1, k2 string) bool {
	return strings.TrimSpace(strings.ToLower(k1)) ==
		strings.TrimSpace(strings.ToLower(k2))
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
		v, _ := strconv.ParseFloat(string(vb), 10)
		a.Number = v
	case AttributeDefinition_boolean.String():
		v, _ := val.(bool)
		a.Boolean = v
	case AttributeDefinition_list.String():
		ss, _ := val.([]string)
		a.List = ss
	case AttributeDefinition_datetime.String():
		t, _ := val.(time.Time)
		a.Datetime = t.Format(time.RFC3339)
	}
	for _, i := range u.GetAttributes() {
		if !SameKey(i.GetKey(), key) {
			continue
		}

		i.Text, i.Number, i.Boolean, i.Datetime, i.List =
			a.Text, a.Number, a.Boolean, a.Datetime, a.List
		return
	}

	u.Attributes = append(u.Attributes, a)
}

func SetTextAttr(u *User, key string, val string) {
	key = strings.ToLower(strings.TrimSpace(key))
	if u == nil || key == "" {
		return
	}

	a := &Attribute{}
	a.Key = key
	a.Text = val
	for _, i := range u.GetAttributes() {
		if !SameKey(i.GetKey(), key) {
			continue
		}

		i.Text = a.Text
		return
	}

	u.Attributes = append(u.Attributes, a)
}

func GetTextAttr(u *User, key string) string {
	key = strings.ToLower(strings.TrimSpace(key))
	for _, a := range u.GetAttributes() {
		if !SameKey(key, a.GetKey()) {
			continue
		}
		return a.GetText()
	}
	return ""
}
