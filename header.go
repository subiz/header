package header

import (
	hexa "encoding/hex"
	"encoding/json"
	"math"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	cpb "github.com/subiz/header/common"
	"github.com/subiz/log"
)

var updateTable = map[string]bool{
	"user-user":      true,
	"user-agent":     false,
	"user-connector": false,
	"user-system":    false,

	"agent-user":      true,
	"agent-agent":     true,
	"agent-connector": true,
	"agent-system":    true,

	"system-user":      true,
	"system-agent":     true,
	"system-connector": true,
	"system-system":    true,

	"connector-user":      true,
	"connector-agent":     false,
	"connector-connector": true,
	"connector-system":    true,
}

func CanUpdate(byType string, oldType string) bool {
	if byType == cpb.Type_subiz.String() {
		byType = "system"
	} else if byType == cpb.Type_agent.String() {
		byType = "agent"
	} else if byType == cpb.Type_connector.String() {
		byType = "connector"
	} else {
		byType = "user"
	}

	if oldType == cpb.Type_subiz.String() {
		oldType = "system"
	} else if oldType == cpb.Type_agent.String() {
		oldType = "agent"
	} else if oldType == cpb.Type_connector.String() {
		oldType = "connector"
	} else {
		oldType = "user"
	}

	return updateTable[byType+"-"+oldType]
}

// make value in prop primary if existed
func MakePrimaryValue(user *User, prop string, value string) {
	if value == "" {
		return
	}
	for _, attr := range user.Attributes {
		if attr.Key != prop {
			continue
		}
		if attr.Text == value {
			return // already primary
		}

		if attr.Text == "" {
			attr.Text = value
			return
		}

		found := false
		newOtherValues := []string{}
		for _, val := range attr.OtherValues {
			if val == value {
				found = true
				continue
			}

			// dont dup attr.Text
			if val == attr.Text {
				continue
			}

			newOtherValues = append(newOtherValues, val)
		}

		if !found {
			return
		}
		attr.OtherValues = append([]string{attr.Text}, newOtherValues...)
		attr.Text = value
		return
	}
}

func cleanOtherValues(attr *Attribute) {
	newOtherValues := []string{} // do not init nil
	for _, val := range attr.OtherValues {
		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}
		if val == attr.Text {
			continue
		}
		newOtherValues = append(newOtherValues, val)
	}
	// other values should be unique
	attr.OtherValues = Unique(newOtherValues)
}

// PushOtherValues adds newvalue to attr.OtherValues
// also, it keeps the size of attr.OtherValues in check, it will remove old data
// if the size is too big. The max size is 5KB
func PushOtherValues(attr *Attribute, newvalue string) {
	const maxSize = 5024
	// this value is too big, cannot push in
	if len(newvalue) > maxSize {
		attr.OtherValues = Unique(attr.OtherValues)
		return
	}
	if newvalue == "" {
		return
	}
	if attr.Text == newvalue {
		return
	}
	attr.OtherValues = Unique(append(attr.OtherValues, newvalue))

	nOtherValues := len(attr.OtherValues)
	if nOtherValues == 0 || nOtherValues == 1 {
		return
	}

	size := 0
	nKeep := 0
	for i := nOtherValues - 1; i >= 0; i-- {
		val := attr.OtherValues[i]
		if size+len(val)+len(newvalue) > maxSize {
			// cannot add must break
			break
		}
		nKeep++
		size += len(val)
	}
	attr.OtherValues = attr.OtherValues[nOtherValues-nKeep:]
}

func IsSameAttr(def *AttributeDefinition, a, b *Attribute) bool {
	typ := def.GetType()
	if typ == "" {
		typ = "text"
	}

	if a.OtherValues == nil && b.OtherValues != nil {
		return false
	}

	// check other value
	if b.OtherValues != nil {
		if len(a.OtherValues) != len(b.OtherValues) {
			return false
		}

		for i := range a.OtherValues {
			if a.OtherValues[i] != b.OtherValues[i] {
				return false
			}
		}
	}

	if typ == "text" {
		return a.Text == b.Text
	}

	if typ == "number" {
		return math.Abs(a.Number-b.Number) <= 0.0000003
	}

	if typ == "boolean" {
		return a.Boolean == b.Boolean
	}

	if typ == "datetime" {
		return a.Datetime == b.Datetime
	}

	return false
}

