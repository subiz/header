package header

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func DeltaToPlainText(delta string) string {
	del := make(map[string]interface{})
	if err := json.Unmarshal([]byte(delta), &del); err != nil {
		return ""
	}

	if del["ops"] == nil {
		return ""
	}

	deltas, ok := del["ops"].([]interface{})
	if !ok {
		return ""
	}

	output := ""
	for _, itemi := range deltas {
		item, ok := itemi.(map[string]interface{})
		if !ok {
			continue
		}

		if item["insert"] == nil {
			continue
		}

		// is text
		if str, ok := item["insert"].(string); ok {
			output += str
			continue
		}

		obj, ok := item["insert"].(map[string]interface{})
		if !ok {
			continue
		}

		// is emoji
		if obj["emoji"] != nil {
			code, _ := obj["emoji"].(string)
			output += ":" + code + ":"
			continue
		}

		// is mention
		if obj["mention"] != nil {
			mention, ok := obj["mention"].(map[string]interface{})
			if !ok {
				continue
			}
			fullname, _ := mention["fullname"].(string)
			output += "@" + fullname
			continue
		}
	}
	return output
}

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
		case "list":
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
		v, _ := strconv.ParseFloat(string(vb), 64)
		a.Number = v
	case AttributeDefinition_boolean.String():
		v, _ := val.(bool)
		a.Boolean = v
	case "list":
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

