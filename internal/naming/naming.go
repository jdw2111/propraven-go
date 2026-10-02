// Package naming converts OpenAPI names to Go identifiers. The generator and
// the generated-surface test share it so both agree on every name.
package naming

import (
	"go/token"
	"strings"
	"unicode"
)

// initialisms are written in upper case when they form a whole word.
var initialisms = map[string]bool{
	"api": true, "apn": true, "cbsa": true, "cmbs": true, "csv": true,
	"fips": true, "html": true, "http": true, "https": true, "id": true,
	"ids": false, "ip": true, "json": true, "llc": true, "sql": true,
	"ucc": true, "uri": true, "url": true, "uuid": true, "xml": true,
}

// Words splits s on non-alphanumerics and camelCase boundaries.
func Words(s string) []string {
	var words []string
	var cur []rune
	flush := func() {
		if len(cur) > 0 {
			words = append(words, string(cur))
			cur = cur[:0]
		}
	}
	rs := []rune(s)
	for i, r := range rs {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			flush()
			continue
		}
		if len(cur) > 0 && unicode.IsUpper(r) {
			prev := rs[i-1]
			nextLower := i+1 < len(rs) && unicode.IsLower(rs[i+1])
			if unicode.IsLower(prev) || unicode.IsDigit(prev) || (unicode.IsUpper(prev) && nextLower) {
				flush()
			}
		}
		cur = append(cur, r)
	}
	flush()
	return words
}

func capWord(w string) string {
	lw := strings.ToLower(w)
	if initialisms[lw] {
		return strings.ToUpper(lw)
	}
	if lw == "ids" {
		return "IDs"
	}
	rs := []rune(lw)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}

// Pascal converts s to an exported Go identifier ("parcel_id" -> "ParcelID",
// "trafficHistory" -> "TrafficHistory"). A leading digit gets a "V" prefix.
func Pascal(s string) string {
	s = strings.ReplaceAll(s, "+", " Plus ")
	var b strings.Builder
	for _, w := range Words(s) {
		b.WriteString(capWord(w))
	}
	out := b.String()
	if out == "" {
		return "Empty"
	}
	if unicode.IsDigit([]rune(out)[0]) {
		out = "V" + out
	}
	return out
}

// Camel converts s to an unexported identifier that is never a Go keyword
// or predeclared name used by generated code.
func Camel(s string) string {
	ws := Words(s)
	if len(ws) == 0 {
		return "v"
	}
	var b strings.Builder
	b.WriteString(strings.ToLower(ws[0]))
	for _, w := range ws[1:] {
		b.WriteString(capWord(w))
	}
	out := b.String()
	if unicode.IsDigit([]rune(out)[0]) {
		out = "v" + out
	}
	if token.IsKeyword(out) || reserved[out] {
		out += "_"
	}
	return out
}

var reserved = map[string]bool{
	"ctx": true, "params": true, "opts": true, "req": true, "out": true,
	"err": true, "s": true, "string": true, "int": true, "bool": true,
	"any": true, "error": true, "len": true, "new": true, "nil": true,
	"true": true, "false": true, "json": true, "url": true, "http": true,
}

// HeaderField names a header parameter: "X-CREDIT-TOKEN" -> "CreditToken",
// "X-PAYMENT" -> "Payment".
func HeaderField(name string) string {
	n := name
	if len(n) > 2 && strings.EqualFold(n[:2], "x-") {
		n = n[2:]
	}
	return Pascal(strings.ToLower(n))
}

// Service is the Go type of a namespace ("parcels" -> "ParcelsService").
func Service(group string) string { return Pascal(group) + "Service" }

// Field is the Client field of a namespace ("parcels" -> "Parcels").
func Field(group string) string { return Pascal(group) }

// Method is the Go method of an operation ("trafficHistory" -> "TrafficHistory").
func Method(method string) string { return Pascal(method) }

// IterMethod is the auto-paginating sibling ("absentee" -> "AbsenteeIter").
func IterMethod(method string) string { return Pascal(method) + "Iter" }