func UpdateAttribute(defM map[string]*AttributeDefinition, user *User, attr *Attribute) (updatedUser bool) {
	if attr.GetAction() == "noop" {
		return false
	}

	if attr.GetKey() == "" || user == nil || defM == nil {
		return false
	}
	by := attr.By
	bytype := attr.ByType
	// find the oldprop
	isBySystem := bytype == cpb.Type_subiz.String()
	isManually := bytype == cpb.Type_agent.String()
	isByConnector := bytype == cpb.Type_connector.String()
	isByCollector := !isByConnector && !isBySystem && !isManually // (bot, widget, user)

	def := defM[attr.Key]
	// undefined attribute, ignore
	if def == nil {
		return false
	}
	if def.IsSystem && def.IsReadonly {
		if isManually || isByCollector {
			return false
		}
	}
	isEmpty := false
	if def.Type == "text" {
		isEmpty = attr.Text == ""
	}
	if def.Type == "datetime" {
		isEmpty = attr.Datetime == ""
	}
	action := attr.GetAction()
	if action == "" {
		action = Attribute_upsert.String()
	}
	var oldattr *Attribute
	for _, a := range user.Attributes {
		if SameKey(a.Key, attr.Key) {
			oldattr = a
			break
		}
	}

	if oldattr == nil {
		oldattr = &Attribute{Key: attr.Key}
		if def.IsSystem && def.IsReadonly {
			oldattr.ByType = cpb.Type_subiz.String()
		}
		user.Attributes = append(user.Attributes, oldattr)
		updatedUser = true
	}

	canWrite := CanUpdate(bytype, oldattr.GetByType())

	// user can only push
	if action == Attribute_unshift.String() && (isByCollector || isByConnector) {
		action = Attribute_push.String()
	}

	// allow to push even user
	if action == Attribute_push.String() || action == Attribute_unshift.String() {
		canWrite = true
	}

	oldIsEmpty := ((def.Type == "text" && oldattr.GetText() == "") || (def.Type == "datetime" && oldattr.GetDatetime() == "") || oldattr == nil)
	if oldIsEmpty {
		canWrite = true
	}

	if def.PreventAutoOverride && !oldIsEmpty && isByCollector {
		canWrite = false
	}

	if canWrite && action == Attribute_delete.String() {
		oldattr.Text = ""
		oldattr.Number = 0
		oldattr.Boolean = false
		oldattr.Datetime = ""
		oldattr.OtherValues = []string{}
		oldattr.Modified = attr.Modified
		if oldattr.Modified == 0 {
			oldattr.Modified = time.Now().UnixMilli()
		}
		oldattr.By = by
		oldattr.ByType = bytype
		updatedUser = true
		return
	}

	if isEmpty {
		return
	}

	if !canWrite || IsSameAttr(def, oldattr, attr) {
		return
	}

	if action == Attribute_push.String() { // text only
		if oldattr.Text != "" {
			PushOtherValues(oldattr, attr.Text)
		} else {
			oldattr.Text = attr.Text
		}
		cleanOtherValues(oldattr)
		oldattr.Modified = attr.Modified
		if oldattr.Modified == 0 {
			oldattr.Modified = time.Now().UnixMilli()
		}
		oldattr.By = by
		oldattr.ByType = bytype
		updatedUser = true
	} else if action == "unshift" { // text only
		if oldattr.Text != "" {
			PushOtherValues(oldattr, oldattr.Text)
		}
		oldattr.Text = attr.Text
		cleanOtherValues(oldattr)
		oldattr.Modified = attr.Modified
		if oldattr.Modified == 0 {
			oldattr.Modified = time.Now().UnixMilli()
		}
		oldattr.By = by
		oldattr.ByType = bytype
		updatedUser = true
	} else {
		if (action == Attribute_insert.String() && oldIsEmpty) || action != Attribute_insert.String() {
			if attr.OtherValues != nil {
				oldattr.OtherValues = attr.OtherValues
			}
			oldattr.Text, oldattr.Number, oldattr.Boolean, oldattr.Datetime = attr.Text, attr.Number, attr.Boolean, attr.Datetime
			oldattr.Modified = attr.Modified
			if oldattr.Modified == 0 {
				oldattr.Modified = time.Now().UnixMilli()
			}
			oldattr.By = by
			oldattr.ByType = bytype
			updatedUser = true
		}
	}
	return updatedUser
}

