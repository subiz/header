package header

import (
	_ "embed"
	hexa "encoding/hex"
	"encoding/json"
	"html"
	"math"
	"net/mail"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	cpb "github.com/subiz/header/common"
	"github.com/subiz/log"
	"google.golang.org/protobuf/proto"
)

//go:embed perm.json
var permJSON string

var updateTable = map[string]bool{
	"user-user":      true,
	"user-agent":     false,
	"user-connector": false,
	"user-system":    false,

	// oser: other user, (user A is on of secondary users of B, if A update on B, A is oser)
	"oser-user":      false,
	"oser-agent":     false,
	"oser-connector": false,
	"oser-system":    false,

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

func CanUpdate(userId string, byId string, byType string, oldType string) bool {
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

	if byType == "user" {
		if byType != "" && userId != byId {
			byType = "oser"
		}
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

func AttributeValue(defM map[string]*AttributeDefinition, attr *Attribute) string {
	def := defM[attr.Key]
	// undefined attribute, ignore
	if def == nil {
		return attr.Text
	}

	if def.Type == "datetime" {
		return attr.Datetime
	}

	if attr.Type == "boolean" {
		if attr.Boolean {
			return "true"
		}
		return "false"
	}

	if attr.Type == "number" {
		return strconv.FormatFloat(attr.Number, 'f', -1, 64)
	}

	// if def.Type == "text" || def.Type == "list" || def.Type == "" {
	return attr.Text
}

// system => only coin
// this function expects the parent pass in trusted data
func UpdateAttribute(cred *cpb.Credential, defM map[string]*AttributeDefinition, user *User, attr *Attribute, now int64) (updatedUser bool) {
	if attr.GetKey() == "" || user == nil || defM == nil || attr.GetAction() == "noop" {
		return false
	}

	def := defM[attr.Key]
	if def == nil {
		return false // ignore undefined attribute
	}

	byId, bytype := attr.By, attr.ByType
	if byId == "" {
		byId = cred.GetIssuer()
	}

	if bytype == "" {
		bytype = cred.GetType().String()
	}
	isBySystem := bytype == cpb.Type_subiz.String()
	isManually := bytype == cpb.Type_agent.String()
	isByConnector := bytype == cpb.Type_connector.String()
	isByUser := !isByConnector && !isBySystem && !isManually // (bot, widget, user)

	// only allow subiz to update system read-only field
	if def.IsSystem && def.IsReadonly {
		if isManually || isByUser {
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
		if a.Key == attr.Key {
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

	if isByConnector {
		oldattr.ConnectorValue = AttributeValue(defM, attr)
		oldattr.Modified = now
		updatedUser = true
	}

	if isByUser && byId == user.GetId() {
		oldattr.UserValue = AttributeValue(defM, attr)
		oldattr.Modified = now
		updatedUser = true
	}

	canWrite := CanUpdate(user.GetId(), byId, bytype, oldattr.GetByType())
	oldIsEmpty := ((def.Type == "text" && oldattr.GetText() == "") || (def.Type == "datetime" && oldattr.GetDatetime() == "") || oldattr == nil)
	if oldIsEmpty {
		canWrite = true
	}

	if def.PreventAutoOverride && !oldIsEmpty && isByUser {
		canWrite = false
	}

	if canWrite && action == Attribute_delete.String() {
		if attr.Text != "" {
			// remove specific value
			vals := []string{}
			if oldattr.Text != attr.Text {
				vals = append(vals, oldattr.Text)
			}

			for _, v := range oldattr.OtherValues {
				if v != attr.Text {
					vals = append(vals, v)
				}
			}
			oldattr.Text = ""
			oldattr.OtherValues = nil
			if len(vals) > 0 {
				oldattr.Text = vals[0]
				oldattr.OtherValues = vals[1:]
			}
			if len(oldattr.OtherValues) == 0 {
				oldattr.OtherValues = nil
			}
			oldattr.Number = 0
			oldattr.Boolean = false
			oldattr.Datetime = ""
			oldattr.Modified = attr.Modified
			if oldattr.Modified == 0 {
				oldattr.Modified = now
			}
			oldattr.By = byId
			oldattr.ByType = bytype
			updatedUser = true
			return
		}

		oldattr.Text = ""
		oldattr.Number = 0
		oldattr.Boolean = false
		oldattr.Datetime = ""
		oldattr.OtherValues = nil
		oldattr.Modified = attr.Modified
		if oldattr.Modified == 0 {
			oldattr.Modified = now
		}
		oldattr.By = byId
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
			oldattr.Modified = now
		}
		oldattr.By = byId
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
			oldattr.Modified = now
		}
		oldattr.By = byId
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
				oldattr.Modified = now
			}

			oldattr.By = byId
			oldattr.ByType = bytype
			updatedUser = true
		}
	}
	return updatedUser
}

func UpdateAttributeForce(cred *cpb.Credential, defM map[string]*AttributeDefinition, user *User, attr *Attribute, now int64) {
	if attr.GetKey() == "" || user == nil || defM == nil || attr.GetAction() == "noop" {
		return
	}

	def := defM[attr.Key]
	if def == nil {
		return // ignore undefined attribute
	}

	// dont trust user
	if cred.GetType() == cpb.Type_user || cred.GetType() == cpb.Type_unknown {
		attr.By = cred.GetIssuer()
		attr.ByType = cpb.Type_user.String()
	}

	byId, bytype := attr.By, attr.ByType
	if byId == "" {
		byId = cred.GetIssuer()
	}

	if bytype == "" {
		bytype = cred.GetType().String()
	}
	isBySystem := bytype == cpb.Type_subiz.String()
	isManually := bytype == cpb.Type_agent.String()
	isByConnector := bytype == cpb.Type_connector.String()
	isByUser := !isByConnector && !isBySystem && !isManually // (bot, widget, user)

	action := attr.GetAction()
	if action == "" {
		action = Attribute_upsert.String()
	}
	var oldattr *Attribute
	for _, a := range user.Attributes {
		if a.Key == attr.Key {
			oldattr = a
			break
		}
	}

	if oldattr == nil {
		oldattr = &Attribute{Key: attr.Key}
		user.Attributes = append(user.Attributes, oldattr)
	}

	if isByConnector {
		oldattr.ConnectorValue = AttributeValue(defM, attr)
		oldattr.Modified = now
	}

	if isByUser && byId == user.GetId() {
		oldattr.UserValue = AttributeValue(defM, attr)
		oldattr.Modified = now
	}

	if action == Attribute_delete.String() {
		if attr.Text != "" {
			// remove specific value
			vals := []string{}
			if oldattr.Text != attr.Text {
				vals = append(vals, oldattr.Text)
			}

			for _, v := range oldattr.OtherValues {
				if v != attr.Text {
					vals = append(vals, v)
				}
			}
			oldattr.Text = ""
			oldattr.OtherValues = nil
			if len(vals) > 0 {
				oldattr.Text = vals[0]
				oldattr.OtherValues = vals[1:]
			}
			oldattr.Number = 0
			oldattr.Boolean = false
			oldattr.Datetime = ""
			oldattr.Modified = attr.Modified
			if oldattr.Modified == 0 {
				oldattr.Modified = now
			}
			oldattr.By = byId
			oldattr.ByType = bytype
			return
		}

		oldattr.Text = ""
		oldattr.Number = 0
		oldattr.Boolean = false
		oldattr.Datetime = ""
		oldattr.OtherValues = nil
		oldattr.Modified = attr.Modified
		if oldattr.Modified == 0 {
			oldattr.Modified = now
		}
		oldattr.By = byId
		oldattr.ByType = bytype
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
			oldattr.Modified = now
		}
		oldattr.By = byId
		oldattr.ByType = bytype
		return
	} else if action == "unshift" { // text only
		if oldattr.Text != "" {
			PushOtherValues(oldattr, oldattr.Text)
		}
		oldattr.Text = attr.Text
		cleanOtherValues(oldattr)
		oldattr.Modified = attr.Modified
		if oldattr.Modified == 0 {
			oldattr.Modified = now
		}
		oldattr.By = byId
		oldattr.ByType = bytype
		return
	} else {
		oldIsEmpty := ((def.Type == "text" && oldattr.GetText() == "") || (def.Type == "datetime" && oldattr.GetDatetime() == "") || oldattr == nil)
		if (action == Attribute_insert.String() && oldIsEmpty) || action != Attribute_insert.String() {
			if attr.OtherValues != nil {
				oldattr.OtherValues = attr.OtherValues
			}
			oldattr.Text, oldattr.Number, oldattr.Boolean, oldattr.Datetime = attr.Text, attr.Number, attr.Boolean, attr.Datetime
			oldattr.Modified = attr.Modified
			if oldattr.Modified == 0 {
				oldattr.Modified = now
			}

			oldattr.By = byId
			oldattr.ByType = bytype
			return
		}
	}
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

// scope name -> perm -> bool
//
//	example: {
//	 "agent": {"account:read": true, ...},
//	 "account:read": {"account:read": true},
//	}
var ScopeM = map[string]map[string]bool{}

func _setPrimaryScope(k string) {
	if ScopeM[k] == nil {
		ScopeM[k] = map[string]bool{}
	}
	ScopeM[k][k] = true
}

type PermStruct struct {
	Scope string
	Perm  map[string]string
}

func expandScope(permfile map[string]*PermStruct, scope string) {
	permStruct := permfile[scope]

	if ScopeM[scope] == nil {
		ScopeM[scope] = map[string]bool{}
	}
	for _, subScope := range strings.Split(permStruct.Scope, " ") {
		if subScope == "" {
			continue
		}
		expandScope(permfile, subScope)
		for k, v := range ScopeM[subScope] {
			ScopeM[scope][k] = v
		}
	}
}

func init() {
	permfile := map[string]*PermStruct{}
	if err := json.Unmarshal([]byte(permJSON), &permfile); err != nil {
		panic(err)
	}
	objM := map[string]bool{} // list of object
	for scope, p := range permfile {
		if ScopeM[scope] == nil {
			ScopeM[scope] = map[string]bool{}
		}
		for obj, str := range p.Perm {
			for _, k := range strings.Split(str, " ") {
				if k == "" {
					continue
				}
				if ScopeM[scope] == nil {
					ScopeM[scope] = map[string]bool{}
				}
				ScopeM[scope][obj+":"+k] = true
				objM[obj] = true // ticket, ticket_type, account...
			}
		}
	}

	for scope := range permfile {
		expandScope(permfile, scope)
	}

	for obj := range objM {
		_setPrimaryScope(obj + ":read")
		_setPrimaryScope(obj + ":read:all")
		_setPrimaryScope(obj + ":read:own")
		_setPrimaryScope(obj + ":read:unassigned")
		_setPrimaryScope(obj + ":read:none")

		_setPrimaryScope(obj + ":import")
		_setPrimaryScope(obj + ":import:all")
		_setPrimaryScope(obj + ":import:own")
		_setPrimaryScope(obj + ":import:unassigned")
		_setPrimaryScope(obj + ":import:none")

		_setPrimaryScope(obj + ":invite")
		_setPrimaryScope(obj + ":invite:all")
		_setPrimaryScope(obj + ":invite:own")
		_setPrimaryScope(obj + ":invite:unassigned")
		_setPrimaryScope(obj + ":invite:none")

		_setPrimaryScope(obj + ":update")
		_setPrimaryScope(obj + ":update:all")
		_setPrimaryScope(obj + ":update:own")
		_setPrimaryScope(obj + ":update:unassigned")
		_setPrimaryScope(obj + ":update:none")

		_setPrimaryScope(obj + ":create")
		_setPrimaryScope(obj + ":create:all")
		_setPrimaryScope(obj + ":create:own")
		_setPrimaryScope(obj + ":create:unassigned")
		_setPrimaryScope(obj + ":create:none")

		_setPrimaryScope(obj + ":delete")
		_setPrimaryScope(obj + ":delete:all")
		_setPrimaryScope(obj + ":delete:own")
		_setPrimaryScope(obj + ":delete:unassigned")
		_setPrimaryScope(obj + ":delete:none")
	}
}

type ObjectAction string

const READ ObjectAction = "read"
const IMPORT ObjectAction = "import"
const INVITE ObjectAction = "invite" // member
const UPDATE ObjectAction = "update"
const CREATE ObjectAction = "create"
const DELETE ObjectAction = "delete"

type ObjectType string

const TICKET ObjectType = "ticket"
const TICKET_TYPE ObjectType = "ticket_type"
const TICKET_TEMPLATE ObjectType = "ticket_template"
const ACCOUNT ObjectType = "account"
const CONVERSATION ObjectType = "conversation"
const AGENT_GROUP ObjectType = "agent_group"
const CALL_CAMPAIGN ObjectType = "call_campaign"
const CAMPAIGN ObjectType = "campaign"
const RULE ObjectType = "rule"
const INTEGRATION ObjectType = "integration"
const MESSAGETEMPLATE ObjectType = "message_template"
const EMAIL_SIGNATURE ObjectType = "email_signature"
const TAG ObjectType = "tag"
const WEB_PLUGIN ObjectType = "web_plugin"
const FILE ObjectType = "file"
const LANG ObjectType = "lang"
const USER_LABEL ObjectType = "user_label"
const USER_VIEW ObjectType = "user_view"
const ORDER ObjectType = "order"
const SHOP_SETTING ObjectType = "shop_setting"
const CONVERSATION_MODAL ObjectType = "conversation_modal"
const CONVERSATION_SETTING ObjectType = "conversation_setting"
const PHONE_DEVICE = "phone_device"
const CALL_SETTING ObjectType = "call_setting"
const GREETING_AUDIO ObjectType = "greeting_audio"
const PRODUCT ObjectType = "product"
const USER ObjectType = "user"
const SEGMENT ObjectType = "segment"
const AUTO_SEGMENT ObjectType = "auto_segment"
const WHITELIST ObjectType = "whitelist"
const SUBSCRIPTION ObjectType = "subscription"
const WORKFLOW ObjectType = "workflow"
const ZNS_TEMPLATE ObjectType = "zns_template"
const AI_DATA ObjectType = "ai_data"
const AI_AGENT ObjectType = "ai_agent"
const BANK_ACCOUNT ObjectType = "bank_account"

// const INVOICE ObjectType = "invoice"
const BOT ObjectType = "bot"
const LIVE ObjectType = "live"
const ATTRIBUTE ObjectType = "attribute"
const AGENT ObjectType = "agent"
const WEBHOOK ObjectType = "webhook"
const SLA_POLICY ObjectType = "sla_policy"
const KNOWLEGE_BASE ObjectType = "knowledge_base"
const ARTICLE ObjectType = "article"
const ACCMGR ObjectType = "accmgr"
const REPORT ObjectType = "report"

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

// Assign(ag1, ag2, ["fullname", "email"])
func Assign(dst, src proto.Message, fields, skipfields []string) proto.Message {
	if dst == nil {
		return src
	}

	if src == nil {
		return dst
	}

	dst = proto.Clone(dst)
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
		return dst
	}
	if dstTypeOf != srcTypeOf {
		return dst
	}

	for i := 0; i < srcValueOf.NumField(); i++ {
		tf := srcTypeOf.Field(i)
		jsonname := strings.Split(tf.Tag.Get("json"), ",")[0]
		if jsonname == "-" {
			continue
		}

		// skip if not field that user interested in
		if ContainString(skipfields, jsonname) {
			continue
		}

		if len(fields) > 0 {
			if !ContainString(fields, jsonname) {
				continue
			}
		}

		if !dstValueOf.Field(i).CanSet() {
			continue
		}
		dstValueOf.Field(i).Set(srcValueOf.Field(i))
	}
	return dst
}

// NormPhone converts an user input phone number to
// standalized phone number
// (84)35 9423 423 => 0359423423
//
// A good phone number only contain numbers
//
//	if the number starts with 84 => replace to 0 since we mostly serve
//	Vietnamese customers
func NormPhone(phone string) []string {
	phonesplit := strings.FieldsFunc(phone, func(r rune) bool {
		return r == '\000' || r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/'
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

	// strings.Join(phones, ",")
	return phones
}

// VNPhone converts an user input string to
// an standalized phone number. The number must be a valid vietnamese phone
// number
// "Xin chao, sdt cua minh la (84)35 9423 423 => 84359423423
var acceptedVnPhoneCharacter = map[rune]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	'-': true, // 043-46345345
	' ': true, // 036 48 435
	',': true, // 036,4543
	'.': true, // 04346345345
	'(': true, // (036)435345345
	')': true, // (036)435345345
}

var acceptedVnPhoneCharacterStrict = map[rune]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	'-': true, // 043-46345345
	',': true, // 036,4543
	'.': true, // 04346345345
}

func VNPhone(str string) []string {
	phonesplit := strings.FieldsFunc(str, func(r rune) bool {
		return !acceptedVnPhoneCharacter[r]
	})

	strictphonesplit := strings.FieldsFunc(str, func(r rune) bool {
		return !acceptedVnPhoneCharacterStrict[r]
	})

	phonesplit = append(phonesplit, strictphonesplit...)

	phones := map[string]bool{}
	for _, phone := range phonesplit {
		arr := make([]rune, 0)
		for _, r := range phone {
			if r >= '0' && r <= '9' {
				arr = append(arr, r)
			}
		}
		number := string(arr)
		if number != "" {
			phones[number] = true
		}
	}

	validPhone := []string{}
	for phone := range phones {
		if len(phone) < 6 {
			continue
		}

		if len(phone) > 20 {
			continue
		}

		if strings.HasPrefix(phone, "1900") || strings.HasPrefix(phone, "1800") {
			validPhone = append(validPhone, phone)
			continue
		}

		// 364348593 => 0364348593
		if len(phone) == 9 && phone[0] != 0 {
			phone = "0" + phone
		}

		if strings.HasPrefix(phone, "0") {
			phone = "84" + phone[1:]
		}

		if len(phone) != 11 {
			continue
		}

		if strings.HasPrefix(phone, "84") {
			if GetVietnamPhoneISP("0"+phone[2:]) != "other" {
				// it is avalid vietnam phone
				validPhone = append(validPhone, phone)
				continue
			}
			continue
		}
		continue
	}
	return validPhone
}

// NormPhone converts an user input phone number to
// standalized phone number
// (84)35 9423 423 => 0359423423
//
// A good phone number only contain numbers
//
//	if the number starts with 84 => replace to 0 since we mostly serve
//	Vietnamese customers
func NormEmail(email string) []string {
	emailsplit := strings.FieldsFunc(email, func(r rune) bool {
		return r == '\000' || r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/' || r == ' '
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

	ems := []string{}
	for email := range emails {
		ems = append(ems, email)
	}
	return ems
}

func EmailAddress(email string) string {
	email = strings.TrimSpace(email)
	emailsplit := strings.FieldsFunc(email, func(r rune) bool {
		return r == '\000' || r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/' || r == ' '
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
		return r == '\000' || r == ',' || r == ';' || r == '\n' || r == '\\' || r == '/'
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
	if len(ss) == 0 {
		return false
	}

	if len(ss) == 1 {
		return ss[0] == s
	}

	if len(ss) == 2 {
		return ss[0] == s || ss[1] == s
	}

	if len(ss) == 3 {
		return ss[0] == s || ss[1] == s || ss[2] == s
	}

	if len(ss) == 4 {
		return ss[0] == s || ss[1] == s || ss[2] == s || ss[3] == s
	}

	if len(ss) == 5 {
		return ss[0] == s || ss[1] == s || ss[2] == s || ss[3] == s || ss[4] == s
	}

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
	if s == "" {
		return ""
	}

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

	return log.Error3(err, field, log.E_invalid_input, log.E(code.String()))
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
		if i.GetKey() != key {
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
		keys["call_isp."+Hex(GetVietnamPhoneISP(user.ProfileId))] = true //
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

func GetVietnamPhoneISP(number string) string {
	first3, first4 := "", ""
	if len(number) < 5 {
		return "other"
	}
	if number[0] == '8' && number[1] == '4' {
		first3 = "0" + number[2:4]
		first4 = "0" + number[2:5]
	} else if number[0] == '0' {
		first3 = "0" + number[1:3]
		first4 = "0" + number[1:4]
	}
	if first3 == "059" || first3 == "099" || first4 == "0199" {
		return "Gmobile"
	}
	if first3 == "052" || first3 == "056" || first3 == "058" || first3 == "092" || first4 == "0182" || first4 == "0186" || first4 == "0188" {
		return "Vietnammobile"
	}
	if first3 == "070" || first3 == "079" || first3 == "077" || first3 == "076" || first3 == "078" || first3 == "090" || first3 == "093" || first3 == "089" ||
		first4 == "0120" || first4 == "0121" || first4 == "0122" || first4 == "0126" || first4 == "0128" {
		return "Mobifone"
	}

	if first3 == "087" {
		return "Itelecom"
	}

	if first3 == "083" || first3 == "084" || first3 == "085" || first3 == "081" || first3 == "082" || first3 == "088" || first3 == "091" || first3 == "094" ||
		first4 == "0128" || first4 == "0123" || first4 == "0125" || first4 == "0127" || first4 == "0124" || first4 == "0126" || first4 == "0129" {
		return "Vinaphone"
	}

	if first3 == "032" || first3 == "033" || first3 == "034" || first3 == "035" || first3 == "036" || first3 == "037" || first3 == "038" || first3 == "039" || first3 == "096" || first3 == "097" || first3 == "098" || first3 == "086" ||
		first4 == "0169" || first4 == "0168" || first4 == "0167" || first4 == "0166" || first4 == "0165" || first4 == "0164" || first4 == "0163" || first4 == "0162" {
		return "Viettel"
	}
	return "other"
}

// getConversationIdFromEv lookups conversation id from *header.Event
// it returns empty string "" if not found any conversation id
// see automation-bot/utils.go for more details
func GetConversationIdFromEv(e *Event) string {
	switch e.GetType() {
	case RealtimeType_message_sent.String(),
		RealtimeType_message_updated.String(),
		RealtimeType_message_pong.String(),
		RealtimeType_message_referral.String():
		return e.GetData().GetMessage().GetConversationId()
	case RealtimeType_bot_terminated.String():
		return e.GetData().GetBotTerminated().GetConversationId()
	}

	switch e.GetType() {
	case RealtimeType_task_created.String(),
		RealtimeType_task_updated.String():
		if e.GetData().GetConversation().GetId() != "" {
			return e.GetData().GetConversation().GetId()
		}
		task := e.GetData().GetTask()
		if len(task.GetAssociatedConversations()) > 0 {
			return task.AssociatedConversations[0]
		}
		return ""
	case RealtimeType_call_answered.String(),
		RealtimeType_call_rang.String(),
		RealtimeType_call_transferred.String():
		return e.GetData().GetCallInfo().GetConversationId()
	}

	// header.RealtimeType_conversation_state_updated.String(),
	//header.RealtimeType_conversation_updated.String(),
	//header.RealtimeType_conversation_joined.String(),
	//header.RealtimeType_conversation_invited.String(),
	//	header.RealtimeType_conversation_left.String(),
	//header.RealtimeType_conversation_tagged.String(),
	//header.RealtimeType_conversation_untagged.String(),
	//header.RealtimeType_conversation_rating_requested.String(),
	//header.RealtimeType_conversation_rated.String():
	return e.GetData().GetConversation().GetId()
}

// setConversationIdFromEv lookups conversation id from *header.Event
// it returns empty string "" if not found any conversation id
func SetConversationId(e *Event, cid string) {
	data := e.GetData()
	if data == nil {
		data = &Data{}
	}
	switch e.GetType() {
	case RealtimeType_conversation_state_updated.String():
	case RealtimeType_conversation_updated.String():
	case RealtimeType_conversation_joined.String():
	case RealtimeType_conversation_invited.String():
	case RealtimeType_conversation_left.String():
	case RealtimeType_conversation_tagged.String():
	case RealtimeType_conversation_untagged.String():
	case RealtimeType_conversation_spam_unmarked.String():
	case RealtimeType_conversation_spam_marked.String():
	case RealtimeType_conversation_rating_requested.String():
	case RealtimeType_conversation_rated.String():
	default:
		switch e.GetType() {
		case RealtimeType_message_updated.String():
		case RealtimeType_message_pong.String():
		case RealtimeType_message_sent.String():
		case RealtimeType_message_referral.String():
		default:
			return
		}
		msg := data.GetMessage()
		if msg == nil {
			msg = &Message{}
			data.Message = msg
		}
		msg.ConversationId = cid
	}
	convo := data.GetConversation()
	if convo == nil {
		convo = &Conversation{}
		data.Conversation = convo
	}
	convo.Id = cid
}

type IResourceGroup interface {
	GetId() string
	GetPermissions() []*ResourceGroupMember
}

func DeltaToBlock(delta string) *Block {
	delta = strings.TrimSpace(delta)
	if delta == "" {
		return &Block{}
	}

	deltab := []byte(delta)
	var deltas = []any{}
	m := map[string]any{}
	json.Unmarshal(deltab, &m)
	if ops, has := m["ops"]; has {
		deltas, _ = ops.([]any)
	} else {
		json.Unmarshal(deltab, &deltas)
	}

	if len(deltas) == 0 {
		return &Block{}
	}

	blocks := []*Block{}
	for _, delta := range deltas {
		if delta == nil {
			continue
		}
		block := &Block{Type: "text"}
		dt := delta.(map[string]any)
		if dt == nil {
			continue
		}

		if attri := dt["attributes"]; attri != nil {
			attrs, _ := attri.(map[string]any)
			if attrs != nil {
				if _, has := attrs["italic"]; has {
					block.Italic = true
				}
				if _, has := attrs["bold"]; has {
					block.Bold = true
				}
				if _, has := attrs["underline"]; has {
					block.Underline = true
				}
			}
		}
		inserti := dt["insert"]
		if inserti == nil {
			continue
		}

		switch v := inserti.(type) {
		case string:
			block.Text = v
		case map[string]any:
			if emoji, has := v["emoji"]; has {
				block.Type = "emoji"
				emojicode, _ := emoji.(string)
				block.Attrs = map[string]string{"code": emojicode}
				break
			}

			if dynamicField, has := v["dynamicField"]; has {
				// block.Attrs = map[string]string{"code": emojicode}
				if dynamicField == nil {
					break
				}

				df, _ := dynamicField.(map[string]any)
				if df == nil {
					break
				}

				keyi := df["key"]
				if keyi == nil {
					break
				}
				key, _ := keyi.(string)

				block.Type = "dynamic-field"
				block.Text = key

				block.Attrs = map[string]string{}
				for k, v := range df {
					block.Attrs[k], _ = v.(string)
				}
				break
			}

			if mentioni, has := v["mention"]; has {
				// block.Attrs = map[string]string{"code": emojicode}
				if mentioni == nil {
					break
				}

				df, _ := mentioni.(map[string]any)
				if df == nil {
					break
				}

				mentiontype := "agent"
				if typi := df["type"]; typi != nil {
					mentiontype, _ = typi.(string)
				}

				var mentionid string
				if mentionidi := df["id"]; mentionidi != nil {
					mentionid, _ = mentionidi.(string)
				}

				var mentionfullname string
				if fullnamei := df["fullname"]; fullnamei != nil {
					mentionfullname, _ = fullnamei.(string)
				}

				if mentiontype == "agent" {
					block.Type = "mention-agent"
				}

				block.Text = "@" + mentionid
				if mentionid == "" {
					block.Text = "@" + mentionfullname
				}
				break
			}
		}

		blocks = append(blocks, block)
	}

	if len(blocks) == 0 {
		return &Block{}
	}

	if len(blocks) == 1 {
		return blocks[0]
	}

	return &Block{Type: "div", Content: blocks}
}

func CompileBlock(block *Block, data map[string]string) {
	if block == nil {
		return
	}
	if data["user.display_name"] == "" && data["user.fullname"] != "" {
		data["user.display_name"] = data["user.fullname"]
	}
	if data["user.display_name"] != "" && data["user.fullname"] == "" {
		data["user.fullname"] = data["user.display_name"]
	}

	if block.Type == "dynamic-field" {
		if len(block.Attrs) == 0 {
			return
		}
		key := block.Attrs["key"]
		key = strings.ToLower(strings.TrimSpace(key))
		if key == "" {
			return
		}
		value, has := data[key]
		if !has {
			if block.Text == "" {
				// do not change type since we are not able to compile it
				block.Text = block.Attrs["text"]
			}
			if block.Text == "" {
				block.Text = block.Attrs["fallback"]
			}
			return
		}

		if value == "" && strings.HasPrefix(key, "user") && block.Attrs["text"] != "" {
			value = block.Attrs["text"]
		}
		if value == "" {
			value = block.Attrs["fallback"]
		}

		// block.Type = "text"
		block.Text = value
		return
	}
	for _, block := range block.GetContent() {
		CompileBlock(block, data)
	}
}

func BlockToPlainText(block *Block) string {
	return strings.TrimSpace(blockToPlainText(block))
}

func blockToPlainText(block *Block) string {
	if block == nil {
		return ""
	}
	out := ""
	if block.Type == "bullet_list" || block.Type == "ordered_list" {
		for i, item := range block.GetContent() {
			prefix := "\n* "
			if block.Type == "ordered_list" {
				prefix = strconv.Itoa(i) + "\n. "
			}
			out += prefix + strings.TrimSpace(blockToPlainText(item))
		}
		return out
	}

	if block.Type == "heading" || block.Type == "paragraph" || block.Type == "div" {
		out += "\n"
	}

	if block.Type == "" || block.Type == "text" || block.Type == "link" || block.Type == "dynamic-field" {
		return out + block.Text
	}

	if block.Type == "mention-agent" || block.Type == "mention" {
		var name = ""
		if attrs := block.GetAttrs(); attrs != nil {
			name = attrs["name"]
			if name == "" {
				name = attrs["id"]
			}
		}
		return out + "@" + name
	}

	if block.Type == "horizontal_rule" {
		return out + "\n---\n"
	}

	if block.Type == "emoji" {
		var code string
		if len(block.Attrs) > 0 {
			code = block.Attrs["code"]
		}
		if code == "" {
			code = block.Text
		}

		if code == "" {
			return out
		}
		return out + EmojiM[":"+code+":"]
	}

	for _, block := range block.GetContent() {
		out += blockToPlainText(block)
	}
	return out
}

func BlockToHTML(block *Block) string {
	return eleToHTML(blockToEle(block))
}

func eleToHTML(root *sanitiziedHTMLElement) string {
	if root == nil {
		return ""
	}
	out := "<" + root.Tag

	if root.ExtraTag != "" {
		out += "><" + root.ExtraTag
	}
	if root.Style != "" {
		out += " style=\"" + root.Style + "\""
	}

	if root.Class != "" {
		out += " class=\"" + root.Class + "\""
	}

	attrs := []string{}
	for k, v := range root.Attrs {
		attrs = append(attrs, k+"=\""+v+"\"")
	}

	sort.Strings(attrs)
	if len(attrs) > 0 {
		out += " " + strings.Join(attrs, " ")
	}

	if selftClosingTag[root.Tag] {
		return out + "/>"
	}

	if len(root.Content) == 0 {
		out += ">" + root.Text
	} else {
		out += ">"
		for _, child := range root.Content {
			out += eleToHTML(child)
		}
	}
	if root.ExtraTag != "" {
		out += "</" + root.ExtraTag + ">"
	}
	out += "</" + root.Tag + ">"
	return out
}

var selftClosingTag = map[string]bool{"br": true, "img": true, "hr": true}

type sanitiziedHTMLElement struct {
	Tag      string
	ExtraTag string
	Content  []*sanitiziedHTMLElement
	Style    string
	Text     string
	Attrs    map[string]string
	Class    string
}

func blockToEle(block *Block) *sanitiziedHTMLElement {
	if block == nil {
		return nil
	}
	ele := &sanitiziedHTMLElement{}
	ele.Style = ""
	if block.Style != nil {
		b, _ := json.Marshal(StyleToString(block.Style))
		m := map[string]string{}
		json.Unmarshal(b, &m)
		for k, v := range m {
			ele.Style += html.EscapeString(k) + ": " + html.EscapeString(v) + ";"
		}
	}

	ele.Attrs = map[string]string{}
	if block.Title != "" {
		ele.Attrs["title"] = html.EscapeString(block.Title)
	}
	if block.Href != "" {
		ele.Attrs["href"] = html.EscapeString(block.Href)
	}

	ele.Class = html.EscapeString(block.Class)
	for k, v := range block.Attrs {
		ele.Attrs[html.EscapeString(k)] = html.EscapeString(v)
	}
	if block.Type == "bullet_list" || block.Type == "ordered_list" {
		if block.Type == "bullet_list" {
			ele.Tag = "ul"
		} else {
			ele.Tag = "ol"
		}
	}

	if block.Type == "list_item" {
		ele.Tag = "li"
	}

	if block.Type == "heading" {
		if block.Level < 1 {
			block.Level = 1
		}

		if block.Level > 6 {
			block.Level = 6
		}
		ele.Tag = "h" + strconv.Itoa(int(block.Level))
	}

	if block.Type == "paragraph" {
		ele.Tag = "p"
	}

	if block.Type == "div" {
		ele.Tag = "div"
	}

	if block.Type == "image" {
		ele.Tag = "img"
		ele.Attrs["src"] = block.GetImage().GetUrl()

		if block.GetImage().GetWidth() > 0 {
			ele.Attrs["width"] = strconv.Itoa(int(block.GetImage().GetWidth()))
		}

		if block.GetImage().GetHeight() > 0 {
			ele.Attrs["height"] = strconv.Itoa(int(block.GetImage().GetHeight()))
		}
	}

	ele.Text = html.EscapeString(block.Text)
	if block.Type == "" || block.Type == "text" || block.Type == "dynamic-field" {
		ele.Tag = "span"
		if block.Bold {
			ele.Tag = "b"
		}
		if block.Italic {
			if ele.Tag == "b" {
				ele.ExtraTag = "em"
			} else {
				ele.Tag = "em"
			}
		}

		if block.Underline {
			ele.Style += "text-decoration:underline;"
		}

		if block.StrikeThrough {
			ele.Style += "text-decoration: line-through;"
		}
	}

	if block.Type == "link" {
		ele.Tag = "a"
	}

	if block.Type == "mention-agent" {
		ele.Tag = "span"
	}

	if block.Type == "horizontal_rule" {
		ele.Tag = "hr"
	}

	if block.Type == "hard_break" {
		ele.Tag = "br"
	}

	if block.Type == "emoji" {
		var code string
		if len(block.Attrs) > 0 {
			code = block.Attrs["code"]
		}
		if code == "" {
			code = block.Text
		}

		if code != "" {
			ele.Text = EmojiM[":"+code+":"]
		}
	}

	if block.Type == "table" {
	}

	var contents []*sanitiziedHTMLElement
	for _, block := range block.GetContent() {
		contents = append(contents, blockToEle(block))
	}
	ele.Content = contents
	return ele
}

var EmojiM = map[string]string{
	":like:":                "👍",
	":unlike:":              "👎",
	":wink:":                "😉",
	":tongue-out:":          "😛",
	":tired:":               "😫",
	":surprised:":           "😲",
	":smiling:":             "😊",
	":sleepy:":              "😴 ",
	":sad:":                 "😔",
	":neutral:":             "😐",
	":heart-eyes:":          "😍",
	":grinning:":            "😀",
	":crying:":              "😭",
	":confused:":            "😕",
	":angry:":               "😠",
	":emoji--disappointed:": "😞",
	":smile:":               "😊",
	":open_mouth:":          "😮",
	":tired_face:":          "😫",
	":stuck_out_tongue:":    "😛",
	":moyai:":               "🗿",
	":thumbsdown:":          "👎",
	":dislike:":             "👎",
	":thumbsup:":            "👍",
}

func MapBlock(block *Block, predicate func(block *Block)) {
	if block == nil {
		return
	}
	predicate(block)
	for _, content := range block.GetContent() {
		if content == nil {
			continue
		}
		MapBlock(content, predicate)
	}
}

func StyleToString(style *Style) map[string]string {
	var out = map[string]string{}
	if style.BorderRadius != "" {
		out["border-radius"] = style.BorderRadius
	}
	if style.FontFamily != "" {
		out["font-family"] = style.FontFamily
	}
	if style.Color != "" {
		out["color"] = style.Color
	}
	if style.Background != "" {
		out["background"] = style.Background
	}
	if style.TextAlign != "" {
		out["text-align"] = style.TextAlign
	}
	if style.TextTransform != "" {
		out["text-transform"] = style.TextTransform
	}
	if style.FontStyle != "" {
		out["font-style"] = style.FontStyle
	}
	if style.FontWeight != "" {
		out["font-weight"] = style.FontWeight
	}
	if style.Width != "" {
		out["width"] = style.Width
	}
	if style.MaxWidth != "" {
		out["max-width"] = style.MaxWidth
	}
	if style.Height != "" {
		out["height"] = style.Height
	}
	if style.MaxHeight != "" {
		out["max-height"] = style.MaxHeight
	}
	if style.PaddingLeft != "" {
		out["padding-left"] = style.PaddingLeft
	}
	if style.PaddingRight != "" {
		out["padding-right"] = style.PaddingRight
	}
	if style.PaddingTop != "" {
		out["padding-top"] = style.PaddingTop
	}
	if style.PaddingBottom != "" {
		out["padding-bottom"] = style.PaddingBottom
	}
	if style.MarginLeft != "" {
		out["margin-left"] = style.MarginLeft
	}
	if style.MarginRight != "" {
		out["margin-right"] = style.MarginRight
	}
	if style.MarginTop != "" {
		out["margin-top"] = style.MarginTop
	}
	if style.MarginBottom != "" {
		out["margin-bottom"] = style.MarginBottom
	}
	if style.Position != "" {
		out["position"] = style.Position
	}
	if style.ObjectFit != "" {
		out["object-fit"] = style.ObjectFit
	}
	if style.LineHeight != "" {
		out["line-height"] = style.LineHeight
	}
	if style.BackgroundPosition != "" {
		out["background-position"] = style.BackgroundPosition
	}
	if style.Left != "" {
		out["left"] = style.Left
	}
	if style.Right != "" {
		out["right"] = style.Right
	}
	if style.Top != "" {
		out["top"] = style.Top
	}
	if style.Bottom != "" {
		out["bottom"] = style.Bottom
	}
	if style.Opacity != "" {
		out["opacity"] = style.Opacity
	}
	if style.Rotate != "" {
		out["rotate"] = style.Rotate
	}
	// filter
	if style.Blur != "" {
		out["blur"] = style.Blur
	}
	if style.Grayscale != "" {
		out["grayscale"] = style.Grayscale
	}
	if style.Overlay != "" {
		out["overlay"] = style.Overlay
	}
	if style.OverlayOpacity != "" {
		out["overlay-opacity"] = style.OverlayOpacity
	}
	if style.Flex != "" {
		out["flex"] = style.Flex
	}
	if style.FlexDirection != "" {
		out["flex-direction"] = style.FlexDirection
	}
	if style.FlexShrink != "" {
		out["flex-shrink"] = style.FlexShrink
	}
	if style.AlignItems != "" {
		out["align-items"] = style.AlignItems
	}
	if style.JustifyContent != "" {
		out["justify-content"] = style.JustifyContent
	}
	if style.Transform != "" {
		out["transform"] = style.Transform
	}
	if style.FontSize != "" {
		out["font-size"] = style.FontSize
	}
	if style.ZIndex != "" {
		out["z-index"] = style.ZIndex
	}
	if style.BorderBottom != "" {
		out["border-bottom"] = style.BorderBottom
	}
	if style.BorderLeft != "" {
		out["border-left"] = style.BorderLeft
	}
	if style.BorderTop != "" {
		out["border-top"] = style.BorderTop
	}
	if style.BorderRight != "" {
		out["border-right"] = style.BorderRight
	}
	if style.Border != "" {
		out["border"] = style.Border
	}
	if style.BoxShadow != "" {
		out["box-shadow"] = style.BoxShadow
	}
	if style.Overflow != "" {
		out["overflow"] = style.Overflow
	}
	if style.OverflowX != "" {
		out["overflow-x"] = style.OverflowX
	}
	if style.OverflowY != "" {
		out["overflow-y"] = style.OverflowY
	}
	if style.WhiteSpace != "" {
		out["white-space"] = style.WhiteSpace
	}
	if style.UserSelect != "" {
		out["user-select"] = style.UserSelect
	}
	if style.PointerEvents != "" {
		out["pointer-events"] = style.PointerEvents
	}
	return out
}

func ToErr(err *Error) *log.AError {
	if err == nil {
		return nil
	}
	return &log.AError{
		Id:      err.Id,
		Code:    err.Code,
		Number:  err.Number,
		Fields:  err.Fields,
		XHidden: err.XHidden,
		Message: err.Message,
	}
}

func FromErr(err *log.AError) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		Id:      err.Id,
		Code:    err.Code,
		Number:  err.Number,
		Fields:  err.Fields,
		Message: err.Message,
	}
}

// return valid utf8 string
func CleanString(str string) string {
	str = strings.Join(strings.Split(str, "\000"), "")
	return strings.ToValidUTF8(str, "")
}

var skipuserfields = map[string]bool{
	"owner":           true,
	"modified":        true,
	"seen":            true,
	"lead_source":     true,
	"lead_at":         true,
	"lifecycle_stage": true,
	"related_email":   true,
	"first_interact":  true,
}

func TruncateConversation(convo *Conversation) {
	convo.Tags = nil
	convo.Integration = nil
	if convo.LastMessageSent != nil {
		if convo.LastMessageSent.By != nil {
			convo.LastMessageSent.By.Device = nil
		}
	}
	convo.ResponseSec = 0
	for _, mem := range convo.GetMembers() {
		if mem.Membership != "active" {
			continue
		}
		mem.InvitedBy = nil
		mem.Membership = ""
		mem.ConversationId = ""
		mem.JoinedAt = 0
		mem.FirstMessageAt = 0
		mem.SeenAt = 0
		mem.ReceivedAt = 0
		mem.LastSent = 0
		mem.LastSentEventTime = 0
		mem.InviteReason = ""
		mem.TotalMessages = 0
	}
}

func TruncateUser(user *User) {
	user.LeadOwners = nil
	user.LeadConversionBys = nil
	user.FirstContentView = nil
	user.LatestContentView = nil
	user.StartContentView = nil
	user.Segments = nil
	user.LifecycleStages = nil
	user.Labels = nil
	user.Merged = 0
	user.MergedReason = ""
	user.PrimaryId = ""

	newuserattributes := []*Attribute{}
	for _, attr := range user.Attributes {
		if skipuserfields[attr.Key] {
			continue
		}

		if attr.Key != "created" && attr.Key != "fullname" && attr.Key != "focused" {
			if attr.UserValue == "" {
				continue
			}
		}
		attr.Modified = 0
		attr.ByType = ""
		attr.By = ""
		newuserattributes = append(newuserattributes, attr)
	}
	user.Attributes = newuserattributes
}