func UpdateAttr(u *User, key string, typ string, val interface{}, action string) {
	key = strings.ToLower(strings.TrimSpace(key))
	if u == nil || val == nil || key == "" {
		return
	}
	a := &Attribute{}
	a.Key = key
	foundIndex := -1
	for index, uAttr := range u.Attributes {
		if uAttr == nil {
			continue
		}
		if !SameKey(uAttr.Key, key) {
			continue
		}
		foundIndex = index
		break
	}

	attrs := u.Attributes
	if foundIndex >= 0 {
		attrs = append(attrs[:foundIndex], attrs[foundIndex+1:]...)
	}

	switch typ {
	case AttributeDefinition_text.String(), "string", "str", "":
		v, _ := val.(string)
		a.Text = v
	case AttributeDefinition_number.String(), "int", "float":
		vb, _ := json.Marshal(val)
		v, _ := strconv.ParseFloat(string(vb), 64)
		a.Number = v
	case AttributeDefinition_boolean.String(), "bool":
		v, _ := val.(bool)
		a.Boolean = v
	case "list":
		ss, _ := val.([]string)
		a.List = ss
	case AttributeDefinition_datetime.String(), "date", "time":
		var d time.Time
		switch t := val.(type) {
		case uint:
			d = time.Unix(int64(t)/1000, 0)
		case int:
			d = time.Unix(int64(t)/1000, 0)
		case int64:
			d = time.Unix(int64(t)/1000, 0)
		case int32:
			d = time.Unix(int64(t)/1000, 0)
		case uint64:
			d = time.Unix(int64(t)/1000, 0)
		case uint32:
			d = time.Unix(int64(t)/1000, 0)
		case time.Time:
			d = t
		}
		a.Datetime = d.Format(time.RFC3339)
	default:
		return // undefined type
	}
	a.Action = action
	attrs = append(attrs, a)
	u.Attributes = attrs
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
	m["agent"] = "account:r conversation:rw agent_group:r rule:r integration:r message_template:rw other_message_template:r tag:r whitelist:wr widget:r subscription:r invoice:r user:rw attribute:r facebook:r bot:r agent:r conversation_setting:r web_plugin:r file:wr webpage:r lang:r user_label:r user_view:rw order:rw shop_setting:r conversation_modal:r conversation_automation:r phone_device:r call_setting:r greeting_audio:r"
	m["view_other_convos"] = "other_conversation:r"
	m["view_others"] = "other_conversation:r other_lead:r other_order:r"
	m["export_user"] = "user:e" // export

	m["account_setting"] = m["agent"] + " account:w agent:w agent_group:w rule:w integration:w other_message_template:rw tag:w widget:w attribute:w facebook:w bot:w conversation_setting:w web_plugin:wr webhook:wr webpage:w lang:w user_label:w shop_setting:w conversation_modal:w conversation_automation:w phone_device:rw call_setting:rw greeting_audio:rw"
	m["account_manage"] = m["account_setting"] + "account:w agent_group:wr agent:w subscription:rw payment_method:rw"
	m["owner"] = m["account_manage"] + " " + m["account_setting"]
	m["subiz"] = m["account_manage"] + " " + m["account_setting"] + " payment:w"
	m["crm"] = m["account:r webhook:wr user:rw"]
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

func LimitLength(str string, limit int) string {
	if len(str) < limit {
		return str
	}
	return str[:limit]
}

// []string{"all", "agent"}, "conversation:r tag:wr" => true
// []string{"agent"}, "tag:wr" => false
func CheckAccess(realScopes, authorizedScopes []string, perm string) bool {
	return checkAccess(realScopes, perm) && checkAccess(authorizedScopes, perm)
}

func CalcTotalOrder(order *Order) {
	subtotal := float32(0)
	// compute subtotal first, since it doesn't depend on tax or gloal discount
	for _, item := range order.Items {
		if item.Product == nil {
			continue
		}
		price := item.Product.Price * float32(item.Quantity)

		// some products that are discounted before tax
		// the discount amount is calculated in product price for
		// simpler calucation in future
		if item.DiscountType == "percentage" {
			if item.DiscountPercentage > 0 {
				price = price * (1 - float32(item.DiscountPercentage)/10000)
			}
		} else if item.DiscountType == "amount" {
			if item.DiscountAmount > 0 {
				price = price - item.DiscountAmount
				if price < 0 {
					price = 0
				}
			}
		}

		item.Total = price
		item.FpvTotal = int64(price * order.CurrencyRate * 1000000)

		subtotal += price
	}

	// TAX
	taxM := map[string]*taxitem{}
	for i, item := range order.Items {
		if item.Product == nil {
			continue
		}

		tax := item.Product.Tax
		if tax.GetId() == "" {
			continue
		}

		taxprice := item.Total * float32(tax.Percentage) / 10000
		taxM[strconv.Itoa(i)] = &taxitem{tax: tax, taxprice: taxprice}
	}

	computedDiscount := float32(0)

	totaltax := float32(0)
	for _, t := range taxM {
		totaltax += t.taxprice
	}

	// have discount after tax
	if order.DiscountType == "percentage" && order.DiscountPercentage > 0 {
		computedDiscount = (subtotal + totaltax) * float32(order.DiscountPercentage) / 10000
	} else if order.DiscountType == "amount" {
		computedDiscount = order.DiscountAmount
	}
	if computedDiscount > (subtotal + totaltax) {
		computedDiscount = subtotal + totaltax
	}

	// SHIP
	shippingfee := order.GetShipping().GetNominalFee()
	if shippingfee > 0 {
		if order.Shipping.Tax.GetId() != "" {
			tax := order.Shipping.Tax
			taxprice := (shippingfee * (float32(tax.Percentage) / 10000))
			taxM["ship"] = &taxitem{tax: tax, taxprice: taxprice}

			order.Shipping.TotalTax = taxprice
			order.Shipping.FpvTotalTax = int64(taxprice * order.CurrencyRate * 1000000)
		}
		order.Shipping.FpvNominalFee = int64(order.Shipping.NominalFee * order.CurrencyRate * 1000000)
	}

	totaltax = float32(0)
	for _, t := range taxM {
		totaltax += t.taxprice
	}

	total := subtotal + totaltax - computedDiscount + shippingfee
	total += order.Adjustment
	if total < 0 {
		total = 0
	}

	// calculate total and fpv(s) ...
	order.Subtotal = subtotal
	order.FpvSubtotal = int64(subtotal * order.CurrencyRate * 1000000)

	order.Total = total
	order.FpvTotal = int64(total * order.CurrencyRate * 1000000)

	order.FpvDiscountAmount = int64(order.DiscountAmount * order.CurrencyRate * 1000000)

	order.FpvAdjustment = int64(order.Adjustment * order.CurrencyRate * 1000000)

	order.TotalTax = totaltax
	order.FpvTotalTax = int64(totaltax * order.CurrencyRate * 1000000)
	for i, item := range order.Items {
		if item.Product == nil {
			continue
		}

		t := taxM[strconv.Itoa(i)]
		if t != nil {
			item.TotalTax = t.taxprice
			item.FpvTotalTax = int64(t.taxprice * order.CurrencyRate * 1000000)
		}
	}
}

type taxitem struct {
	tax      *Tax
	taxprice float32
}

// like js Object.assign(dst, src)
// dst and src must be same struct pointer
// AssignObject(ag1, ag2, ["fullname", "email"])
func AssignObject(dst, src interface{}, fields []string) {
	if dst == nil || src == nil {
		return
	}
	var srcValueOf reflect.Value
	var srcTypeOf reflect.Type
	if reflect.TypeOf(src).Kind() == reflect.Ptr {
		srcValueOf, srcTypeOf = reflect.ValueOf(src).Elem(), reflect.TypeOf(src).Elem()
	} else {
		srcValueOf, srcTypeOf = reflect.ValueOf(src), reflect.TypeOf(src)
	}

	var dstValueOf reflect.Value
	var dstTypeOf reflect.Type
	if reflect.TypeOf(dst).Kind() == reflect.Ptr {
		dstValueOf, dstTypeOf = reflect.ValueOf(dst).Elem(), reflect.TypeOf(dst).Elem()
	} else {
		dstValueOf, dstTypeOf = reflect.ValueOf(dst), reflect.TypeOf(dst)
	}

	if reflect.ValueOf(dst).IsNil() || reflect.ValueOf(src).IsNil() {
		return
	}
	if dstTypeOf != srcTypeOf {
		return
	}

	for i := 0; i < srcValueOf.NumField(); i++ {
		tf := srcTypeOf.Field(i)
		jsonname := strings.Split(tf.Tag.Get("json"), ",")[0]
		if jsonname == "-" {
			continue
		}

		// skip if not field that user interested in
		canskip := false
		if len(fields) > 0 {
			canskip = true
			for _, field := range fields {
				if field == jsonname {
					canskip = false
					break
				}
			}
		}
		if canskip {
			continue
		}

		if !dstValueOf.Field(i).CanSet() {
			continue
		}
		dstValueOf.Field(i).Set(srcValueOf.Field(i))
	}
}

// NormPhone converts an user input phone number to
// standalized phone number
// (84)35 9423 423 => 0359423423
//
// A good phone number only contain numbers
//   if the number starts with 84 => replace to 0 since we mostly serve
//   Vietnamese customers
func NormPhone(phone string) string {
	phonesplit := strings.FieldsFunc(phone, func(r rune) bool {
		return r == ',' || r == ';'
	})

	phones := []string{}
	for _, phone := range phonesplit {
		arr := make([]rune, 0)
		for _, r := range phone {
			if r >= '0' && r <= '9' {
				arr = append(arr, r)
			}
		}
		number := string(arr)
		if number != "" {
			phones = append(phones, number)
		}
	}
	// remove 84
	for i, phone := range phones {
		if strings.HasPrefix(phone, "84") {
			phones[i] = "0" + phone[2:]
		}
	}
	return strings.Join(phones, ",")
}

func Phone(phone string) string {
	phonesplit := strings.FieldsFunc(phone, func(r rune) bool {
		return r == ',' || r == ';'
	})

	phones := []string{}
	for _, phone := range phonesplit {
		arr := make([]rune, 0)
		for _, r := range phone {
			if r >= '0' && r <= '9' {
				arr = append(arr, r)
			}
		}
		number := string(arr)
		if number != "" {
			phones = append(phones, number)
		}
	}

	if len(phones) == 0 {
		return ""
	}

	// remove 84
	for i, phone := range phones {
		if strings.HasPrefix(phone, "0") {
			phones[i] = "84" + phone[1:]
		}
	}
	return phones[0]
}

func ContainString(ss []string, s string) bool {
	for _, i := range ss {
		if s == i {
			return true
		}
	}
	return false
}

func firstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

func GetUserDisplayName(u *User) string {
	countryName := ""
	cityName := ""
	fullname := ""
	email := ""

	for _, attr := range u.GetAttributes() {
		if attr.GetKey() == "trace_country_name" {
			countryName = attr.GetText()
		}
		if attr.GetKey() == "trace_city_name" {
			cityName = attr.GetText()
		}
		if attr.GetKey() == "fullname" {
			fullname = attr.GetText()
		}
		if attr.GetKey() == "emails" {
			email = attr.GetText()
		}
	}

	if fullname != "" && email != "" {
		return fullname + " (" + email + ")"
	}
	if fullname != "" {
		return fullname
	}
	if email != "" {
		return email
	}
	if cityName != "" {
		return countryName + " (" + cityName + ") #" + last4Chars(u.GetId())
	}
	return countryName + " #" + last4Chars(u.GetId())
}

func last4Chars(s string) string {
	if len(s) < 4 {
		return s
	}
	return s[len(s)-4:]
}