func GetTimeAttr(u *User, key string) (time.Time, bool) {
	key = strings.ToLower(strings.TrimSpace(key))
	has := false
	for _, a := range u.GetAttributes() {
		if !SameKey(key, a.GetKey()) {
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
		if len(ss) > 0 {
			a.OtherValues = ss[1:]
			a.Text = ss[0]
		}
	case AttributeDefinition_datetime.String():
		t, _ := val.(time.Time)
		a.Datetime = t.Format(time.RFC3339)
	}
	for _, i := range u.GetAttributes() {
		if !SameKey(i.GetKey(), key) {
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

// 3 unmarshal and 2 marshal, very inefficient but its work
func AssignJSONByte(dst interface{}, body []byte) error {
	input := map[string]any{}
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	dstbyte, err := json.Marshal(dst)
	if err != nil {
		return err
	}
	dstmap := map[string]any{}
	if err := json.Unmarshal(dstbyte, &dstmap); err != nil {
		return err
	}

	for k, v := range input {
		dstmap[k] = v
	}

	dstbyte, err = json.Marshal(dstmap)
	if err != nil {
		return err
	}

	return json.Unmarshal(dstbyte, dst)
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
//
//	if the number starts with 84 => replace to 0 since we mostly serve
//	Vietnamese customers
func NormPhone(phone string) string {
	phonesplit := strings.FieldsFunc(phone, func(r rune) bool {
		return r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/'
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

// NormPhone converts an user input phone number to
// standalized phone number
// (84)35 9423 423 => 0359423423
//
// A good phone number only contain numbers
//
//	if the number starts with 84 => replace to 0 since we mostly serve
//	Vietnamese customers
func NormEmail(email string) string {
	emailsplit := strings.FieldsFunc(email, func(r rune) bool {
		return r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/' || r == ' '
	})

	emails := map[string]bool{}
	for _, email := range emailsplit {
		arr := make([]rune, 0)
		for _, r := range email {
			if r != '"' && r != '\'' {
				arr = append(arr, r)
			}
		}
		email = string(arr)
		email = strings.ToLower(email)
		if email == "" {
			continue
		}
		email = Substring(email, 0, 320)
		emails[email] = true
	}
	out := ""
	for em := range emails {
		out += em + ","
	}

	if len(out) > 0 {
		out = out[:len(out)-1] // remove last ,
	}
	return out
}

func EmailAddress(email string) string {
	email = strings.TrimSpace(email)
	emailsplit := strings.FieldsFunc(email, func(r rune) bool {
		return r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/' || r == ' '
	})

	if email == "" {
		return ""
	}

	if len(emailsplit) > 0 {
		email = emailsplit[0]
	} else {
		return ""
	}

	em, _ := mail.ParseAddress(email)
	if em == nil {
		return ""
	}

	email = em.Address
	if email == "" {
		return ""
	}

	arr := make([]rune, 0)
	for _, r := range email {
		if r != '"' && r != '\'' {
			arr = append(arr, r)
		}
	}
	email = strings.ToLower(string(arr))

	// RF3696: That limit is a maximum of 64 characters (octets)
	// in the "local part" (before the "@") and a maximum of 255 characters
	// (octets) in the domain part (after the "@") for a total length of 320
	// characters. However, there is a restriction in RFC 2821 on the length of an
	// address in MAIL and RCPT commands of 256 characters.  Since addresses
	// that do not fit in those fields are not normally useful, the upper
	// limit on address lengths should normally be considered to be 256.
	return Substring(email, 0, 320)
}

func PhoneNumber(phone string) string {
	phonesplit := strings.FieldsFunc(phone, func(r rune) bool {
		return r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/'
	})
	if len(phonesplit) == 0 {
		return ""
	}
	phone = phonesplit[0]
	arr := make([]rune, 0)
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			arr = append(arr, r)
		}
	}
	phone = string(arr)
	if strings.HasPrefix(phone, "0") {
		phone = "84" + phone[1:]
	}
	return phone
}

func ContainString(ss []string, s string) bool {
	for _, i := range ss {
		if s == i {
			return true
		}
	}
	return false
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

func Unique(slice []string) []string {
	if len(slice) == 0 || len(slice) == 1 {
		return slice
	}

	if len(slice) == 2 {
		if slice[0] == slice[1] {
			return slice[:1]
		}
	}

	keys := make(map[string]bool)
	list := []string{}
	hasDupplicated := false
	for _, entry := range slice {
		_, has := keys[entry]
		if !has {
			keys[entry] = true
			list = append(list, entry)
		} else {
			hasDupplicated = true
		}
	}

	// dont try to change the original slice
	if !hasDupplicated {
		return slice
	}
	return list
}

func Substring(s string, start int, end int) string {
	if start == 0 && end >= len(s) {
		return s
	}

	start_str_idx := 0
	i := 0
	for j := range s {
		if i == start {
			start_str_idx = j
		}
		if i == end {
			return s[start_str_idx:j]
		}
		i++
	}
	return s[start_str_idx:]
}

func Fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func E400(err error, code E, v ...interface{}) error {
	field := log.M{}
	for i, vv := range v {
		field[strconv.Itoa(i)] = vv
	}

	return log.Error(err, field, log.E_invalid_input, log.E(code.String()))
}

func E500(err error, code E, v ...interface{}) error {
	field := log.M{}
	for i, vv := range v {
		field[strconv.Itoa(i)] = vv
	}

	return log.Error(err, field, log.E_internal, log.E(code.String()))
}

func GetUserType(u *User) string {
	if u == nil {
		return ""
	}
	if u.Type == "contact" {
		return "contact"
	}

	typ := u.GetType()
	if u.Channel == "account" || u.Channel == "import" || u.Channel == "call" || u.Channel == "email" || u.Channel == "facebook" || u.Channel == "zalo" || u.Channel == "instagram" {
		typ = "lead"
	}

	for _, attr := range u.GetAttributes() {
		if attr.Key == "lifecycle_stage" {
			if attr.Text != "" {
				typ = "lead"
			}
		}

		if attr.Key == "emails" && attr.Text != "" {
			return "contact"
		}

		if attr.Key == "phones" && attr.Text != "" {
			return "contact"
		}

		if attr.Key == "record_id" && attr.Text != "" {
			return "contact"
		}
	}

	if u.PrimaryId != "" && u.PrimaryId != u.Id {
		return "lead"
	}

	return typ
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

func ReportExtractor(user *User, firstContentView *Event) map[string]bool {
	keys := map[string]bool{}
	keys["all.-"] = true
	keys["channel."+Hex(user.Channel)] = true              // touchpoint
	if user.Channel != "email" && user.Channel != "call" { // zalo, instagram
		keys["channel_source."+Hex(user.Channel+"."+user.ChannelSource)] = true //
	}

	if user.Channel == "call" {
		keys["call_isp."+Hex(GetPhoneISP(user.ProfileId))] = true //
	}

	if user.Channel != "subiz" {
		return keys
	}

	// device.Source = com.GetSourceFromRef2(adsum, device.SourceReferrer, device.Referrer)
	dev := firstContentView.GetBy().GetDevice()
	if uurl, _ := url.Parse(firstContentView.GetData().GetProduct().GetUrl()); uurl != nil {
		domain := uurl.Hostname()
		domain = strings.TrimPrefix(domain, "www.")
		if domain != "" {
			keys["domain."+Hex(domain)] = true
		}
		keys["path."+Hex(domain+uurl.Path)] = true
	}

	if dev.GetUtm().GetSource() != "" {
		keys["utm_source."+Hex(dev.GetUtm().GetSource())] = true
	}

	if dev.GetUtm().GetName() != "" {
		keys["utm_campaign."+Hex(dev.GetUtm().GetName())] = true
	}

	if uurl, _ := url.Parse(firstContentView.GetBy().GetDevice().GetSourceReferrer()); uurl != nil {
		refdomain := RefineRefDomain(uurl.Hostname())
		if refdomain != "" {
			keys["ref_domain."+Hex(refdomain)] = true
		}
	}

	if dev.GetSource() != "" {
		keys["source."+Hex(dev.GetSource())] = true
	}
	return keys
}

func RefineRefDomain(refdomain string) string {
	if refdomain == "" {
		return ""
	}

	refdomain = strings.TrimPrefix(refdomain, "www.")
	refdomain = strings.TrimPrefix(refdomain, "m.")

	// facebook
	if strings.HasSuffix(refdomain, ".facebook.com") {
		refdomain = "facebook.com"
	}
	if strings.HasSuffix(refdomain, ".coccoc.com") {
		refdomain = "coccoc.com"
	}
	if strings.Contains(refdomain, "yahoo.co") {
		// yahoo.co.jp, yahoo.com.vn
		return "yahoo.com"
	}

	if strings.HasPrefix(refdomain, "yahoo.") {
		// yahoo.jp
		return "yahoo.com"
	}

	if strings.Contains(refdomain, "google.co") {
		// google.com.au
		refdomain = "google.com"
	}
	if strings.HasPrefix(refdomain, "google.") {
		// google.ca
		return "google.com"
	}

	if strings.HasSuffix(refdomain, ".doubleclick.net") {
		refdomain = "doubleclick.net"
	}
	if strings.HasSuffix(refdomain, ".messenger.com") {
		refdomain = "facebook.com"
	}

	if strings.HasSuffix(refdomain, ".googlesyndication.com") {
		// fd662a7a3f921c440915d527e1165132.safeframe.googlesyndication.com
		refdomain = "googlesyndication.com"
	}

	if strings.HasPrefix(refdomain, ".ampproject.net.") {
		// d-1700324721198241007.ampproject.net
		return "ampproject.net"
	}

	return refdomain
}

func Hex(str string) string {
	return hexa.EncodeToString([]byte(str))
}

func UnHex(str string) string {
	out, _ := hexa.DecodeString(str)
	return string(out)
}

func GetPhoneISP(number string) string {
	first3 := ""
	if len(number) < 5 {
		return "other"
	}
	if number[0] == '8' && number[1] == '4' {
		first3 = "0" + number[2:4]
	} else if number[0] == '0' {
		first3 = "0" + number[1:3]
	}
	if first3 == "059" || first3 == "099" {
		return "Gmobile"
	}
	if first3 == "056" || first3 == "058" || first3 == "092" {
		return "Vietnammobile"
	}
	if first3 == "070" || first3 == "079" || first3 == "077" || first3 == "076" || first3 == "078" || first3 == "090" || first3 == "093" || first3 == "089" {
		return "Mobifone"
	}

	if first3 == "087" {
		return "Itelecom"
	}

	if first3 == "083" || first3 == "084" || first3 == " 085" || first3 == " 081" || first3 == " 082" || first3 == " 088" || first3 == " 091" || first3 == " 094" {
		return "Vinaphone"
	}

	if first3 == "032" || first3 == "033" || first3 == "034" || first3 == "035" || first3 == "036" || first3 == "037" || first3 == "038" || first3 == "039" || first3 == "096" || first3 == "097" || first3 == "098" || first3 == "086" {
		return "Viettel"
	}
	return "other"
}
