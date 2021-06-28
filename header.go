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

var Scopes = makeScopeMap()

func makeScopeMap() map[string]string {
	// scope => permission
	var m = map[string]string{}
	m["agent"] = "account:r conversation:rw agent_group:r rule:r integration:r message_template:rw other_message_template:r tag:r whitelist:wr widget:r subscription:r invoice:r user:rw attribute:r facebook:r bot:r agent:r conversation_setting:r web_plugin:r file:wr webpage:r lang:r user_view:rw user_label:r"
	m["view_other_convos"] = "other_conversation:r"
	m["export_user"] = "user:e" // export

	m["account_setting"] = m["agent"] + " account:w agent:w agent_group:w rule:w integration:w other_message_template:rw tag:w widget:w attribute:w facebook:w bot:w conversation_setting:w web_plugin:wr webhook:wr webpage:w lang:w user_label:w"
	m["account_manage"] = m["account_setting"] + "account:w agent_group:wr agent:w subscription:rw payment_method:rw"
	m["owner"] = m["account_manage"] + " " + m["account_setting"]
	m["subiz"] = m["account_manage"] + " " + m["account_setting"] + " payment:w"
	m["all"] = m["subiz"]
	return m
}

func prettyPerm(perm string) string {
	perms := strings.FieldsFunc(perm, func(r rune) bool {
		return r == ' ' || r == ';' || r == ',' || r == '\n' || r == '\t'
	})
	permM := make(map[string]string)
	for _, p := range perms {
		p = strings.TrimSpace(p)

		psplit := strings.Split(p, ":")
		if len(psplit) != 2 {
			continue
		}
		permM[psplit[0]] += psplit[1]
	}

	out := ""
	for k, v := range permM {
		ppp := ""
		if strings.Contains(v, "w") {
			ppp += "w"
		}
		if strings.Contains(v, "r") {
			ppp += "r"
		}
		if strings.Contains(v, "e") {
			ppp += "e"
		}
		if strings.Contains(v, "p") {
			ppp += "p"
		}
		if ppp != "" {
			out += strings.TrimSpace(k) + ":" + ppp + " "
		}
	}
	return strings.TrimSpace(out)
}

// []string{"all", "agent"}, "conversation:r tag:wr" => true
// []string{"agent"}, "tag:wr" => false
func checkAccess(scopes []string, perm string) bool {
	// make availabe perm map by joining all permision in scopes
	availableperm := make(map[string]string) // {"conversation" => "cr", "user" => "u"}
	joinperm := ""
	for _, scope := range scopes {
		joinperm += " " + Scopes[strings.TrimSpace(scope)]
	}

	joinperm = prettyPerm(joinperm)
	joinpermsplit := strings.Split(joinperm, " ")
	for _, joinpermitem := range joinpermsplit {
		joinpermitemsplit := strings.Split(joinpermitem, ":")
		if len(joinpermitemsplit) != 2 {
			continue
		}
		availableperm[joinpermitemsplit[0]] = joinpermitemsplit[1]
	}

	perm = prettyPerm(perm)
	perms := strings.Split(perm, " ")
	for _, p := range perms {
		ps := strings.Split(p, ":") // conversation:rw
		if len(ps) != 2 {
			continue
		}

		for _, p := range ps[1] {
			if !strings.Contains(availableperm[ps[0]], string(p)) {
				return false
			}
		}
	}
	return true
}

// []string{"all", "agent"}, "conversation:r tag:wr" => true
// []string{"agent"}, "tag:wr" => false
func CheckAccess(realScopes, authorizedScopes []string, perm string) bool {
	return checkAccess(realScopes, perm) && checkAccess(authorizedScopes, perm)
}
