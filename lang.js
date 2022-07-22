var lo = require("lodash");
var langmap = require("langmap");
const fs = require("fs");

// 2021 Nov 2 from https://www.six-group.com/en/products-services/financial-information/data-standards.html
let allCurrencyCodes = [
  "AFN",
  "EUR",
  "ALL",
  "DZD",
  "USD",
  "AOA",
  "XCD",
  "ARS",
  "AMD",
  "AWG",
  "AUD",
  "AZN",
  "BSD",
  "BHD",
  "BDT",
  "BBD",
  "BYN",
  "BZD",
  "XOF",
  "BMD",
  "INR",
  "BTN",
  "BOB",
  "BOV",
  "BAM",
  "BWP",
  "NOK",
  "BRL",
  "BND",
  "BGN",
  "BIF",
  "CVE",
  "KHR",
  "XAF",
  "CAD",
  "KYD",
  "CLP",
  "CLF",
  "CNY",
  "COP",
  "COU",
  "KMF",
  "CDF",
  "NZD",
  "CRC",
  "HRK",
  "CUP",
  "CUC",
  "ANG",
  "CZK",
  "DKK",
  "DJF",
  "DOP",
  "EGP",
  "SVC",
  "ERN",
  "SZL",
  "ETB",
  "FKP",
  "FJD",
  "XPF",
  "GMD",
  "GEL",
  "GHS",
  "GIP",
  "GTQ",
  "GBP",
  "GNF",
  "GYD",
  "HTG",
  "HNL",
  "HKD",
  "HUF",
  "ISK",
  "IDR",
  "XDR",
  "IRR",
  "IQD",
  "ILS",
  "JMD",
  "JPY",
  "JOD",
  "KZT",
  "KES",
  "KPW",
  "KRW",
  "KWD",
  "KGS",
  "LAK",
  "LBP",
  "LSL",
  "ZAR",
  "LRD",
  "LYD",
  "CHF",
  "MOP",
  "MKD",
  "MGA",
  "MWK",
  "MYR",
  "MVR",
  "MRU",
  "MUR",
  "XUA",
  "MXN",
  "MXV",
  "MDL",
  "MNT",
  "MAD",
  "MZN",
  "MMK",
  "NAD",
  "NPR",
  "NIO",
  "NGN",
  "OMR",
  "PKR",
  "PAB",
  "PGK",
  "PYG",
  "PEN",
  "PHP",
  "PLN",
  "QAR",
  "RON",
  "RUB",
  "RWF",
  "SHP",
  "WST",
  "STN",
  "SAR",
  "RSD",
  "SCR",
  "SLL",
  "SGD",
  "XSU",
  "SBD",
  "SOS",
  "SSP",
  "LKR",
  "SDG",
  "SRD",
  "SEK",
  "CHE",
  "CHW",
  "SYP",
  "TWD",
  "TJS",
  "TZS",
  "THB",
  "TOP",
  "TTD",
  "TND",
  "TRY",
  "TMT",
  "UGX",
  "UAH",
  "AED",
  "USN",
  "UYU",
  "UYI",
  "UYW",
  "UZS",
  "VUV",
  "VES",
  "VED",
  "VND",
  "YER",
  "ZMW",
  "ZWL",
  "XBA",
  "XBB",
  "XBC",
  "XBD",
  "XTS",
  "XXX",
  "XAU",
  "XPD",
  "XPT",
  "XAG"
];
allCurrencyCodes = lo.uniqBy(allCurrencyCodes, e => e);

function capitalizeFirstLetter(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

function toGoLocale(locale) {
  return capitalizeFirstLetter(locale.replace("-", "_"));
}

let curm = `
var AllCurrency = map[string]bool{
` + allCurrencyCodes.map(code => `	"`+code+`": true,`).join('\n') + `
}
`

let localem = `
var LocaleM = map[string]bool{
`;

let gocodeallstring = `

func GetAllI18ns(str *I18NString) []string {
	out := make([]string, 0)
`;

let gocode = `// CODE GENERATED BY lang.js. DO NOT MODIFY
package header

import "strings"

// GetI18n returns the value by its locale
func GetI18n(str *I18NString, locale, fallback string) string {
	if str == nil {
		return fallback
	}
`;

let proto = `// CODE GENERATED BY lang.js. DO NOT MODIFY
syntax = "proto3";

package header;

option go_package = "github.com/subiz/header";

message I18nString {`;
var i = 10;
var s = lo
  .map(langmap, (_, k) => k)
  .filter(k => k.indexOf("-") >= 0)
  .map(k => {
    i++;

    gocode += `	if locale == "${k}" {
		if str.${toGoLocale(k)} == "" && fallback != locale && fallback != "" {
			return GetI18n(str, fallback, fallback)
		}
		return str.${toGoLocale(k)}
	}

`;

    gocodeallstring += `	if strings.TrimSpace(str.${toGoLocale(k)}) != "" {
		out = append(out, str.${toGoLocale(k)})
	}

`;

    localem += `	"${k}": true,
`;
    proto += `
	string ${k.replace("-", "_")} = ${i}; // ${langmap[k].englishName}`;
  });

localem += "}\n";
gocode +=
  `	return fallback
}` + localem + curm;
gocodeallstring += `	return out
}`;

proto += `
	string custom = 250;
}`;

let curencySource =
  `
func toFP(f float32) int64 {
	return int64(float64(f) * 1000000)
}

// CalcFPV turns price concurrency code to fixed point arithmetic value
// (https://en.wikipedia.org/wiki/Fixed-point_arithmetic)
// must convert to base concurrency first
func CalcFPV(price *Price, code string) int64 {
	cur := strings.TrimSpace(strings.ToUpper(code))
	if cur != "" {
` +
  allCurrencyCodes
    .map(code => {
      return (
        `
		if cur == "` +
        code +
        `" {
			return toFP(price.Get` +
        code +
        `())
		}`
      );
    })
    .join("\n") +
  `	}

	// unknow currency codd => fallback to the first non-zero
` +
  allCurrencyCodes
    .map(code => {
      return (
        `
	if price.Get` +
        code +
        `() != 0 {
		return toFP(price.Get` +
        code +
        `())
	}`
      );
    })
    .join("\n") +
  `
	return 0
}

func SetCurrency(price *Price, code string, value float32) {
	cur := strings.TrimSpace(strings.ToUpper(code))
	if cur != "" {
		return
	}
` +
  allCurrencyCodes
    .map(code => {
      return (
        `
	if cur == "` +
        code +
        `" {
		price.` +
        code +
        ` = value
	}`
      );
    })
    .join("\n") +
  `
}
`;

// proto
let protoAllCurCodeStr = allCurrencyCodes
  .map((code, i) => "\tfloat " + code + " = " + (i + 10) + ";")
  .join("\n");

let curProto =
  `
message Price {
	string currency = 2;
	int64 FPV = 3; // fixed point arithmetic value, see https://en.wikipedia.org/wiki/Fixed-point_arithmetic

` +
  protoAllCurCodeStr +
  "\n}\n";

proto += curProto;

fs.writeFileSync(
  "/src/locale.generated.go",
  gocode + gocodeallstring + curencySource
);
fs.writeFileSync("/src/locale.generated.proto", proto);
